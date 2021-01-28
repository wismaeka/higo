package main

import (
	"first/kedua"
	"fmt"
)

func main() {
	//fmt.Println("Hi, you can call me", os.Args[1])
	// fmt.Print("Enter your name: \n")
	// var input string
	// _, err := fmt.Scanln(&input)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	fmt.Println(kedua.InputNumber())
}
