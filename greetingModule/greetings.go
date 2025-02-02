package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

//func nameOfFunc(param type) returnType{
func Hello(name string) (string, error) {
	
	if name == "" {
		return "", errors.New("No name provided")
	}
	message := fmt.Sprintf(randomFormat(), name)
	// := declares a value and it's type in 1 line
	/*
		The long way of writing this same line is:
		var message string
		message = fmt.Sprintf("HI......", name)
		$v is filled in using the name variable
	*/
	return message, nil
}

func randomFormat() string {
		// This is a slice, it will change it's size dynamically as you add or remove items
	formats := []string{
		"Hi, %v. Welcome!",
		"Nice to meet you, %v!",
		"YO wazzzzzzup %v!",
	}


		// return a random message.
		// random index for the slice of formats
	return formats[rand.Intn(len(formats))]


}
	/* 
		
	*/
