package main

import "fmt"

var isConnected bool = false


func main() {
	fmt.Println("Connection open: ", isConnected)
	businessLogic()
	fmt.Println("Connection closed: ", isConnected)
}

func businessLogic(){
	defer disConnect()
	connect()
	fmt.Println("DB connection businesslogin status: ",isConnected)
	fmt.Println("BUSINESS logic goes here")
}
func connect(){
	isConnected=true
	fmt.Println("DB Connection is open")
}

func disConnect(){
	isConnected=false
	fmt.Println("DB Connection is open")
}

