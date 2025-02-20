package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

// used this code as the client that sends and recieve data from 3rd party server though the proxy server
// experimental code block
func Call(addr string, num int) (bool, error) {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return false, fmt.Errorf("Error: %v", err)
	}

	defer conn.Close()
	go func() {
		resl := strconv.Itoa(num)
		_, err := conn.Write([]byte(resl))
		if err != nil {
			log.Fatalf("Could not write the number to the server: %v", err)
			return
		}

	}()
	time.Sleep(1 * time.Second)
	log.Printf("Connection made to the server: %s from address: %s", addr, conn.LocalAddr())
	//	buffer := make([]byte, 1024)
	buffer2 := make([]byte, 1024)
	//	n, erri := conn.Read(buffer)
	//	if erri != nil {
	//		return false, fmt.Errorf("Error: %v", erri)
	//	}
	time.Sleep(1 * time.Second)
	n2, err2 := conn.Read(buffer2)
	if err2 != nil {
		return false, fmt.Errorf("Error 2nd buffer: %v", err2)
	}
	//	log.Print(string(buffer[:n]))
	log.Print(string(buffer2[:n2]))
	return true, nil
}

// experimental code block end //
func call_proxy() {
	conn, err := net.Dial("tcp", "***********")
	if err != nil {
		fmt.Printf("Error :%v", err)
		return
	}
	sending := "********"
	_, err3 := conn.Write([]byte(sending))
	if err3 != nil {
		fmt.Printf("Error :%v", err3)
		return
	}
	defer conn.Close()
	time.Sleep(1 * time.Second)
	send2 := "*********" // the IP and port addresses of the server of which the client listens
	_, err6 := conn.Write([]byte(send2))
	if err6 != nil {
		fmt.Printf("Error: %v", err6)
		return
	}
	ln, err2 := net.Listen("tcp", "********")
	if err2 != nil {
		fmt.Printf("Error :%v", err2)
		return
	}

	buffer := make([]byte, 2048)
	for {
		conn2, err4 := ln.Accept()
		if err4 != nil {
			fmt.Printf("Error: %v", err4)
			return
		}
		n, err5 := conn2.Read(buffer)
		if err5 != nil {
			fmt.Printf("Error :%v", err5)
			return
		}

		fmt.Print(string(buffer[:n]))

	}

}

func main() {
	call_proxy()
}
