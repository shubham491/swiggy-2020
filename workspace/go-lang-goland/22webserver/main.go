package main

import (
	"fmt"
	"net/http"
)

type Hello struct {

}

func (h Hello) ServeHTTP(writer http.ResponseWriter,req *http.Request){
	fmt.Fprint(writer,"<h2>Hello welcome from Golang</h2>")
}

func main() {
	var h Hello
	err:=http.ListenAndServe("localhost:7788", h)
	checkError(err)
}

func checkError(err error) {
	if err != nil{
		panic(err)
	}
}
