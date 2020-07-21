package main

import "fmt"

type Dog struct {
	//declaring variables starting with capital letter are public.
	Breed string
	Weight int
}

func main() {
	dog1:=Dog{Breed: "labra",Weight:23}
	fmt.Println(dog1)

	fmt.Printf("Dog breed: %v, dog weight: %v\n",dog1.Breed,dog1.Weight)



}
