package admin

import (
    _ "github.com/GoAdminGroup/go-admin/adapter/gin"
    _ "github.com/GoAdminGroup/themes/adminlte"
    _ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"

    "github.com/GoAdminGroup/go-admin/engine"
    // "github.com/GoAdminGroup/go-admin/examples/datamodel"
    "github.com/GoAdminGroup/go-admin/modules/config"
    "github.com/GoAdminGroup/go-admin/modules/language"
    "github.com/GoAdminGroup/go-admin/plugins/admin"
    // "github.com/GoAdminGroup/go-admin/template/types"
    "github.com/gin-gonic/gin"

    "meeting/pkg/setting"
    "meeting/admin/tables"
)

func InitAdmin(r *gin.Engine) {
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host: 		setting.DatabaseSetting.Host,
				Port: 		setting.DatabaseSetting.Port,
                User: 		setting.DatabaseSetting.User,
                Pwd:  		setting.DatabaseSetting.Password,
                Name: 		setting.DatabaseSetting.Name,
                MaxIdleCon: 50,
                MaxOpenCon: 150,
                Driver:    	config.DriverMysql,
			},
		},
		UrlPrefix: "admin", // 访问网站的前缀
        // Store 必须设置且保证有写权限，否则增加不了新的管理员用户
        Store: config.Store{
            Path:   "./uploads/admin",
            Prefix: "uploads",
        },
        Language: language.CN,
	}
	// r := gin.Default()
	eng := engine.Default()
	adminPlugin := admin.NewAdmin(tables.Generators)
    // adminPlugin := admin.NewAdmin(datamodel.Generators)
	// if err := eng.AddConfigFromJSON("./admin/config.json").
	// 	AddPlugins(adminPlugin).
	// 	Use(r); err != nil {
	// 	panic(err)
	// }
    if err := eng.AddConfig(cfg).
        AddPlugins(adminPlugin).
        Use(r); err != nil {
        panic(err)
    }

	// r.GET("/admin", func(ctx *gin.Context) {
	// 	engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
	// 		return DashboardPage()
	// 	})
	// })

	// _ = r.Run(":9033")
}