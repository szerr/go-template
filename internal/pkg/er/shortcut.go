package er

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// WithStack 带有异常栈,加到所有原始异常上.
func WithStack(err error) IShellError {
	return Unknown.WithErr(err).WithStackSkip(3)
}

// WSEF (With Stack oriErr ZapField)用于包装底层错误，增加原始错误，debug 信息和调用栈。如果 oriErr 为空则返回 nil，适合作为 return 返回。
func WSEF(oriErr error, fields ...zapcore.Field) IShellError {
	if oriErr == nil {
		return nil
	}
	return WithStack(oriErr).WithZapField(fields...)
}

// AWSEF 在高层替代 WSEF，在底层错误没有改完时，兼容原始 error。对 ShellError 调用栈补全，并追加字段。如果 myErr 为空则返回 nil，适合作为 return 返回。
func AWSEF(err error, fields ...zapcore.Field) IShellError {
	if err == nil {
		return nil
	}
	var e IShellError
	var ok bool
	// 对非 ShellError 进行包裹
	if e, ok = err.(IShellError); !ok {
		e = Unknown.WithErr(err)
	}
	// 没有调用栈时添加
	if len(e.Stack()) == 0 {
		e.WithStackSkip(3)
	}

	// 追加 debug 信息
	return e.WithZapField(fields...)
}

// ReplaceErr 替换底层的 BuiltInError，替换 Msg 和 Code。如果 err 不是 IShellError，当作 BuiltInErr.WSEF(err) 处理
func ReplaceErr(err error, BuiltInErr *BuiltInError) IShellError {
	if e, ok := err.(IShellError); ok {
		return e.SetBuiltInErr(BuiltInErr)
	}
	return BuiltInErr.WithStack().WithErr(err)
}

// ToMap 将 myErr 转为 MapObject。原始错误会返回 map{"error": myErr}
func ToMap(err error) map[string]interface{} {
	if e, ok := err.(IShellError); ok {
		return e.ToMap()
	}
	if e, ok := err.(*BuiltInError); ok {
		return e.ToMap()
	}
	return map[string]interface{}{"error": err}
}

// WithMsg 带有错误提示信息
func WithMsg(err *ShellError, msg string) IShellError {
	return err.WithMsg(msg).WithStack()
}

// WithZap 支持 增加 zap 信息
func WithZapField(err *ShellError, msg string, fields ...zapcore.Field) IShellError {
	return err.WithMsg(msg).WithErr(err).WithZapField(fields...)
}

// ToZapFields 将 error 转为 ZapField
func ToZapFields(err error) []zap.Field {
	if e, ok := err.(*ShellError); ok {
		return e.Fields()
	} else {
		return []zap.Field{zap.Error(err)}
	}
}

// ToAddFields 将 error 转为 ZapField，并合并更多信息
func ToAddFields(err error, fields ...zapcore.Field) []zap.Field {
	return append(ToZapFields(err), fields...)
}

func ToJson(err error) string {
	if e, ok := err.(*ShellError); ok {
		return e.JsonIndent("", "  ")
	} else {
		return `"` + err.Error() + `"`
	}
}
