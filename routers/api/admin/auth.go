package admin

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"meeting/pkg/app"
	"meeting/pkg/e"
	"meeting/pkg/util"
	"meeting/service"
)

type adminAuth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func User_Login(c *gin.Context) {
	response := app.Response{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := adminAuth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		response.SendErr(0, "用户名密码格式不正确")
		return
	}

	authService := service.AdminAuth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		response.SendErr(e.ERROR_DB, "")
		return
	}

	if !isExist {
		response.SendErr(0, "用户名密码错误")
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		response.SendErr(e.ERROR_DB, "")
		return
	}

	data := make(map[string]interface{})
	data["token"] = token
	response.SendSucc(data)
}
