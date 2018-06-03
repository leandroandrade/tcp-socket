package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	for {
		read := bufio.NewReader(os.Stdin)

		fmt.Print("message to send: ")

		text, err := read.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		fmt.Fprintf(connection, text+"\n")
		response, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		fmt.Println("response from server:", response)

	}

}
