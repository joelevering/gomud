package main

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			for _, cli := range clients {
				cli.sendMsg(msg)
			}
		case cli := <-entering:
			clients[cli.name] = cli
		case cli := <-leaving:
			delete(clients, cli.name)
			close(cli.channel)
		}
	}
}
