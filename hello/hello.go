package main

import (

	"fmt"

	"example.com/greetings"
)


func main(){

	message := greetings.Hello("Adrian")
	fmt.Println(message)
}
