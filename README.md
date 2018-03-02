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

Most commands have their first letter as a shortcut

# BUGS

# TODO

* Add combat
  * ~~Allow players to initiate combat~~
  * ~~Basic combat algorithm~~
  * Switch to smaller stat values + floats and rounding
  * Test combat
    * Move away from inline code in the combat loop, make it a testable method
    * Make CombatInstance take an interface instead of actual Client (move it out of client package?)
  * Implement NPC death
  * Implement Player death
  * Implement NPC spawning
  * Implement PC death
  * Implement Skills/Abilities
    * Move CombatInstance to live on Client (it needs to be able to take commands while active)
* Add progression
* Add a command to see your stats
* Write tests
  * All commands
* Let NPCs follow scripts (so they can talk and move)
* Chat Improvements
  * Add ability for people to toggle on/off defaulting to 'say'-ing unrecognized commands
* Admin login and features (kick people)
* Add stats
* Add classes
* Add items
* Add UI
  * TERMUI Appears unable to handle custom Writers (e.g. net.Conn)
