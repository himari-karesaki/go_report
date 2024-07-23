package main

type Bicycle struct {
	size  int
	parts Parts
}

func (b Bicycle) spares() int {
	return //Partsにスペアを要求する処理
}

type Parts struct {
	parts []Part
}

func (p Parts) spares() string {
	return //Partにスペアを要求する処理
}

type Part struct {
	name        string
	description string
	needsSpares bool
}
