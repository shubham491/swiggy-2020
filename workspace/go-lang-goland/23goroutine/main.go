package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {

	urls:=[] string{
		"https://golang.org", //string
		"https://jsonplaceholder.typicode.com/posts",//json
		"https://httpbin.org/xml",//xml
	}

	var wg sync.WaitGroup //used so that program doesn't terminate before getting the response
	for _,url:=range urls{
		//go returnType(url)
		//or
		wg.Add(1)
		go func(url string){
			returnType(url)  //async
			wg.Done()
		}(url)
	}
	wg.Wait() //wait till delta becomes zero
	
}

func returnType(url string){
	resp,err:=http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	ctype:=resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n",url,ctype)

}
func checkError(err error) {
	if err != nil{
		panic(err)
	}
}
