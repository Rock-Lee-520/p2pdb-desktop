package main

import (
	"api/models"
	_ "api/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// need to register models in init
	orm.RegisterModel(new(models.Emails))
	orm.RegisterModel(new(models.Configs))

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/p2pdb")

	// automatically build table
	orm.RunSyncdb("default", false, true)
}

func main() {
	//orm.RegisterDataBase("default", "mysql", beego.AppConfig.DefaultString("sqlconn", "root:@tcp(127.0.0.1:3306)/p2pdb"))
	beego.Run()
}
