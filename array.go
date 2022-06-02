package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "muhammad"
	nama[1] = "saeful"
	nama[2] = "ramdan"
	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	// array + isi datanya

	var value = [3]string{
		"10",
		"11",
		"12",
	}
	fmt.Println(value)
	fmt.Println(value[1])

	// cek panjang array
	fmt.Println(len(value))
}
