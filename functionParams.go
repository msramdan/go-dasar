package main

import "fmt"

func main() {
	helloSay("Muhammad", "Ramdan")
}

func helloSay(firstName string, lastName string) {
	fmt.Println("Hello ", firstName+" dan", lastName)
}
