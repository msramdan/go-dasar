package main

import (
	"fmt"
)

func main() {
	name := "ramdan"

	switch name {
	case "muhammad":
		fmt.Println("Hello, ", name)
	case "saeful":
		fmt.Println("Hai, ", name)
	case "ramdan":
		fmt.Println("Hey, ", name)
	}

	// short hand
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu panjang")
	case true:
		fmt.Println("Nama Sudah benar")
	}

}
