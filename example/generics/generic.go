package main

import "fmt"

func main() {
	text := MergeString([]string{"B", "T", "P", "N"})
	fmt.Println("text: ", text)
}

func MergeString[V string](array []V) V {
	var merge V
	for i := 0; i < len(array); i++ {
		merge += array[i]
	}
	return merge
}
