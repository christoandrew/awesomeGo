package main

import (
	"awesomeGo/products"
	"awesomeGo/tours"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	"github.com/GoAdminGroup/go-admin/engine"
	apiadmin "awesomeGo/admin"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte" // Import the theme
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := gin.Default()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	v1 := r.Group("/api")
	products.Products(v1)
	products.GetProduct(v1)
	tours.GetTours(v1)

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	// GoAdmin global configuration, can also be written as a json to be imported.
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": config.Database{
				Host:       "127.0.0.1",
				Port:       "3306",
				User:       "root",
				Pwd:        "",
				Name:       "go_api",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     "mysql",
			},
		},
		UrlPrefix: "admin", // The url prefix of the website.
		// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
	}

	// Import the business table configuration you need to manage here.
	// About Generators，see: https://github.com/GoAdminGroup/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(apiadmin.Generators)

	// Add configuration and plugins, use the Use method to mount to the web framework.
	_ = eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	_ = r.Run(":9033")

}

//func Migrate()  {
//	db := Database()
//	db.AutoMigrate(&tours.Activity{})
//	db.AutoMigrate(&tours.Feature{})
//	db.AutoMigrate(&tours.Lodge{})
//	db.AutoMigrate(&tours.Accommodation{})
//	db.AutoMigrate(&tours.Tour{})
//}
