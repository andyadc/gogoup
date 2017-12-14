package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//single()
	//multi()
}

/*
多并发执行
*/
func multi() {
	service := `:8888`
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			continue
		}
		go handleSimpleClient(conn)
	}
}

func handleSimpleClient(conn net.Conn) {
	defer conn.Close()
	sleep1s()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack

	defer conn.Close()
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == `timestamp` {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // clear last read content
	}
}

/*
单任务
*/
func single() {
	service := `:8888`
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		sleep1s()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func sleep1s() {
	time.Sleep(1 * time.Second)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
