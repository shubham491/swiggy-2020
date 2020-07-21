package main

import (
	"fmt"
	"math/big"
)

func main() {
	//i1,i2,i3:=10,20,30
	//intsum:=i1+i2+i3


	var b1,b2,b3,bitsum big.Float
	b1.SetFloat64(12.3)
	b2.SetFloat64(23.45)
	b3.SetFloat64(43.22)
	bitsum.Add(&b1, &b2).Add(&bitsum,&b3)
	fmt.Printf("Bigsum value:%.2f\n",&bitsum)

}
