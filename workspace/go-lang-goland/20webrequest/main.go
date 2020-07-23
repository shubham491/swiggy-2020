package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url:="https://jsonplaceholder.typicode.com/posts"
	resp,err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Content type %T \n",resp)

	fmt.Println("Status Code:",resp.StatusCode)

	bytes, err :=ioutil.ReadAll(resp.Body)

	if err != nil{
		panic(err)
	}
	content :=string(bytes)
	fmt.Println(content)
}
