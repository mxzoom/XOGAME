package main

import (
	"bufio"
	"time"
)

type user struct { //USER
	id             int
	name           string
	reader         *bufio.Reader //?????? пока не знаю, надо ли
	status         bool          //alive or not
	lastAliveCheck time.Time
}

var users = make(map[int]user) //STORE for USERS, [id]user
