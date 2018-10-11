package main

import (
	"fmt"
	"log"

	"github.com/joelevering/gomud/player"
)

type Gatekeeper struct {
	entering <-chan *player.Player
	leaving  <-chan *player.Player
	state    *GameState
}

func (gk *Gatekeeper) KeepTheGate() {
	for {
		select {
		case player := <-gk.entering:
			gk.logIn(player)
		case player := <-gk.leaving:
			gk.logOut(player)
		}
	}
}

func (gk *Gatekeeper) broadcast(msg string) {
	for _, p := range gk.state.Players {
		p.SendMsg(msg)
	}
}

func (gk *Gatekeeper) logIn(player *player.Player) {
  name := player.GetName()

  log.Printf("User logged in: %s", name)

	gk.state.Players[name] = player

  player.Character.SetSpawn(gk.state.DefaultRoom)
	player.Spawn()

	player.Look()

  go gk.broadcast(fmt.Sprintf("%s has logged in!", name))
}

func (gk *Gatekeeper) logOut(player *player.Player) {
  name := player.GetName()
	player.LeaveRoom("")

	log.Printf("User logged out: %s", name)

	delete(gk.state.Players, name)
	close(player.Channel)

	go gk.broadcast(fmt.Sprintf("%s has logged out!", name))
}
