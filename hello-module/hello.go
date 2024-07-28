package main

import (
	"fmt"

	"modouaicha023/greetings"
)

func main(){
	message := greetings.Hello("modou")
	fmt.Println(message)
}
