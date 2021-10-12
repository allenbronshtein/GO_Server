// Noa Abbo 208523514
// Allen Bronshtein 206228751

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, waitGroup *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	waitGroup.Done()
}
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	waitGroup := sync.WaitGroup{}
	for input.Scan() {
		waitGroup.Add(1)
		go echo(c, input.Text(), 5*time.Second, &waitGroup)
	}
	waitGroup.Wait()
	c.Close() // close quitely
}
func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
