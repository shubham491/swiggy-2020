package main

import (
	"fmt"
	"os"
)


// simple HTTP error
type HttpError struct {
	status int
	method string
}

type SSLError struct {

	message string
	statusCode int
}

// HttpError implements Error method
func (httpError *HttpError) Error() string {
	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		httpError.method, httpError.status)
}

func (sslError *SSLError) Error() string {
	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		sslError.message, sslError.statusCode)
}

// mock function: sends HTTP request
func GetUserEmail(userId int) (string, error) {

	if(userId<100){
	return "shubham.das@gmail.com",nil
	}
	// request failed, return an error
	return "", &HttpError{403, "GET"}
}

func main() {

	// get user email address
	if email, err := GetUserEmail(101); err != nil {
		fmt.Println(err) // print error

		// if user is unauthorized, perform session cleaning
		if errVal := err.(*HttpError); errVal.status == 403 {
			fmt.Println("Performing session cleaning...")
		}

	} else {
		// do something with the `email`
		fmt.Println("User email is", email)
	}
}


func main1() {
	file, err := os.Open("sample.txt") // For read access.
	if err==nil{
	fmt.Println(file)
	} else {
		fmt.Println(err.Error())
	}


}
