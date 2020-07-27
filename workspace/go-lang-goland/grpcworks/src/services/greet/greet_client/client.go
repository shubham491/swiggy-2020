package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"services/greet/greetpb"
)

func main() {
	fmt.Println("Hi i a,m client")
	conn,err:=grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err!=nil{
		log.Fatal("sorry failed to talk to server",err)
	}
	defer conn.Close()

	c:=greetpb.NewGreetServiceClient(conn)

	callGreet(c)
	callGreetFullName(c)

}

func callGreet(c greetpb.GreetServiceClient){
	fmt.Println("in Greet Function..")
	req:=&greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
		FistName :"shubham",
		LastName: "das",
		},
	}

	res,err:=c.Greet(context.Background(), req)

	if err!=nil{
		log.Fatal("Error while calling greet",err)
	}
	log.Println("Response from server",res.Result)

}

func  callGreetFullName(c greetpb.GreetServiceClient)  {
	fmt.Println("In callGreetFullName()... ")

	req:= &greetpb.GreetRequest {
		Greeting: &greetpb.Greeting {
			FistName: "Shubham",
			LastName: "Das",
		},
	}
	res, err := c.GreetFullName(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	log.Println("Response From the GreetFullName:", res.Greeting.FistName + " " + res.Greeting.LastName)
}
