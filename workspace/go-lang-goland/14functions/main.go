package main

import "fmt"

func main() {
	DoSomeWork()

	fmt.Println(AddValues(1,3))

	fmt.Println(AddAllValues(1,2,3,4,5,6,67))
	name, leng := FullNmae("shub","das")
	fmt.Println("Name: ",name,"\nlength: ",leng)

	name, leng = FullNameNakedReturn("shub","das")
	fmt.Println("Name: ",name,"\nlength: ",leng)
}
func DoSomeWork(){
	fmt.Println("Doing work")
}
func AddValues(value1 int, value2 int) int{
	return value1+value2
}

func AddAllValues(values ... int)int{
	sum:=0
	for i:=range values{
		sum += values[i]
	}
	return sum
}

func FullNmae (f,l string) (string,int){
	full := f+" "+l
	length:=len(full)
	return full,length
}

func FullNameNakedReturn (f,l string) (full string,length int){
	full = f+" "+l
	length=len(full)
	return full,length
}