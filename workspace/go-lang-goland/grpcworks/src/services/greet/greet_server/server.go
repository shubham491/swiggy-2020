package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"services/greet/greetpb"
	"strings"
)

type Server struct {

}


func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
fmt.Println("Function called... ")

fname := req.GetGreeting().GetFistName()
lname := req.GetGreeting().GetLastName()

result := "Hello : "+fname+ ", "+lname

res := &greetpb.GreetResponse{
Result:result,
}
return res,nil

}



func (*Server) GreetFullName(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetFullNameResponse, error)  {
	fmt.Println("GreetFullName called...")

	firstName := req.GetGreeting().GetFistName()
	lastName := req.GetGreeting().GetLastName()

	res := &greetpb.GreetFullNameResponse{
		Greeting: &greetpb.Greeting{
			FistName: strings.ToUpper(firstName),
			LastName: strings.ToUpper(lastName),
		},
	}
	return res, nil
}

func main() {
	fmt.Println("Hello from server")

	lis,err:=net.Listen("tcp","0.0.0.0:50051")
	if err!=nil{
		log.Fatal("sorry failed to load server",err)

	}
	s:=grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &Server{})

	if s.Serve(lis); err!=nil{
		log.Fatal("sorry failed to load server",err)
	}


}
