package api

import (
    "net/http"

    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"

    "meeting/pkg/app"
    "meeting/pkg/e"
    "meeting/pkg/util"
    "meeting/service/auth_service"
)

type auth struct {
    Username string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
    response := app.Response{C: c}
    valid := validation.Validation{}

    username := c.Query("username")
    password := c.Query("password")

    a := auth{Username: username, Password: password}
    ok, _ := valid.Valid(&a)

    if !ok {
        app.MarkErrors(valid.Errors)
        response.Send(http.StatusBadRequest, e.INVALID_PARAMS, nil)
        return
    }

    authService := auth_service.Auth{Username: username, Password: password}
    isExist, err := authService.Check()
    if err != nil {
        response.Send(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
        return
    }

    if !isExist {
        response.Send(http.StatusUnauthorized, e.ERROR_AUTH, nil)
        return
    }

    token, err := util.GenerateToken(username, password)
    if err != nil {
        response.Send(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
        return
    }

    response.Send(http.StatusOK, e.SUCCESS, map[string]string{
        "token": token,
    })
}