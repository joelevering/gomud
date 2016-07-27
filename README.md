# gomud

A golang-based MUD client.
* `go build`
* `./gomud`
* `telnet <ip> <port>`

Available commands:
'say <message>' to communicate with people in your room
'move <exit key>' to move to a new room
'look' to see where you are
'look <npc name>' to see more details about an NPC
'list' to see who is currently in your room
'help' to repeat this message

Anything else will be broadcast as a message to the people in your room

# BUGS
* Help method says to do /help but that doesn't work anymore

# TODO
* Confirm name on entry during sign in (some people don't realize they're entering a name)
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
  * ~~Don't send blank messages~~
  * ~~Require 'say'~~
  * ~~Add 'yell'~~
* Admin login and features (kick people)
* Add stats
* Add classes
* Add skills
* Add progression
* Add items
* Add UI
  * TERMUI Appears unable to handle custom Writers (e.g. net.Conn)
