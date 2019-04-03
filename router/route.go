package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goweb/controller/user"
)

func Route(e *gin.Engine)  {

	e.GET("/save",user.Save)
}
