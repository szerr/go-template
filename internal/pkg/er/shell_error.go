package er

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"strconv"
	"strings"
)

type IShellError interface {
	Error() string
	Unwrap() []error
	WithErr(error) IShellError
	WithMsg(string) IShellError
	WithZapField(fields ...zapcore.Field) IShellError
	WithStack() IShellError
	WithStackSkip(skip int) IShellError
	StackStringSkip() string
	StackString() string
	WSEF(oriErr error, fields ...zapcore.Field) IShellError
	Code() uint32
	Name() string
	Msg() string
	SendToClient() bool
	SetLevel(level zapcore.Level) *ShellError
	SetBuiltInErr(err *BuiltInError) IShellError
	FieldsToMapObject() map[string]interface{}
	Fields() []zapcore.Field
	Stack() []uintptr
	Level() zapcore.Level
	ToMap() map[string]interface{}
	GoString() string
	Format(f fmt.State, verb rune)
	Json() string
}

const (
	Debug  zapcore.Level = zapcore.DebugLevel
	Info                 = zapcore.InfoLevel
	Warn                 = zapcore.WarnLevel
	Error                = zapcore.ErrorLevel
	DPanic               = zapcore.DPanicLevel
	Panic                = zapcore.PanicLevel
	Fatal                = zapcore.FatalLevel

	_minLevel = Debug
	_maxLevel = Fatal

	// InvalidLevel is an invalid value for zapcore.Level.
	//
	// Core implementations may panic if they see messages of this level.
	InvalidLevel = _maxLevel + 1
)

type ShellError struct {
	BuiltInError
	oriErr []error         // 原始错误，支持包裹多个错误
	fields []zapcore.Field // 整合 zap，带有更多 debug 信息
	stack  []uintptr       // 调用 WithStack 填充
}

func (e *ShellError) Error() string {
	return e.BuiltInError.Error() + fmt.Sprintf(", oriErr: %v, fields:%v", e.oriErr, e.FieldsToMapObject())
}

func (e *ShellError) String() string {
	return e.Error()
}

func (e *ShellError) GoString() string {
	return e.Error()
}

// Format Printf("%+v") 输出调用栈
func (e *ShellError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		fmt.Fprintf(s, e.Error())
		if s.Flag('+') {
			fmt.Fprintf(s, e.StackStringSkip())
		}
	case 's':
		fmt.Fprintf(s, e.Error())
	default:
		fmt.Fprintf(s, e.Error())
	}
}

func (e *ShellError) Stack() []uintptr {
	return e.stack
}

// Unwrap 兼容 1.20+
func (e *ShellError) Unwrap() []error {
	return e.oriErr
}

// Is 为了实现 1.19 兼容而做
func (e *ShellError) Is(target error) bool {
	if e == target {
		return true
	}
	for _, err := range e.oriErr {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

// As 为了实现 1.19 兼容而做
func (e *ShellError) As(target interface{}) bool {
	for _, err := range e.oriErr {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}

// SetLevel 设置错误级别
func (e *ShellError) SetLevel(level zapcore.Level) *ShellError {
	if level < _minLevel || level > _maxLevel {
		panic("level out of range")
	}
	e.level = level
	return e
}

// SetErr 替换底层的 BuiltInError，替换 Msg 和 Code。
func (e *ShellError) SetBuiltInErr(err *BuiltInError) IShellError {
	e.oriErr = append(e.oriErr, &e.BuiltInError)
	e.BuiltInError = *err
	return e
}

// FieldsToMapObject 将 fields 转为 MapObject。
func (e *ShellError) FieldsToMapObject() map[string]interface{} {
	enc := zapcore.NewMapObjectEncoder()
	for _, f := range e.fields {
		// TODO 测试
		//if f.Type == zapcore.ErrorType
		f.AddTo(enc)
	}
	return enc.Fields
}

// ToMap 将 ShellError 转为 MapObject。如果转换过程中发生异常，会变成其中的 myErr 字段。因为 map 的机制多个 myErr 会相互覆盖。
func (e *ShellError) ToMap() map[string]interface{} {
	m := e.FieldsToMapObject()
	for k, v := range e.BuiltInError.ToMap() {
		m[k] = v
	}
	m["oriErr"] = e.oriErr
	m["stack"] = e.StackStringSkip()
	return m
}

// ToJson 将数据转为 json，如果遇到错误会将错误一起封装
func (e *ShellError) Json() string {
	data, err := json.Marshal(e.ToMap())
	if err != nil {
		return fmt.Sprintln(`{"json marshal myErr": "`, err.Error(), `", "ShellError":"}`, e.ToMap(), `"`)
	}
	return string(data)
}

// ToJson 将数据转为 json，如果遇到错误会将错误一起封装
func (e *ShellError) JsonIndent(prefix string, indent string) string {
	data, err := json.MarshalIndent(e.ToMap(), prefix, indent)
	if err != nil {
		return fmt.Sprintln(`{
	"json marshal myErr": "`, err.Error(), `", 
	"ShellError":"`, e.ToMap(), `"
}`)
	}
	return string(data)
}

// ToZapFields 将 所有信息 打包返回 zapcore.Field
func (e *ShellError) Fields() []zapcore.Field {
	fields := append(e.fields, e.BuiltInError.Fields()...)
	fields = append(fields, zap.String("stack", e.StackStringSkip()))
	for _, e := range e.oriErr {
		fields = append(fields, zap.Error(e))
	}
	return fields
}

// StackStringSkip 自定义异常栈输出，尽量减少数据量
func (e *ShellError) StackStringSkip() string {
	frames := runtime.CallersFrames(e.stack)
	var stackBuilder strings.Builder
	for {
		frame, more := frames.Next()
		stackBuilder.WriteString(frame.Function)
		stackBuilder.WriteString("\n\t")
		stackBuilder.WriteString(frame.File)
		stackBuilder.WriteString(":")
		stackBuilder.WriteString(strconv.Itoa(frame.Line))
		stackBuilder.WriteString("\n")

		//fmt.Printf("- more:%v | %s\n", more, frame.Function)

		//进入到 runtime 或 testing 后不再抛出后续栈
		if strings.Contains(frame.File, "src/runtime/") || strings.Contains(frame.File, "src/testing/") {
			break
		}

		// Check whether there are more frames to process after this one.
		if !more {
			break
		}
	}
	return stackBuilder.String()
}

// StackString 输出全部异常栈
func (e *ShellError) StackString() string {
	frames := runtime.CallersFrames(e.stack)
	var stackBuilder strings.Builder
	for {
		frame, more := frames.Next()
		stackBuilder.WriteString(frame.Function)
		stackBuilder.WriteString("\n\t")
		stackBuilder.WriteString(frame.File)
		stackBuilder.WriteString(":")
		stackBuilder.WriteString(strconv.Itoa(frame.Line))
		stackBuilder.WriteString("\n")
		// Check whether there are more frames to process after this one.
		if !more {
			break
		}
	}
	return stackBuilder.String()
}

// WithErr 增加原始错误
func (e *ShellError) WithErr(oriErr error) IShellError {
	e.oriErr = append(e.oriErr, oriErr)
	return e
}

// WithMsg 更改 msg
func (e *ShellError) WithMsg(msg string) IShellError {
	e.msg = msg
	return e
}

// WithMsg 更改 msg
func (e *ShellError) WithMsgf(format string, a ...any) IShellError {
	e.msg = fmt.Sprintf(format, a...)
	return e
}

// WithZapField 支持 zapcore.Field，带上更多信息
func (e *ShellError) WithZapField(fields ...zapcore.Field) IShellError {
	e.fields = fields
	return e
}

// WithStack 记录当前调用栈，skip 表示要记录之前跳过的堆栈帧数
func (e *ShellError) WithStackSkip(skip int) IShellError {
	var pc [20]uintptr
	n := runtime.Callers(skip, pc[:])
	// 防止被多次调用覆盖，追加到 stack
	e.stack = append(e.stack, pc[:n]...)
	return e
}

// WithStack 记录当前调用栈
func (e *ShellError) WithStack() IShellError {
	return e.WithStackSkip(3)
}

// WSF 结合 WithZapField、WithStack 在记录调用栈后记录 Field
func (e *ShellError) WSF(fields ...zapcore.Field) IShellError {
	return e.WithStackSkip(3).WithZapField(fields...)
}

// WSEF 用于包装底层错误，增加原始错误，debug 信息和调用栈。如果 oriErr 为空则返回 nil，适合作为 return 返回。
func (e *ShellError) WSEF(oriErr error, fields ...zapcore.Field) IShellError {
	if oriErr == nil {
		return nil
	}
	return e.WithStackSkip(3).WithErr(oriErr).WithZapField(fields...)
}
