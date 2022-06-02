package main

import "fmt"

func main() {
	// membuat alias
	// membuat type ari baru dari type data yang sudah ada
	type nama string
	type status bool

	var firstName nama = "Muhammad"
	statusKawin := true
	fmt.Println(firstName)
	fmt.Println(statusKawin)
}
