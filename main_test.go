package main

import (
	"net"
	"testing"
)

func Test_sendMessageLoop(t *testing.T) {
	type args struct {
		conn   *net.UDPConn
		closed chan interface{}
	}
	conn, _ := net.ListenUDP("udp", &net.UDPAddr{Port: 37540})
	tests := []struct {
		name string
		args args
	}{
		{name: "Test01", args: args{
			conn:   conn,
			closed: make(chan interface{}),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
