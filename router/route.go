package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goweb/controller/user"
)

func Route(r *gin.Engine)  {


	e := r.Group("/user")
	{
		e.POST("/save",user.Save)
		e.GET("/", user.GetUserList)
		e.GET("/:id", user.GetUserById)
		e.DELETE("/:id", user.DeleteUserById)
	}
	e2 := r.Group("/active")
	{

		e2.GET("/:name", user.ActiveUser)
	}
}
