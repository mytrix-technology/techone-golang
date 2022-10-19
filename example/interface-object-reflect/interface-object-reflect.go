package main

import (
	"fmt"
	"math"
	"reflect"
)

type hitung1 interface {
	luas1() float64
	keliling1() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas1() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling1() float64 {
	return math.Pi * l.diameter
}

type persegi1 struct {
	sisi float64
}

func (p persegi1) luas1() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi1) keliling1() float64 {
	return p.sisi * 4
}

func main() {
	// Interface
	var bangunDatar hitung1
	bangunDatar = persegi1{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas1())
	fmt.Println("keliling  :", bangunDatar.keliling1())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas1())
	fmt.Println("keliling  :", bangunDatar.keliling1())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	// Reflect
	var number = 23
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
}
