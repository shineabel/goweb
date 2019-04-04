package user

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/goweb/model"
	"github.com/goweb/db"
)

func Save(c *gin.Context)  {
	fmt.Println("save user %s in controller",c.Param("name"))

	//type UserData struct {
	//	Name string `json:"name" binding:"required"`
	//	Password string `json:"password" binding:"required"`
	//}
	//var ud UserData
	//if err := c.ShouldBindWith(&ud, binding.JSON); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	var u model.User


	if err := db.DB.Where(" name = ?", c.Param("name")).Find(&u).Error; err == nil {
		fmt.Println("sorry,user name %s exist.", u.Name)
	} else {
		fmt.Println("error occured:", err)
	}

}
