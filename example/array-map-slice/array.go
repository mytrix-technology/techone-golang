package main

import "fmt"

func main() {
	// Array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Array")
	fmt.Println("Jumlah elemen: ", len(fruits))
	fmt.Println("Isi semua elemen: ", fruits)

	// Array Multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("\nArray Multi Dimensi")
	fmt.Println("numbers1: ", numbers1)
	fmt.Println("numbers2: ", numbers2)
}
