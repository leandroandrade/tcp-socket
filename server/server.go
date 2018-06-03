package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("server waiting connections")

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	defer listener.Close()

	fmt.Println("connection accepted")
	for {
		message, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		fmt.Println("message received:", string(message))

		newMessage := strings.ToUpper(message)
		connection.Write([]byte(newMessage + "\n"))

	}
}
