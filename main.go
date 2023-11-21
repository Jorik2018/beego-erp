package main

import (
	"crudDemo/models"
	_ "crudDemo/routers"

	_ "github.com/lib/pq"

	"github.com/beego/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
	orm.RegisterModel(new(models.UserMasterTable), new(models.CarsMasterTable))
	orm.RunSyncdb("default", false, true)
	beego.Run()

}
