package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func fib(n int) int {
	if n <= 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			fmt.Println("failed to read data, err:", err)
			return
		}

		fmt.Printf("Got : %s", bytes)

		line := fmt.Sprintf("%s", bytes)
		templine := strings.Trim(line, "\n")
		if len(templine) > 0 {
			num, err := strconv.Atoi(templine)

			fmt.Println("num is ", num)
			if err != nil {
				fmt.Println("failed to conver string to int, err:", err)
			}
			result := fib(num)
			fmt.Println("Result :", result)
			conn.Write([]byte(strconv.Itoa(result) + "\n"))
		}

	}
}

func main() {
	listener, err := net.Listen("tcp", ":25000")
	if err != nil {
		fmt.Println("Failed to create listener, err:", err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		go handleConnection(conn)
	}

}
