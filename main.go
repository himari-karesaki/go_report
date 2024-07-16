package main

import "fmt"

// Item構造体の定義
type Item struct {
	Category string
	Price    int
}

func main() {
	fmt.Println("Hello, World!")
	//品物を入れる変数の定義
	var category string

	//値段を入れる変数の定義
	var price int

	fmt.Print("品物を入力してください: ")
	fmt.Scan(&category)

	fmt.Print("値段を入力してください: ")
	fmt.Scan(&price)

	fmt.Println("=============")

	fmt.Printf("%sの値段は%d円です\n", category, price)

	fmt.Println("=============")
}

func inputItem() Item {
	var item Item
	fmt.Print("品物を入力してください: ")
	fmt.Scan(&item.Category)
	fmt.Print("値段を入力してください: ")
	fmt.Scan(&item.Price)
	return item
}
