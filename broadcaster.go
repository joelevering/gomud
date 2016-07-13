package main

import "log"

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			for _, cli := range GameState.clients {
				cli.sendMsg(msg)
			}
		case cli := <-entering:
			log.Print("User logged in: " + cli.name)

			GameState.clients[cli.name] = cli

			cli.sendMsg(GameState.defaultRoom.name)
			cli.sendMsg(GameState.defaultRoom.desc)

			ListClients(cli)

			go sendMsg(cli.name + " has arrived!")
		case cli := <-leaving:
			log.Print("User logged out: " + cli.name)

			delete(GameState.clients, cli.name)
			close(cli.channel)

			go sendMsg(cli.name + " has left!")
		}
	}
}
