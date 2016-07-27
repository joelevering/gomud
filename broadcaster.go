package main

import (
	"fmt"
	"log"
)

func Broadcaster() {
	for {
		select {
		case cli := <-Entering:
			log.Printf("User logged in: %s", cli.name)

			GameState.clients[cli.name] = cli

			SetCurrentRoom(cli, GameState.defaultRoom)

			cli.Look()

			go broadcast(fmt.Sprintf("%s has logged in!", cli.name))
		case cli := <-Leaving:
			RemoveClientFromRoom(cli, "")

			log.Print("User logged out: " + cli.name)

			delete(GameState.clients, cli.name)
			close(cli.channel)

			go broadcast(fmt.Sprintf("%s has logged out!", cli.name))
		}
	}
}

func broadcast(msg string) {
	for _, cli := range GameState.clients {
		cli.sendMsg(msg)
	}
}
