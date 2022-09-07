package main

import (
	_ "api/routers"

	"github.com/beego/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.DefaultString("sqlconn", "root:@tcp(127.0.0.1:3306)/p2pdb"))
	beego.Run()
}
