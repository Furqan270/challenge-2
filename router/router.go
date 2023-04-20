package router

import (
	"chall2/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	controller.NewBookController().BookRoutes(r)

	return r
}
