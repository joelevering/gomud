# gomud

A golang-based MUD client.
* `go build`
* `./gomud`
* `telnet <ip> <port>`

Available commands:
'move <exit key>' to move to a new room
'/look' or 'look' to see where you are
'/look <npc name>' or 'look <npc name>' to see more details about an NPC
'/list' or '/ls' to see who is currently in your room
'/help' or 'help' to repeat this message

Anything else will be broadcast as a message to the people in your room

# BUGS

# TODO
* Write a boat load of tests
  * All commands
* Add NPCs/bots
  * ~~NPCs have stats and live in rooms~~
  * ~~Load NPCs~~
  * ~~Show NPCs in the room~~
  * ~~Make the /look command give details on NPCs~~
  * Let NPCs follow scripts (so they can talk and move)
* Add combat
* Chat Improvements
  * Don't send blank messages
  * Require 'say'
  * Add 'yell'
* Admin login and features (kick people)
* Add stats
* Add classes
* Add skills
* Add progression
* Add items
* Add UI
  * TERMUI Appears unable to handle custom Writers (e.g. net.Conn)
