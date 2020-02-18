package admin

import (
    _ "github.com/GoAdminGroup/go-admin/adapter/gin" // 引入适配器，必须引入，如若不引入，则需要自己定义
    _ "github.com/GoAdminGroup/themes/adminlte" // 引入主题，必须引入，不然报错
    // _ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // 引入对应数据库引擎
    _ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"

    "github.com/GoAdminGroup/go-admin/engine"
    // "github.com/GoAdminGroup/go-admin/examples/datamodel"
    // "github.com/GoAdminGroup/go-admin/modules/config"
    // "github.com/GoAdminGroup/go-admin/modules/language"
    "github.com/GoAdminGroup/go-admin/plugins/admin"
    // "github.com/GoAdminGroup/go-admin/template/types"
    "github.com/gin-gonic/gin"

    // "meeting/pkg/setting"
    "meeting/admin/tables"
)

func (r *gin.Engine) InitAdmin() {
	// cfg := config.Config{
	// 	Databases: config.DatabaseList{
	// 		// "default": {
	// 		// 	Host: 		setting.DatabaseSetting.Host,
	// 		// 	Port: 		setting.DatabaseSetting.Port,
 //   //              User: 		setting.DatabaseSetting.User,
 //   //              Pwd:  		setting.DatabaseSetting.Password,
 //   //              Name: 		setting.DatabaseSetting.Name,
 //   //              MaxIdleCon: 50,
 //   //              MaxOpenCon: 150,
 //   //              Driver:    	config.DriverMysql,
	// 		// },
 //            "default": {
 //                Host:       "127.0.0.1",
 //                Port:       "3306",
 //                User:       "root",
 //                Pwd:        "root1234",
 //                Name:       "mt-admin-go",
 //                MaxIdleCon: 50,
 //                MaxOpenCon: 150,
 //                Driver:     config.DriverMysql,
 //            },
	// 	},
	// 	UrlPrefix: "admin", // 访问网站的前缀
 //        // Store 必须设置且保证有写权限，否则增加不了新的管理员用户
 //        Store: config.Store{
 //            Path:   "./admin/uploads",
 //            Prefix: "uploads",
 //        },
 //        Language: language.CN,
	// }
	// r := gin.Default()
	eng := engine.Default()
	adminPlugin := admin.NewAdmin(tables.Generators)
    // adminPlugin := admin.NewAdmin(datamodel.Generators)
	if err := eng.AddConfigFromJSON("./admin/config.json").
		AddPlugins(adminPlugin).
		Use(r); err != nil {
		panic(err)
	}
    // if err := eng.AddConfig(cfg).
    //     AddPlugins(adminPlugin).
    //     Use(r); err != nil {
    //     panic(err)
    // }

	// r.GET("/admin", func(ctx *gin.Context) {
	// 	engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
	// 		return DashboardPage()
	// 	})
	// })

	// _ = r.Run(":9033")
}