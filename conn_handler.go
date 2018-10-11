package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/joelevering/gomud/player"
)

type ConnHandler struct {
	entering chan *player.Player
	leaving  chan *player.Player
  state    *GameState
}

func (handler *ConnHandler) Handle(conn net.Conn) {
	defer conn.Close()

	ch := make(chan string)
	player := player.NewPlayer(ch, handler.state.Queue)
	go player.StartWriter(conn)

	player.SetName(confirmName(player, conn))

	handler.entering <- player

	input := bufio.NewScanner(conn)
	for input.Scan() {
    player.Cmd(input.Text())
	}

	handler.leaving <- player
}

func confirmName(player *player.Player, conn net.Conn) string {
	var confirmed, who string

	for strings.ToUpper(confirmed) != "Y" {
		player.SendMsg("Who are you?")
		input := bufio.NewScanner(conn)
		input.Scan()
		who = input.Text()

		player.SendMsg(fmt.Sprintf("Are you sure you want to be called \"%s\"? (Y to confirm)", who))
		input.Scan()
		confirmed = input.Text()
	}

	return who
}
