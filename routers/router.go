package routers

import (
    // "net/http"

    "github.com/gin-gonic/gin"

    _ "meeting/docs"
    // "github.com/swaggo/gin-swagger"
    // "github.com/swaggo/gin-swagger/swaggerFiles"

    // "meeting/middleware/jwt"
    // "meeting/pkg/export"
    // "meeting/pkg/qrcode"
    // "meeting/pkg/upload"
    // "meeting/routers/api"
    // "meeting/routers/api/v1"
    "meeting/routers/api/admin_v1"
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


    adminv1 := r.Group("/api/admin/v1")
    // adminv1.Use(jwt.JWT())
    {
        // 获取活动列表
        adminv1.GET("/meets", admin_v1.Meet_List)
        // 添加活动
        adminv1.POST("/meets", admin_v1.Meet_Add)
        // 获取活动详情
        adminv1.GET("/meets/:id", admin_v1.Meet_Detail)
        // 更新活动
        adminv1.PUT("/meets/:id", admin_v1.Meet_Update)
        // 删除活动
        adminv1.DELETE("/meets/:id", admin_v1.Meet_Delete)
    }
    // apiv1 := r.Group("/api/v1")
    // apiv1.Use(jwt.JWT())
    // {
        // //获取标签列表
        // apiv1.GET("/tags", v1.GetTags)
        // //新建标签
        // apiv1.POST("/tags", v1.AddTag)
        // //更新指定标签
        // apiv1.PUT("/tags/:id", v1.EditTag)
        // //删除指定标签
        // apiv1.DELETE("/tags/:id", v1.DeleteTag)
        // //导出标签
        // r.POST("/tags/export", v1.ExportTag)
        // //导入标签
        // r.POST("/tags/import", v1.ImportTag)

        // //获取文章列表
        // apiv1.GET("/articles", v1.GetArticles)
        // //获取指定文章
        // apiv1.GET("/articles/:id", v1.GetArticle)
        // //新建文章
        // apiv1.POST("/articles", v1.AddArticle)
        // //更新指定文章
        // apiv1.PUT("/articles/:id", v1.EditArticle)
        // //删除指定文章
        // apiv1.DELETE("/articles/:id", v1.DeleteArticle)
        //生成文章海报
        // apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
    // }
}