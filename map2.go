package main

import "fmt"

// pake interface untuk isi data map nya bisa berbeda type
func main() {
	book := make(map[string]interface{})
	book["title"] = "Belajar Pemrograman"
	book["pengarang"] = "Ramdan"
	book["harga"] = 2000

	fmt.Println(book)
	fmt.Println(len(book))

	// hapus map
	delete(book, "harga")

	fmt.Println(book)
	fmt.Println(len(book))
}
