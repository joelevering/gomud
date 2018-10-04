# gomud

A golang-based MUD client.
* `go build`
* `./gomud`
* In another terminal: `telnet <ip> <port>`
* 'help' to show available commands

Notes:
* Update channels to be unique to the NPCs room
* Let NPCs move between rooms (will require them to unsub/re-sub behavior)
* Add a "move" action
* Add some more triggers
* Refactor action pivot?
  * Make unique types of NPCs (possible starting point: slimes vs dogs vs crabs)
* Test "chance" system
* Test triggered behavior in general
  * Add mock queue
