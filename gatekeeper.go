package main

import (
	"fmt"
	"log"

	"github.com/joelevering/gomud/client"
)

type Gatekeeper struct {
	entering <-chan *client.Client
	leaving  <-chan *client.Client
	state    *GameState
}

func (gk *Gatekeeper) KeepTheGate() {
	for {
		select {
		case cli := <-gk.entering:
			gk.logIn(cli)
		case cli := <-gk.leaving:
			gk.logOut(cli)
		}
	}
}

func (gk *Gatekeeper) broadcast(msg string) {
	for _, cli := range gk.state.Clients {
		cli.SendMsg(msg)
	}
}

func (gk *Gatekeeper) logIn(cli *client.Client) {
	log.Printf("User logged in: %s", cli.Name)

	gk.state.Clients[cli.Name] = cli

  cli.SetSpawn(gk.state.DefaultRoom)
  cli.Spawn()

	cli.Look()

	go gk.broadcast(fmt.Sprintf("%s has logged in!", cli.Name))
}

func (gk *Gatekeeper) logOut(cli *client.Client) {
	cli.LeaveRoom("")

	log.Printf("User logged out: %s", cli.Name)

	delete(gk.state.Clients, cli.Name)
	close(cli.Channel)

	go gk.broadcast(fmt.Sprintf("%s has logged out!", cli.Name))
}
