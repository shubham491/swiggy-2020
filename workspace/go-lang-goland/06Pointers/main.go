package main

import "fmt"

func main() {
	v := 44
	var p *int = &v
	if p!=nil {
		fmt.Println("Value at p: ", *p)
	} else
	{
		fmt.Println("sorry.. p is nil")
	}
	*p = *p/2;
	if p!=nil {
		fmt.Println("Value at p: ", *p)
	} else
	{
		fmt.Println("sorry.. p is nil")
	}

}
