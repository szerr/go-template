package main

import (
	"github.com/gin-gonic/gin"
	"go-template/api/engine"
	_ "go-template/docs" // 注意一定要加这一条，不然 swagger 会报错 找不到 doc.json
	"log"
	"net/http"
)

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

func goToDoc(c *gin.Context) {
	c.Redirect(http.StatusFound, "/swagger/index.html")
}

func newApp(e *engine.EngineDoc) (func() error, error) {
	return e.Run, nil
}
