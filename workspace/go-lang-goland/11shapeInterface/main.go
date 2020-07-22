package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct{
	Length float64

}

func (s *Square) Area() float64{
	return s.Length*s.Length

}
func SumOfAreas(shapes []Shape) float64{
	total:=0.0
	for _,shape:=range shapes{
		total += shape.Area()
	}
	return total
}

type Circle struct{
	Radius float64
}

func (c *Circle) Area() float64{
	return math.Pi*c.Radius*c.Radius

}


func main() {
	s:=&Square{
		15,
	}
	fmt.Println(s.Area())
	c:=&Circle{
		3.5,
	}

	fmt.Println(c.Area())
	c1:=&Circle{
		3.5,
	}
	fmt.Println(c1.Area())
	shapes:=[] Shape{s,c,c1}
	result:=SumOfAreas(shapes)
	fmt.Println(result)
	
}
