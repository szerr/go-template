package er

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"time"
)

// Go 运行一个 goroutine， logTag 用于标识 goroutine，防止
func Go(g func(), fields ...zapcore.Field) {
	defer RecoverWithStack("goroutine panic", fields...)
	g()
}

// GoWithRestart 运行一个 goroutine，recover 后自动重启
func GoWithRestart(ctx context.Context, g func(), fields ...zapcore.Field) {
	// TODO 考虑实现异常传递和基于心跳的监控。
	go func(g func()) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				Go(g)
				time.Sleep(time.Millisecond * 100) // 避免大批量持续的 Panic 重启
			}
		}
	}(g)
}

// RecoverWithStack 固定大小和缓冲区，使用更少的代码逻辑输出 Panic 的异常栈
func RecoverWithStack(msg string, fields ...zapcore.Field) any {
	err := recover()
	if err != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		fmt.Fprintf(os.Stderr, "global.Lg is not available: %v\n%s", err, buf[:n])
		// 为了让错误更快暴露
		panic(err)
	}
	return err
}
