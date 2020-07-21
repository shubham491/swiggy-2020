package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal interface{
	Eat() string
	Sleep() string
	Breathe() string
	Details()

}
func (e Emp) Eat() string{
	return "Employee is eating.."
}
func (e Emp) Sleep() string{
	return "Employee is sleeping..."
}
func (e Emp) Breathe() string{
	return "Employee is breathing..."
}

type Address struct {
	house_no string
	street string
	city string
	pin string
}

type Emp struct {
	empid string
	emp_name string
	salary float64
	Address

}

func (e Emp) Details() {
	fmt.Println("Employee Id: ", e.empid)
	fmt.Println("Emplyee Name: ", e.emp_name)
	fmt.Printf("Employee's Salary: %f\n", e.salary)
	fmt.Printf("Address: %v,%v, %v - %v \n\n", e.house_no,e.street,e.city,e.pin)
}

func main() {
	employees := make([] Animal, 0, 5)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the employee detail entry portal:\n")
	for true {
		fmt.Println("Enter your choice")
		fmt.Printf("C : Continue to enter details of %vth employee.\n",len(employees)+1)
		fmt.Println("Q: Quit")
		s,_:=reader.ReadString('\n')
		option:=strings.TrimSpace(s)
		if option=="Q"{
			break;
		} else {
			fmt.Println("Enter employee ID")
			s,_:=reader.ReadString('\n')
			id := strings.TrimSpace(s)
			fmt.Println("Enter employee name")
			s,_=reader.ReadString('\n')
			name := strings.TrimSpace(s)
			fmt.Println("Enter employee's salary")
			s,_=reader.ReadString('\n')
			salary,_:= strconv.ParseFloat(strings.TrimSpace(s),64)
			fmt.Println("Enter employee's house no.")
			s,_=reader.ReadString('\n')
			houseno := strings.TrimSpace(s)
			fmt.Println("Enter street name the employee lives in")
			s,_=reader.ReadString('\n')
			street := strings.TrimSpace(s)
			fmt.Println("Enter employee's city")
			s,_=reader.ReadString('\n')
			city := strings.TrimSpace(s)
			fmt.Println("Enter employee's pincode")
			s,_=reader.ReadString('\n')
			pincode := strings.TrimSpace(s)
			address1:= Address{
				houseno,
				street,
				city,
				pincode,
			}
			new_employee:= Animal(Emp{
				id,
				name,
				salary,
				address1,
			})

			employees= append(employees, new_employee)
			fmt.Printf("Total %v employees details filled.\n",len(employees))


		}

	}
	fmt.Println("Employees Details:\n")
	//fmt.Println(employees)
	if len(employees) == 0 {
		fmt.Println("No employee record added")
	} else {
		for i := range employees {
			//fmt.Println(i)
			fmt.Printf("Employee No. %v:\n", i+1)
			employees[i].Details()
		}
	}



}