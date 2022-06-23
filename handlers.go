package main

import (
	"bufio"
	"fmt"
)

func handleinput(s []string, reader *bufio.Reader) {
	switch s[0] {
	case "ADDUSER":
		createUser(s, reader)
	case "PRINTUSERS":
		fmt.Println(userArr)
	default:
		fmt.Println("incorrect message")
	}
}
