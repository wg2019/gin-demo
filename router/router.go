// Package router 路由
package router

import (
	"github.com/gin-gonic/gin"
)

var engine = gin.New()

// Engine 路由
func Engine() *gin.Engine {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	return engine
}
