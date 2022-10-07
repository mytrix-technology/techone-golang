package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Error
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	// Panic
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}

	// Recover
	defer catch()

	var name1 string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name1)

	if valid, err := validate(name1); valid {
		fmt.Println("halo", name1)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
