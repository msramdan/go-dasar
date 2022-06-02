package main

import "fmt"

func main() {
	// break
	for i := 0; i < 10; {
		fmt.Println("Perulangan ke : ", i)
		if i == 5 {
			break
		}
		i++
	}

}
