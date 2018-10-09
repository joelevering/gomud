package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/joelevering/gomud/client"
)

type ConnHandler struct {
	entering chan *client.Client
	leaving  chan *client.Client
  state    *GameState
}

func (handler *ConnHandler) Handle(conn net.Conn) {
	defer conn.Close()

	ch := make(chan string)
	cli := client.NewClient(ch, handler.state.Queue)
	go cli.StartWriter(conn)

	cli.SetName(confirmName(cli, conn))

	handler.entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
    cli.Cmd(input.Text())
	}

	handler.leaving <- cli
}

func confirmName(cli *client.Client, conn net.Conn) string {
	var confirmed, who string

	for strings.ToUpper(confirmed) != "Y" {
		cli.SendMsg("Who are you?")
		input := bufio.NewScanner(conn)
		input.Scan()
		who = input.Text()

		cli.SendMsg(fmt.Sprintf("Are you sure you want to be called \"%s\"? (Y to confirm)", who))
		input.Scan()
		confirmed = input.Text()
	}

	return who
}
