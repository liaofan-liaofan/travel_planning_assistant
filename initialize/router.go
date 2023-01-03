package initialize

import (
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.New()
	return router
}
