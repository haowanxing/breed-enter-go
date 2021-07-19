package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: nil,
		Port: 37540,
	})
	if err != nil {
		showErrorAndExit(err)
	}
	done := make(chan interface{})
	go sendMessageLoop(conn, done)
	recv := make([]byte, 14)
	n, _, err := conn.ReadFromUDP(recv)
	done <- 1
	if err != nil {
		showErrorAndExit(err)
	}
	log.Println(string(recv[:n]))
}

func sendMessageLoop(conn *net.UDPConn, closed chan interface{}) {
	tick := time.NewTicker(time.Millisecond * 300)
	log.Println("Start sending Abort command per 0.3s")
	for {
		select {
		case <-closed:
			log.Println("Stop sending.")
			return
		case <-tick.C:
			_, _ = conn.WriteToUDP([]byte("BREED:ABORT"), &net.UDPAddr{
				IP:   net.ParseIP("255.255.255.255"),
				Port: 37541,
			})
		}
	}
}

func showErrorAndExit(err error) {
	fmt.Println(err)
	fmt.Println("----------按任意键退出---------")
	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)
	os.Exit(1)
}
