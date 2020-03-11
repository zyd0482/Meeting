package routers

import (
	// "net/http"

	"github.com/gin-gonic/gin"

	// _ "meeting/docs"
	"meeting/middleware/jwt"
	"meeting/routers/api/admin"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	// "meeting/middleware/jwt"
	// "meeting/pkg/export"
	// "meeting/pkg/qrcode"
	// "meeting/pkg/upload"
	// "meeting/routers/api"
	// "meeting/routers/api/v1"
)

func InitRouter(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	// r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	// r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	// r.GET("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.POST("/upload", api.UploadImage)

	// 用户登录
	r.POST("/api/admin/user/login", admin.User_Login)
	adminapi := r.Group("/api/admin")
	adminapi.Use(jwt.JWT())
	{
		// 获取用户信息
		adminapi.POST("/user/info", admin.User_Info)

		// 获取活动列表
		adminapi.POST("/meet/list", admin.Meet_List)
		// 添加活动
		adminapi.POST("/meet/add", admin.Meet_Add)
		// 获取活动详情
		adminapi.POST("/meet/detail", admin.Meet_Detail)
		// 更新活动
		adminapi.POST("/meet/update", admin.Meet_Update)
		// 删除活动
		adminapi.POST("/meet/delete", admin.Meet_Delete)
	}
}
