package routers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/wzcourses/demo-gads-app/controller"
)

func AppRouter() *gin.Engine {
	router := gin.Default()

	userCtrl := controller.NewUserController()
	router.GET("/ping", userCtrl.Ping)
	return router
}
