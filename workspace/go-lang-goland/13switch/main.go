package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow:=rand.Intn(6)+1
	results:=""
	switch dow{
	case 1:
		results="Its sunday"
		fallthrough
	case 2:
			results="Its saturday"
	default:
		results="Its week day"


	}


	fmt.Println(dow)
	fmt.Println(results)

}
