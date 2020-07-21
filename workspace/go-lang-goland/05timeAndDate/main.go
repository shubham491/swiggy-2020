package main

import (
	"fmt"
	"time"
)

func main() {
	//t:=time.Date(2020,time.July,20,10,10,10,0,time.UTC)
	t:=time.Now()
	fmt.Printf("Time is: %s\n",t)
	fmt.Println(t.Weekday())
	tommorow:=t.AddDate(0,0,1)
	fmt.Printf("Tommorow is: %s\n",tommorow)
	fmt.Printf("Tommorow is: %s\n",tommorow.Format("02/01/2006"))
}
