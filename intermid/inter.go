package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Expiremental code block //
type ResToClient struct {
	Status  int
	ResBody string
	Error   error
}

func handler(conn net.Conn) {
	defer conn.Close()
	msg := "Welcome to the proxy server, it is working"
	_, err := conn.Write([]byte(msg))
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}

// func proxy_server
func proxy_call(addr string) interface{} {
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		//log.Fatal("Error in dialling the server")
		r3 := ResToClient{
			Status: 501,
			Error:  fmt.Errorf("Error, from proxy_call function, :%v", err),
		}
		return r3
	}

	defer conn.Close()
	buffer := make([]byte, 2048)
	n, err1 := conn.Read(buffer)
	if err1 != nil {
		r1 := ResToClient{
			Status: 502,
			Error:  fmt.Errorf("Could not read the buffer: %v", err),
		}
		return r1
	}
	msg := string(buffer[:n])

	r := ResToClient{
		Status:  200,
		ResBody: msg,
	}
	return r

}

// Exp. Code Block end //

func SendRequest(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	buffer := make([]byte, 2048)
	n, err2 := conn.Read(buffer)
	if err2 != nil {
		fmt.Printf("Error in write buffer: %v", err2)
		return
	}

	fmt.Print(string(buffer[:n]))

}

func send(addr, addr2 string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	buffer := make([]byte, 2048)
	n, err2 := conn.Read(buffer)
	if err2 != nil {
		fmt.Printf("Error :%v", err2)
		return
	}

	// fmt.Print(string(buffer[:n]))
	conn2, err3 := net.Dial("tcp", addr2)
	if err3 != nil {
		fmt.Printf("Error :%v", err3)
		return
	}

	_, err4 := conn2.Write(buffer[:n])
	if err4 != nil {
		fmt.Printf("Error :%v", err4)
		return
	}

}

// The proxy server itself
func ProxyServer() {
	ln, err := net.Listen("tcp", "**********")
	if err != nil {
		log.Fatalf("Could not listen to upcomming requests: %v", err)
		return
	}
	buffer := make([]byte, 1024)
	buffer2 := make([]byte, 1024)
	log.Print("Welcome to the proxy server")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Could not accpet: %v", err)
			return
		}
		n, err2 := conn.Read(buffer)
		if err2 != nil {
			fmt.Printf("An error occurred when reading the buffer: %v", err)
			return
		}
		n2, err3 := conn.Read(buffer2)
		if err3 != nil {
			fmt.Printf("Error :%v", err3)
			return
		}

		go send(string(buffer[:n]), string(buffer2[:n2]))
	}
}

func main() {

	ProxyServer()

}
