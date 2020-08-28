package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wg2019/gin-demo/docs"
	"github.com/wg2019/gin-demo/load"
	"log"
)

// @title gin-demo
// @version 1.0.0
// @description  gin-demo项目。
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json")))
	err := r.Run(fmt.Sprintf(":%d", load.ServerSetting.Port))
	if err != nil {
		log.Fatalf("ListenAndServe fail. err: %+v", err)
	}
}
