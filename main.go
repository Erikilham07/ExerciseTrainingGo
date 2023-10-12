package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	var (
		name, jobtitle, department string
		age                        int
	)

	name = "Muhammad Ilham Farerik"
	jobtitle = "Implementation Consultant"
	department = "YO"
	age = 25

	fmt.Println("Nama saya adalah", name)
	fmt.Println("Saya bekerja sebagai", jobtitle, "di", department)
	fmt.Println("Umur saya adalah", age)

	const phi = 3.14
	fmt.Println(phi)

	var bilanganPositif uint8 = 89
	bilanganPositif = 100 // fixed the problem by assigning the value to the variable
	fmt.Println(bilanganPositif)

	//array
	var arr [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(arr)

	for a := range arr {
		fmt.Println(a)
	}

	//slice
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice[2])
	fmt.Println(len(slice))

	//map
	var person = map[string]string{
		"name":    "Ilham",
		"address": "Jakarta",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	if val, isExist := person["name"]; isExist {
		fmt.Println(val)
	}

	//function
	fmt.Println(add(2, 1))

	defer func() {
		fmt.Println("Ini dijalankan terakhir")
	}()

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Terjadi panic", r)
		}
	}()

	_, err := bagi(10, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func add(a int, b int) int {
	return a + b
}

func bagi(a, b int) (int, error) {
	// if b == 0 {
	// 	return 0, errors.New("Tidak bisa membagi dengan 0")
	// }
	return a / b, nil
}
