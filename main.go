package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

import (
	"web-backend/go/dao"
	"web-backend/go/entity"
	"web-backend/go/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.SqlSession.AutoMigrate(&entity.Card{})
	r := routes.SetRouter()
	r.Run(":8081")
}
