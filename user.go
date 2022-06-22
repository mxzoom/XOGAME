package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type user struct {
	id             int
	name           string
	reader         *bufio.Reader //?????? пока не знаю, надо ли
	status         bool          //alive or not
	lastAliveCheck time.Time
}

var userArr []user

func createUser(s []string, reader *bufio.Reader) { //ADDUSER:NAME:TODO??
	var m sync.Mutex
	name := s[1]
	m.Lock()
start:
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(99999)

	for _, elem := range userArr {
		if elem.id == id {
			fmt.Println("Gen new ID")
			goto start
		}
		if name == elem.name {
			name = name + "*"
			fmt.Println("!")
			goto start
		}
	}
	user := user{id: id, name: name, reader: reader, status: true, lastAliveCheck: time.Now()}
	userArr = append(userArr, user)
	fmt.Println("user created", name)
	m.Unlock()
}
