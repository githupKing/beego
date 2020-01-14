package main

import (
	_ "beego/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("MyDarkSecret")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(hashedPassword)

	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // nil means it is a match

	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource := getConfig("MySqlDataSource")
	//注册读写数据库
	orm.RegisterDataBase("default", "mysql", dataSource)
	//阿里云上主数据库的连接数配置是600
	orm.SetMaxOpenConns("default", 550)
	orm.RunSyncdb("default", false, false) //自动生成数据库
	beego.Info("MYSQL主数据库配置成功", dataSource)
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
		orm.RunCommand()
	}
	if beego.BConfig.WebConfig.EnableDocs {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
func getConfig(key string) string {
	return beego.AppConfig.String(key)
}
