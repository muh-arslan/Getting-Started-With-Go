package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//message, err := greetings.Hello("Arslan")
	names := []string{"Arslan", "Mujeeb", "Oreed", "Talha", "Uzair"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
