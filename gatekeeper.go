package main

import (
	"fmt"
	"log"

	"github.com/joelevering/gomud/player"
)

type nameClaimRequest struct {
  name  string
  reply chan bool
}

type Gatekeeper struct {
	entering     <-chan *player.Player
	leaving      <-chan *player.Player
	claimName    <-chan nameClaimRequest
	releaseName  <-chan string
	state        *GameState
	pendingNames map[string]bool
}

func (gk *Gatekeeper) KeepTheGate() {
	for {
		select {
		case player := <-gk.entering:
			gk.logIn(player)
		case player := <-gk.leaving:
			gk.logOut(player)
		case req := <-gk.claimName:
			available := gk.isNameAvailable(req.name)
			if available {
				gk.pendingNames[req.name] = true
			}
			req.reply <- available
		case name := <-gk.releaseName:
			delete(gk.pendingNames, name)
		}
	}
}

func (gk *Gatekeeper) isNameAvailable(name string) bool {
	if _, taken := gk.state.Players[name]; taken {
		return false
	}
	_, reserved := gk.pendingNames[name]
	return !reserved
}

func (gk *Gatekeeper) broadcast(msg string) {
	for _, p := range gk.state.Players {
		p.SendMsg(msg)
	}
}

func (gk *Gatekeeper) logIn(player *player.Player) {
  name := player.GetName()

  log.Printf("User logged in: %s", name)

	delete(gk.pendingNames, name)
	gk.state.Players[name] = player

	player.Look()

  go gk.broadcast(fmt.Sprintf("%s has logged in!", name))
}

func (gk *Gatekeeper) logOut(player *player.Player) {
  player.Save()
  player.LeaveRoom("")
  close(player.Logout)

  name := player.GetName()
  log.Printf("User logged out: %s", name)

  delete(gk.state.Players, name)
  close(player.Channel)

  go gk.broadcast(fmt.Sprintf("%s has logged out!", name))
}
