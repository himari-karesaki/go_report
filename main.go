package main

import "fmt"

// Item構造体の定義
type Item struct {
	Category string
	Price    int
}

func main() {

	var n int
	fmt.Print("何個の品物を登録しますか: ")
	fmt.Scan(&n)
	//Item構造体のスライスを作成
	items := []Item{}

	for i := 0; i < n; i++ {
		items = inputItem(items)
	}
	showItems(items)
	// showItems(items)
	// fmt.Println("===========")
	// for i := 0; i < len(items); i++ {
	// 	fmt.Println(items[i].Category, ":", items[i].Price, "円")
	// }
	// fmt.Println("===========")
}

func inputItem(items []Item) []Item {
	var item Item

	fmt.Print("品物を入力してください: ")
	fmt.Scan(&item.Category)
	fmt.Print("値段を入力してください: ")
	fmt.Scan(&item.Price)
	items = append(items, item)
	fmt.Println(items)
	return items
}

func showItems(items []Item) {
	fmt.Println("===========")

	// itemsの長さだけforを回す
	// len(items)はitemsの長さを返す
	for i := 0; i < len(items); i++ {

		fmt.Println(items[i].Category, ":", items[i].Price, "円")

	}

	fmt.Println("===========")
}
