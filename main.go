package main

import "fmt"

// Item構造体の定義
type Item struct {
	Category string
	Price    int
}

func main() {
	//inputItem関数を呼び出してItem構造体の値を取得
	item := inputItem()

	fmt.Println("=============")

	fmt.Printf("%sの値段は%d円です\n", item.Category, item.Price)

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
