package util

import "github.com/gin-gonic/gin"

func Quickprintln(c *gin.Context) {
	println("c.Request.Method:" + c.Request.Method)
	println("c.Request.ContentType:" + c.ContentType())
	// println("c.Request.Body:" + c.Request.Body)
	// c.Request.ParseForm()
	// println("c.Request.Form: %s", c.Request.PostForm)
	for k, v := range c.Request.PostForm {
		println("k:", k)
		println("v:", v)
	}

	// println("c.Request.ContentLength: %v", c.Request.ContentLength)
	// data, _ := ioutil.ReadAll(ctx.Request.Body)

	// logging.Debugf("c.Request.GetBody: %v", string(data))
}
