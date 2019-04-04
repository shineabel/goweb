package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goweb/controller/user"
)

func Route(r *gin.Engine)  {


	e := r.Group("/user")
	{
		e.POST("/save",user.Save)
	}
}
