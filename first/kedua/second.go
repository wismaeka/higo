package kedua

import (
	"fmt"
	"log"
	"strconv"
)

// InputNumber accept numbr
func InputNumber() string {

	fmt.Print("Enter a number: \n")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
		return "error"
	}
	result, err := strconv.Atoi(input)

	if result < 1 || result > 10 {
		return "Angka yang dimasukan di luar range"
	} else {
		return "Angka yang dimasukkan sudah didalam range"
	}
}
