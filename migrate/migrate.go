package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	// DBインスタンスのアドレスを受け取る
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// Goの構造体の定義に基づき、DBにテーブルの作成や、既存テーブルの更新を実施する
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
