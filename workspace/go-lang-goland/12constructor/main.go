package main

import (
	"fmt"
	"os"
)

type Trade struct{
	CompanyName string
	Volume int
	Price float64
	Buy bool

}

func (t *Trade) Value() float64{
	value:=float64(t.Volume)*t.Price;
	if t.Buy{
		fmt.Println("You are buying with value ",value)
	} else {
		fmt.Println("You are selling for value ", value)
	}
	return value
}

func NewTrade(companyName string, volume int, price float64, buy bool)(*Trade, error){
	if companyName==""{
		return nil,fmt.Errorf("Sorry comapany name cant be nil");
	}
	if price<0{
		return nil,fmt.Errorf("Sorry price cant be -ve");
	}
	if volume<=0{
		return nil,fmt.Errorf("Sorry you have to trade at least one stock");
	}

	trade:=&Trade{
		CompanyName: companyName,
		Price: price,
		Volume: volume,
		Buy: buy,
	}
	return trade, nil
}

func main() {
	comp,err := NewTrade("Swiggy",10,56.5,true);
	if err!=nil{
		fmt.Println("Error creating trade",err);
		os.Exit(1);
	}
	fmt.Println(comp)
	comp.Value()

}
