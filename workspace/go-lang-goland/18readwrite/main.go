package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content :="Hello WORLD FROM GO"
	file,err:=os.Create("./sample.txt")
	checkError(err)

	defer file.Close()

	ln,err :=io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Saved %v chars to file",ln)


	bytes:=[] byte(content)

	err=ioutil.WriteFile("./inbytes.txt",bytes,0644)
	checkError(err)
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}