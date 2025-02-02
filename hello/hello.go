package main

import (

	"fmt"
	"log"
	"example.com/greetings"
)


func main(){
	//properties of the logger:
	/*
		Log prefix: "Greetings: "
		Flag: 0, disables printing the time, source file and line numbers
		
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(message)
}
