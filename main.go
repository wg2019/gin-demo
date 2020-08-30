package main

import (
	"fmt"
	_ "github.com/wg2019/gin-demo/docs"
	"github.com/wg2019/gin-demo/load"
	"github.com/wg2019/gin-demo/router"
	"log"
)

// @title gin-demo
// @version 1.0.0
// @description  gin-demo项目。
func main() {
	engine := router.Engine()
	err := engine.Run(fmt.Sprintf(":%d", load.ServerSetting.Port))
	if err != nil {
		log.Fatalf("ListenAndServe fail. err: %+v", err)
	}
}
