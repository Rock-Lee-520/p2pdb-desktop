package main

import (
	"api/models"
	_ "api/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// need to register models in init
	orm.RegisterModel(new(models.Emails))
	orm.RegisterModel(new(models.Configs))
	orm.Debug = true
	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/p2pdb")

	// automatically build table
	orm.RunSyncdb("default", false, true)

}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
	}))

	//orm.RegisterDataBase("default", "mysql", beego.AppConfig.DefaultString("sqlconn", "root:@tcp(127.0.0.1:3306)/p2pdb"))
	beego.Run()
}
