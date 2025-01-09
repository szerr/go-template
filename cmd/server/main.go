package main

import (
	"go-template/api/engine"
	"log"
)

// 启动服务
func main() {
	run, cleanup, err := wireApp()
	if err != nil {
		log.Panic(err)
	}
	defer cleanup()
	err = run()
	if err != nil {
		log.Panic(err)
	}
}

func newApp(e *engine.EngineAuth) (func() error, error) {
	return e.Run, nil
}
