package db

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBの起動
func NewDB() *gorm.DB {

	var flag_env = flag.String("GO_ENV", "", "開発環境フラグ")
	flag.Parse()
	// 環境変数のdevの場合は .envの内容を読み込む
	//if os.Getenv("GO_ENV") == "dev" {
	if *flag_env == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// URLを生成(.env内の情報をURLに埋め込む)
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// postgreSQLへの接続
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db

}

// DBのクローズ
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
