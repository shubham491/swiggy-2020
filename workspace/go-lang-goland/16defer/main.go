package main

import "fmt"

func main() {
	//defer executes in reverse order like push in stack then poping and statment have to be placed before the panic line
	defer fmt.Println("Connection close...")
	fmt.Println("Connection open...")
	//panic("Some ERror")
	myFunc()
	defer fmt.Println("stat1")
	defer fmt.Println("stat2")
	fmt.Println("Do the work")
}
func myFunc(){
	defer fmt.Println("im defer in myfunc")
	fmt.Println("im not defer in myfunc")
}
