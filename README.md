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

* Ask for more information when receiving single arg for multi arg commands
  * 's' as a command returns nothing

# TODO

* Combat
  * ~~Allow players to initiate combat~~
  * ~~Basic combat algorithm~~
  * ~~Test existing combat~~
    * ~~Move away from inline code in the combat loop, make it a testable method~~
    * Could potentially test `report` method by having it return strings and having `Start` send them to the client
  * Switch to smaller stat values + floats and rounding
  * ~~Implement NPC death~~
  * ~~Implement PC death~~
  * ~~Implement NPC respawn~~
    * ~~NPCs spawn after rooms are generated~~
    * ~~NPCs can leave/join rooms~~
  * Implement Skills/Abilities
    * Move CombatInstance to live on Client (it needs to be able to take commands while active)
  * Implement scripted NPC combat behavior
    * How can NPC behavior (in combat and out of combat) be stored? A list of methods to call in combat?
  * Handle leaving the room during combat (e.g. fleeing)
* ~~Add progression~~
  * ~~PC has an exp stat~~
  * ~~PC can gain experience~~
  * ~~PC can level up~~
  * ~~PC heals on level up~~
  * ~~PC gains stats on level up~~
  * ~~Level up exp required scales~~
* ~~Add a command to see your stats~~
* Write tests
  * All commands
* Refactor Cli/NPC code to share an interface
  * Entering/Leaving rooms
  * Saying/Yelling
  * Attacking
* Implement scripted NPC behavior (out of combat, e.g. changing rooms, talking etc.)
* Chat Improvements
  * Add ability for people to toggle on/off defaulting to 'say'-ing unrecognized commands
* Admin login and features (kick people)
* Add classes
* Add items
* Add UI
  * TERMUI Appears unable to handle custom Writers (e.g. net.Conn)
