# gomud

A golang-based MUD client.
* `go build`
* `./gomud`
* In another terminal: `telnet <ip> <port>`
* 'help' to show available commands

Ideas for actually building out the game after the skeleton is done:

* Easy/traditional RPG with RNG elements (a la Dragon Quest)
  * Monsters use % chance to choose their action
  * Interesting/unique settings
  * Interesting/unique classes and abilities
  * Questions:
    * Static ability gain vs. ability points/skill tree
    * Is there anyway to incorporate party-style RPG combat into a MUD? E.g. each player controls a 'party'
* Souls-like combat
  * Use a mostly-static list of abiltiies (e.g. dodge, parry)
  * Require players to learn how to respond to each enemy/room
  * Combat command input timing would be pretty strict which could isolate potential playerbase
* Secrets-based gameplay
  * Create a lot of hidden rooms (secret exits e.g. Tornado)
  * No traditional quest system, but objectives can be pieced together from NPC dialogue and items
    * E.g. give only certain players certain parts of the puzzle
  * Low-chance/limited-time enemy drops
  * Higher upkeep, especially if playerbase grows (would require adding way more items/rooms/etc.)
