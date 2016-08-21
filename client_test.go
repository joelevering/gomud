package main

//
// import (
// 	"strings"
// 	"testing"
// )
//
// func Test_SendMsg(t *testing.T) {
// 	ch := make(chan string)
// 	cli := NewClient(ch)
//
// 	go cli.SendMsg("testing SendMsg")
//
// 	res := <-ch
//
// 	if !strings.Contains(res, "testing SendMsg") {
// 		t.Error("Expected SendMsg to send 'testing SendMsg' to channel, but it didn't")
// 	}
// }
//
// // This test is broken because the Client does not have a Room
// // I want to mock Room but this would require that Client take
// // An interface instead of a room, which would in turn necessitate
// // Getters for all exposed Room attributes and (if making a separate
// // mocks package) somehow getting the Mocks package to recognize
// // structs defined in main (e.g. Exits and other return values required
// // for the mock Room to adhere to the new interface)
// func TestSay(t *testing.T) {
// 	ch := make(chan string)
// 	cli := NewClient(ch)
//
// 	go cli.Say("testing Say")
//
// 	res := <-ch
//
// 	if !strings.Contains(res, "testing Say") {
// 		t.Error("Expected Say to send 'testing Say' to the room, but it didn't")
// 	}
// }
