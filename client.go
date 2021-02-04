package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECTION := arguments[1]
	c, err := net.Dial("tcp", CONNECTION)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//READ THE USER INPUT
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("CLIENT : ")
		text, _ := reader.ReadString('\n')
		//user input is sent to the TCP server over the network
		fmt.Fprintf(c, text+"\n")

		//read the TCP server’s response
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		// terminate when you send the STOP command to the TCP server
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

// go run client.go 127.0.0.1:8080
//go mod init Tcp_connection
