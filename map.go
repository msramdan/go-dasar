package main

import "fmt"

func main() {
	// tipe data map ini seperti object
	person := map[string]string{
		"nama":      "Muhammad Saeful Ramdan",
		"umur":      "26",
		"ttl":       "Sukabumi, 23 Maret 1996",
		"namaIstri": "Anisa Rahmawati",
		"alamat":    "Bogor",
	}
	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["umur"])

	// panjang map
	fmt.Println(len(person))

	// hapus data di map dengan key
	delete(person, "nama")

	fmt.Println(person)
}
