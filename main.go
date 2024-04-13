package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

func (p *persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p *persegi) keliling() float64 {
	return 4 * p.sisi
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func main() {
	var bangunDatarLingkaran hitung = lingkaran{diameter: 20.0}
	var bangunDatarPersegi hitung = &persegi{sisi: 10.0}

	fmt.Println(bangunDatarLingkaran.keliling())
	fmt.Println(bangunDatarPersegi.keliling())
}
