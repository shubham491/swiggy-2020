package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Scanln(&s)
	//fmt.Println(s)
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	s,_:=reader.ReadString('\n')
	fmt.Println("Your name is:",s)
	fmt.Println("Enter your monthly salary:")
	salary,_:=reader.ReadString('\n')
	fmt.Println("Enter no. of months since you joined:")
	months,_:=reader.ReadString('\n')
	salary1,_:=strconv.ParseFloat(strings.TrimSpace(salary),64)
	months1,_:=strconv.ParseInt(strings.TrimSpace(months),10, 64)
	totalSalary:=salary1*float64(months1)
	fmt.Println("Total salary is: ",totalSalary)

}
