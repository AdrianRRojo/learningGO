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


	names := []string{"Adrian", "Colby", "Lily"}

	messages, err := greetings.MultipleHellos(names)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(messages)
}


