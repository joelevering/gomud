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
  name := cli.GetName()

  log.Printf("User logged in: %s", name)

	gk.state.Clients[name] = cli

  cli.Character.SetSpawn(gk.state.DefaultRoom)
	cli.Spawn()

	cli.Look()

  go gk.broadcast(fmt.Sprintf("%s has logged in!", name))
}

func (gk *Gatekeeper) logOut(cli *client.Client) {
  name := cli.GetName()
	cli.LeaveRoom("")

	log.Printf("User logged out: %s", name)

	delete(gk.state.Clients, name)
	close(cli.Channel)

	go gk.broadcast(fmt.Sprintf("%s has logged out!", name))
}
