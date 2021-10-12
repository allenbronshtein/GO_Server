// Noa Abbo 208523514
// Allen Bronshtein 206228751

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// TCP echo server.
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

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	wait_group := sync.WaitGroup{}
	defer func() {
		wait_group.Wait()
		c.Close()
	}()
	duration := 10 * time.Second
	timer := time.NewTimer(duration)
	str := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			str <- input.Text()
		}
		if input.Err() != nil {
			log.Println("no input for 10 seconds")
		}
	}()
	for {
		select {
		case input := <-str:
			timer.Reset(duration)
			wait_group.Add(1)
			go func() {
				echo(c, input, 1*time.Second)
				wait_group.Done()
			}()
		case <-timer.C:
			return
		}
	}
}
func main() {
	s, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
