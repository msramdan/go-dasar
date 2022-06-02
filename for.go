package main

import "fmt"

func main() {
	// example 1
	angka := 1
	for angka <= 10 {
		fmt.Println("Perulangan ke : ", angka)
		angka++
	}

	// example 2
	slice := []string{"muhammad", "saeful", "ramdan", "kira", "anisa", "rahmawati"}

	for i := 0; i < len(slice); {
		fmt.Println("Perulangan ke : ", slice[i])
		i++
	}

}
