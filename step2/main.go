// todo: データベースへの記録

package main

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"

	_ "github.com/go-sql-driver/mysql"
)

// Item構造体の定義
type Item struct {
	Category string `json:"category" gorm:"primaryKey"`
	Price    int    `json:"price"`
}

// posts スライスを宣言
var items []Item

var db *gorm.DB

func main() {

	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベースへの接続に失敗しました:", err)
	}

	// マイグレーションを実行
	db.AutoMigrate(&Item{})
	var category string

	fmt.Print("品目を入力してください：")
	fmt.Scan(&category)

	// TODO:
	// データベースへ接続
	db, err := sql.Open("mysql", "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました：", err)
		return
	}
	defer db.Close()

	// クエリを実行
	var dbCategory string = category

	var item Item

	err = db.QueryRow("SELECT * FROM items WHERE title = ?", dbCategory).Scan(&item.Category, &item.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("指定されたタイトルのレコードが見つかりませんでした")
		} else {
			fmt.Println("クエリの実行に失敗しました:", err)
		}
		return
	}

	// 結果を出力
	fmt.Println(item.Category, item.Price)
}
