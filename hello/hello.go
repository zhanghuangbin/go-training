package main

import (
	"fmt"
	"log"

	"github.com/zhanghuangbin/go-training/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	msg1, err1 := greetings.Hello("zhanghuangbin")

	if err1 != nil {
		log.Fatal(err1)
	} else {
		fmt.Println(msg1)
	}

	names := []string{"zhang", "huang", "bin"}
	msg2, err2 := greetings.Hellos(names)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(msg2)

}
