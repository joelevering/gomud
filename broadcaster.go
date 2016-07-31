package main

import (
	"fmt"
	"log"
)

func Broadcaster() {
	for {
		select {
		case cli := <-Entering:
			log.Printf("User logged in: %s", cli.Name)

			GameState.Clients[cli.Name] = cli

			SetCurrentRoom(cli, GameState.DefaultRoom)

			cli.Look()

			go broadcast(fmt.Sprintf("%s has logged in!", cli.Name))
		case cli := <-Leaving:
			RemoveClientFromRoom(cli, "")

			log.Print("User logged out: " + cli.Name)

			delete(GameState.Clients, cli.Name)
			close(cli.Channel)

			go broadcast(fmt.Sprintf("%s has logged out!", cli.Name))
		}
	}
}

func broadcast(msg string) {
	for _, cli := range GameState.Clients {
		cli.SendMsg(msg)
	}
}
