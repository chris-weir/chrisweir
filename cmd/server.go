package main

import "github.com/chris-weir/chrisweir/server"

func main() {
	server := server.GetServer()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
