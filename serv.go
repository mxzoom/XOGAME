package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Listener struct {
	Ip   string
	Port string
	Auth string
}

func NewListener(ip, port string, auth string) *Listener {
	return &Listener{Ip: ip, Port: port, Auth: auth}
}

func (l *Listener) runListener() {
	netadr := l.Ip + ":" + l.Port
	fmt.Println("start listening on -", netadr)
	listener, err := net.Listen("tcp4", netadr)
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func(conn net.Conn) {
			reader := bufio.NewReader(conn)
			for {
				msg, err := reader.ReadString('\n')
				if err != nil {
					conn.Close()
					return
				}
				msg = strings.Replace(msg, "\r\n", "", -1)
				msgresult := strings.Split(msg, ":")
				handleinput(msgresult, reader)
			}
		}(conn)
	}
}
