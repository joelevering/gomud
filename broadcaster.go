package main

import "log"

func Broadcaster() {
	for {
		select {
		case msg := <-Messages:
			for _, cli := range GameState.clients {
				cli.sendMsg(msg)
			}
		case cli := <-Entering:
			log.Print("User logged in: " + cli.name)

			GameState.clients[cli.name] = cli

			SetCurrentRoom(cli, GameState.defaultRoom)

			DescribeCurrentRoom(*cli)

			go sendMsg(cli.name + " has arrived!")
		case cli := <-Leaving:
			log.Print("User logged out: " + cli.name)

			delete(GameState.clients, cli.name)
			close(cli.channel)

			go sendMsg(cli.name + " has left!")
		}
	}
}
