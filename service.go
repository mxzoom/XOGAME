package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func createUser(s []string, reader *bufio.Reader) { //ADDUSER:NAME:TODO??
	var m sync.Mutex
	name := s[1]
	m.Lock()
start:
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(99999)

	for key, elem := range users {
		if key == id {
			fmt.Println("Gen new ID")
			goto start
		}
		if name == elem.name {
			name = name + "*"
			goto start
		}
	}
	user := user{id: id, name: name, reader: reader, status: true, lastAliveCheck: time.Now()}
	users[id] = user
	fmt.Println("user created", name)
	m.Unlock()
}
