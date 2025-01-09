package er

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// BuiltInError 定义内置错误，保护错误信息不被更改。所有更改都派生一个 ShellError
type BuiltInError struct {
	code         uint32        // 错误码
	name         string        // 错误名，解释 Code 的意义
	msg          string        // 面向客户的错误描述
	level        zapcore.Level // 引入日志系统的错误等级，用于最终处理时的日志记录。
	sendToClient bool          // 是否发送给前端，如果是 false，返回 Unknown
}

func (e *BuiltInError) Error() string {
	return fmt.Sprintf("Error:%s, msg:%s, code:%d, level:%s, sendToClient:%v",
		e.name, e.msg, e.code, e.level.String(), e.sendToClient)
}

func (e *BuiltInError) String() string {
	return e.Error()
}

func (e *BuiltInError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, e.Error())
		} else {
			fmt.Fprintf(s, e.name)
		}
	case 's':
		fmt.Fprintf(s, e.Error())
	default:
		fmt.Fprintf(s, e.Error())
	}
}

func (e *BuiltInError) Code() uint32 {
	return e.code
}

func (e *BuiltInError) Name() string {
	return e.name
}

func (e *BuiltInError) Msg() string {
	return e.msg
}

func (e *BuiltInError) Level() zapcore.Level {
	return e.level
}

func (e *BuiltInError) SendToClient() bool {
	return e.sendToClient
}

// Fields 转成 zap.Field
func (e *BuiltInError) Fields() []zapcore.Field {
	fields := make([]zapcore.Field, 5)
	fields[0] = zap.Uint32("code", e.code)
	fields[1] = zap.String("name", e.name)
	fields[2] = zap.String("msg", e.msg)
	fields[3] = zap.String("level", e.level.String())
	fields[4] = zap.Bool("sendToClient", e.sendToClient)
	return fields
}

// ToMap 将 ShellError 转为 Map
func (e *BuiltInError) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":         e.code,
		"name":         e.name,
		"msg":          e.msg,
		"level":        e.level.String(),
		"sendToClient": e.sendToClient,
	}
}

func (e *BuiltInError) new() *ShellError {
	return &ShellError{
		BuiltInError: *e,
		oriErr:       []error{e},
	}
}

// WithErr 添加原始错误
func (e *BuiltInError) WithErr(err error) IShellError {
	return e.new().WithErr(err)
}

func (e *BuiltInError) WithStackSkip(skip int) IShellError {
	return e.new().WithStackSkip(skip)
}

func (e *BuiltInError) WithStack() IShellError {
	return e.new().WithStackSkip(3)
}

func (e *BuiltInError) WithZapField(fields ...zapcore.Field) IShellError {
	return e.new().WithZapField(fields...)
}

// WSF 结合 WithZapField、WithStack 在记录调用栈后记录 Field
func (e *BuiltInError) WSF(fields ...zapcore.Field) IShellError {
	return e.new().WithStackSkip(3).WithZapField(fields...)
}

// WSEF 整合 WithStack, WithErr, WithZapField。注意：如果 oriErr 为空则返回 nil
func (e *BuiltInError) WSEF(oriErr error, fields ...zapcore.Field) IShellError {
	if oriErr == nil {
		return nil
	}
	return e.WithErr(oriErr).WithStackSkip(3).WithZapField(fields...)
}

// WithMsg 对用户解释错误的原因
func (e *BuiltInError) WithMsg(msg string) IShellError {
	err := e.new()
	err.msg = msg
	return err
}

// WithMsg 对用户解释错误的原因
func (e *BuiltInError) WithMsgf(format string, a ...any) IShellError {
	err := e.new()
	err.msg = fmt.Sprintf(format, a...)
	return err
}

// NewErr 声明内置错误，为了 errors.Is 正常捕获错误，返回指针。
func NewErr(code uint32, errDesc string, msg string, level zapcore.Level, sendToClient bool) *BuiltInError {
	return &BuiltInError{
		code:         code,
		name:         errDesc,
		msg:          msg,
		sendToClient: sendToClient,
		level:        level,
	}
}
