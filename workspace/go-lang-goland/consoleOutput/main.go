package main
import (
	"fmt"
)
func main() {
	str1:="This is String one"
	str2:="This is second string"
	str3:="third string"
	aNumber := 33
	isTrue:=true
	strLength, _ :=fmt.Println(str1, str2, str3);
	fmt.Println(strLength);
	fmt.Printf("Value of aNumber %v\n",aNumber);
	fmt.Printf("Value of Boolean %v\n",isTrue);
	fmt.Printf("Value of aNumber in float %.2f\n",float64(aNumber))

	fmt.Printf("Data types %T %T %T %T %T\n",
		str1,str2,str3,aNumber,isTrue);

	myString:=fmt.Sprintf("Data types %T %T %T %T %T",
		str1,str2,str3,aNumber,isTrue);

	fmt.Println(myString)
}
