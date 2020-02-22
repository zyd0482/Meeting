package app

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

	"meeting/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		fmt.Println(err)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		for _, vErr := range valid.Errors {
            fmt.Println(vErr.Key, vErr.Message)
        }
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}