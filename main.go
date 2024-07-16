package main

import "fmt"

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
