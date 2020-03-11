package admin

import (
	"meeting/models"

	"github.com/gin-gonic/gin"

	"meeting/pkg/app"
	"meeting/pkg/e"
)

func User_Info(c *gin.Context) {
	response := app.Response{C: c}

	id := c.PostForm("id")

	userinfo, err := models.Admin_GetUserInfo(id)
	if err != nil {
		response.SendErr(e.ERROR_DB, "")
		return
	}
	response.SendSucc(userinfo)
}
