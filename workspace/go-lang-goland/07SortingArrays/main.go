package main

import (
	"fmt"
	"sort"
)

func main() {
	var names [5] string
	names[0] = "Rohan"
	names[1] = "Aman"
	names[2] = "Bhavi"
	names[3] = "Shashi"
	names[4] = "Varun"
	fmt.Println(names)


	var numbers = [] int{1,2,3}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	numbers = append(numbers,8)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	numbers= append(numbers[1:len(numbers)])
	fmt.Println(numbers)
	numbers = numbers[:len(numbers)-1]
	fmt.Println(numbers)
	ints := make([]int,5,5)
	ints[0]=0
	ints[1]=4
	ints[2]=56
	ints[3]=3
	ints[4]=4
	fmt.Println(ints)
	fmt.Println("Initial capacity",cap(ints))
	ints=append(ints,100)
	fmt.Println(ints)
	fmt.Println("Final capacity",cap(ints))
	sort.Ints(ints)
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)






}
