package main

import (
	"log"
	"net"
	"time"
)

const Port = "1919"

var clients = make(map[string]Client)
var messages = make(chan string)

type Client struct {
	channel chan<- string
	name    string
}

func (cli Client) sendMsg(msg string) {
	stamp := time.Now().Format(time.Kitchen)
	cli.channel <- stamp + " " + msg
}

func sendMsg(msg string) {
	messages <- msg
}

func clientEnters(cli Client) {
	entering <- cli
}

func clientLeft(cli Client) {
	leaving <- cli
}

func localIp() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		ipnet := addr.(*net.IPNet)
		if !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return ""
}

func main() {
	host := localIp() + ":" + Port
	log.Print("Hosting on: " + host)
	listener, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, clients)
	}
}
