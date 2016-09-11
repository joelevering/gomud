package main

import (
	"fmt"
	"log"
)

func Gatekeeper(entering, leaving <-chan *Client) {
	for {
		select {
		case cli := <-entering:
			logIn(cli)
		case cli := <-leaving:
			logOut(cli)
		}
	}
}

func broadcast(msg string) {
	for _, cli := range GameState.Clients {
		cli.SendMsg(msg)
	}
}

func logIn(cli *Client) {
	log.Printf("User logged in: %s", cli.Name)

	GameState.Clients[cli.Name] = cli

	cli.EnterRoom(GameState.DefaultRoom)

	cli.Look()

	go broadcast(fmt.Sprintf("%s has logged in!", cli.Name))
}

func logOut(cli *Client) {
	cli.LeaveRoom("")

	log.Printf("User logged out: %s", cli.Name)

	delete(GameState.Clients, cli.Name)
	close(cli.Channel)

	go broadcast(fmt.Sprintf("%s has logged out!", cli.Name))
}
