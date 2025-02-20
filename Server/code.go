package main

// Server to send request to
// used this code as 3rd party server that the proxy server will contact on behalf of the client
import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func handler(conn net.Conn) {
	defer conn.Close()
	msg := []byte("message from the server, final destination server")
	_, err := conn.Write(msg)
	log.Print("Welcome, from inside the server")
	if err != nil {
		log.Printf("Failed to write a message: %v", err)
		return
	}
}

// experimental code to practice for the first time sending a response back to the client
func handle_mult(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err2 := conn.Read(buffer)
	if err2 != nil {
		log.Printf("Error in reading buffer: %v", err2)
		return
	}
	// read the buffer before anything
	log.Print(string(buffer[n:]))
	time.Sleep(1 * time.Second)
	resp := string(buffer[:n])
	num, err3 := strconv.Atoi(resp)
	if err3 != nil {
		log.Fatalf("Error in conversion to integer: %v", err3)
		return
	}
	log.Print(num)
	time.Sleep(2 * time.Second)
	res := strconv.Itoa(2 * num)
	send := []byte(res)
	_, err := conn.Write(send)
	log.Printf("Multiply by 2 server. welcome")
	if err != nil {
		log.Printf("Error in sending the math result: %v", err)
		return
	}
}

func Server() {

	ln, err := net.Listen("tcp", "*******")
	if err != nil {
		log.Fatalf("Failed to listent on port **********: %v \n", err)
	}
	log.Print("Welcome to my server listening on port *******, host: *********")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error in accepting the connection: %v", err)
			return
		}
		go handler(conn)

		fmt.Printf("IP: %s \n", conn.RemoteAddr())
	}

}

func main() {
	Server()
}
