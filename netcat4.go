// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 227.

// Netcat is a simple read/write client for TCP servers.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpconn := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println("go.", err)
			//log.Fatal(err)
		}
		fmt.Println("See you.")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	fmt.Println("wait for done.")
	tcpconn.CloseWrite()
	//conn.Close()
	<-done // wait for background goroutine to finish
}

//!-
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		//log.Fatal(err)
		log.Println("main.", err)
	}
	fmt.Println("out.")
}
