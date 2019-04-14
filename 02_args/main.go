package main

import "fmt"

var lala = "ssu"

func main() {
	var name string = "Sasha"
	var age int = 37
	var isCool = true
	isCool = false
	kuku := "Putu"
	size := 3.72
	name, email := "Brad", "sasha@brad.tr"
	fmt.Println(lala, kuku)
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
