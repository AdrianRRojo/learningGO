package greeings

import "fmt"


//func nameOfFunc(param type) returnType{
func Hello(name string) string {

	message := fmt.Sprintf("HI, %v. Welcome", name)
	// := declares a value and it's type in 1 line
	/*
		The long way of writing this same line is:
		var message string
		message = fmt.Sprintf("HI......", name)
		$v is filled in using the name variable
	*/
	return message
}
