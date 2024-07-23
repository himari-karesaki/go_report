package main

import (
	"fmt"
)

type preparer interface {
	prepare_trip() int
}

// type Trip struct {
// 	bikes    int
// 	custmers int
// 	vehicle  string
// 	prepare preparer
// }

func (t Trip) prepare_trip() int {
	return //bikesかcustmersかvehicleかを判定する処理
}

type Mechanic struct {
	prepare_bike int
}

// func (m Mechanic) prepare() {
// 	return //どのバイクを準備するかの処理
// 	//
// }

func (m Mechanic) prepare_trip() {
	return //どのバイクを準備するかの処理
}

type TripCoordinator struct {
	prepare_coordinator int
}

func (t TripCoordinator) prepare_trip() {
	return //どの食べ物を買ってくるかの処理
}

type Driver struct {
	prepare_vehicle string
	gas_up
}

func (d Driver) prepare_trip() {
	return //どの車両を準備するかの処理
	//ガスを満タンにする処理
	//タンクを水で満タンにする処理
}

func main() {
	var trip prepareInterface = &Trip{2, 3, "車"}
	m := &Mechanic{2, 3, "車"}
	c := &TripCoordinator{trip}
	d := &Driver{trip}

	fmt.Printf("bikesかcustomersかvehicleか: %d\n", d.prepare())

	fmt.Printf("bikes: %d\n", m.prepare())
	fmt.Printf("customers: %d\n", c.prepare())
	fmt.Printf("Driver: %s\n", d.prepare())

}