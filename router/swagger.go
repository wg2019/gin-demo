package router

import ginSwagger "github.com/swaggo/gin-swagger"
import swaggerFiles "github.com/swaggo/files"

func init() {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json")))
}
