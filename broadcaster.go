package main

import "log"

func Broadcaster() {
	for {
		select {
		case cli := <-Entering:
			log.Print("User logged in: " + cli.name)

			GameState.clients[cli.name] = cli

			SetCurrentRoom(cli, GameState.defaultRoom)

			DescribeCurrentRoom(*cli)

			go broadcast(cli.name + " has logged in!")
		case cli := <-Leaving:
			RemoveClientFromRoom(cli)

			log.Print("User logged out: " + cli.name)

			delete(GameState.clients, cli.name)
			close(cli.channel)

			go broadcast(cli.name + " has logged out!")
		}
	}
}

func broadcast(msg string) {
	for _, cli := range GameState.clients {
		cli.sendMsg(msg)
	}
}
