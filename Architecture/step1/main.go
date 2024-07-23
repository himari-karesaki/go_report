package main

import (
	"fmt"
	"math"
)

// インターフェースで一旦メソッドと戻り値の型を定義している
type gear struct {
	flont_gear int
	rear_gear  int
	wheel      WheelInterface
}

// 　Gear インターフェースの定義
type GearInterface interface {
	gear_inches() int
	ratio() float64
}

func (g gear) gear_inches() int {
	return int(g.ratio()) * int(g.wheel.diameter())
}

func (g gear) ratio() float64 {
	return float64(g.flont_gear) / float64(g.rear_gear)
}

// wheelインターフェースの定義
type WheelInterface interface {
	diameter() float64
}

type wheel struct {
	rim  int
	tire int
}

func (w wheel) diameter() float64 {
	return float64(w.rim) + (float64(w.tire) * 2)
}

func main() {
	var Wheel WheelInterface = &wheel{rim: 16, tire: 14}
	var Gear GearInterface = &gear{flont_gear: 10, rear_gear: 20, wheel: Wheel}

	fmt.Printf("Front gear ratio: %.2f\n", Gear.ratio())
	fmt.Printf("Wheel diameter: %.2f inches\n", Wheel.diameter())
	fmt.Printf("Gear inches: %d\n", Gear.gear_inches())
	fmt.Printf("Wheel circumference: %.2f inches\n", Wheel.diameter()*math.Pi)                                 // 円周率πを使う場合
	fmt.Printf("Gear inches with wheel circumference: %d\n", Gear.gear_inches()*int(Wheel.diameter()*math.Pi)) // 円周率πを使う場合
}
