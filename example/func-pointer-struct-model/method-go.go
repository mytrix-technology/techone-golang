package main

import "fmt"

func main() {
	// Method
	var s11 = student{"john wick", 21}
	fmt.Println("s11 before", s11.name)
	// john wick

	s11.changeName1("jason bourne")
	fmt.Println("s11 after changeName1", s11.name)
	// john wick

	s11.changeName2("ethan hunt")
	fmt.Println("s11 after changeName2", s11.name)
	// ethan hunt
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}
