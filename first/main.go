package main

import (
	"fmt"
	"log"
)

func main() {
	//fmt.Println("Hi, you can call me", os.Args[1])

	fmt.Print("Enter your name: \n")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("hi, your name is:", input)

}
