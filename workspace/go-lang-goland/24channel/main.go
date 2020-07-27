package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	DoMyWork("Hemant")
	//	wg.Done()
	//}()
	//wg.Wait()

	c:=make(chan string)
	go DoMyWork("priya",c)

	//for {
	//	msg,open := <-c
	//	if !open{
	//		break
	//	}
	//	fmt.Println(msg)
	//}

	//or

	for msg:=range c{
		fmt.Println(msg)
	}

	//go DoMyWork("Hemant")
	//
	//time.Sleep(time.Second*3)
}

func DoMyWork(name string, c chan string){
	for i:=1;i<=4;i++ {
		c<- "Doing work of "+ name+" executed "+strconv.Itoa(i)
		time.Sleep(time.Millisecond*500)
	}
	close(c)

}
