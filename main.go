package main

import (
	_ "sitepointgoapp/routers"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/mattn/go-sqlite3"
  models "sitepointgoapp/models"
)

func init() {
  orm.RegisterDriver("sqlite", orm.DRSqlite)
  orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
  orm.RegisterModel(new(models.Article))
}

func main() {
	beego.Run()
}

