package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	ID        uint `gorm:"primaryKey"`
	Category  string
	Price     int
	CreatedAt time.Time
}

func main() {
	// データベースに接続
	dsn := "root:_MySQLPassword1@tcp(localhost:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("データベースへの接続に失敗しました:", err)
	}

	// テーブルを自動的に作成
	err = db.AutoMigrate(&Item{})
	if err != nil {
		fmt.Println("データベースのマイグレーションに失敗しました:", err)
	}

	// レコードを挿入
	items := []Item{
		{Category: "ピザ", Price: 2980},
		{Category: "焼肉セット", Price: 1980},
		{Category: "人参", Price: 1000},
		{Category: "白菜鍋", Price: 1000},
	}
	result := db.Create(&items)
	if result.Error != nil {
		fmt.Println("レコードの挿入に失敗しました")
	}
	// fmt.Println(result)
	fmt.Printf("%d レコード挿入しました。\n", result.RowsAffected)
}
