package main

import (
	"fmt"
	"math"
)

type persegi struct {
	sisi float64
}
type hitung interface {
	luas() float64
}

func main() {
	var bangunDatar hitung
	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
}
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}
