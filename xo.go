package main

func main() {

	lsn := NewListener("", "8082", "auth")
	lsn.runListener()
}
