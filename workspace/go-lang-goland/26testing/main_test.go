// +build unit

package main

import (
	"io/ioutil"
	"testing"
)

type AddResult struct{
	x int
	y int
	expected int
}
var AddResults=[] AddResult{
	{10,20,30},
	{
		100,200,300},

}
func TestAdd(t *testing.T) {
	for _,test:=range AddResults{
		result:=Add(test.x,test.y)
		if(result != test.expected){
			t.Fatal("Expected result not found")
		}

	}
}


func TestReadFile(t *testing.T){
	data,err:=ioutil.ReadFile("testdata/test.data")
	if err!=nil{
		t.Fatal("File not found")
	}
	expected:="Hello World"
	if string(data) != expected{
		t.Fatal("Data mismatch from file")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i:=0;i<b.N;i++{
		Add(i,i)
	}
}