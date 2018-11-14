package player

import (
  "fmt"
  "strings"

  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/skills"
)

func Help(words []string) string {
  if len(words) == 1 {
    return helpMsg
  }

  switch strings.ToLower(words[1]) {
  case "say":
    return sayMsg
  case "yell":
    return yellMsg
  case "look":
    return lookMsg
  case "list":
    return listMsg
  case "move":
    return moveMsg
  case "status":
    return statusMsg
  case "classes":
    return classesMsg
  case "change subclass":
    return changeSubclassMsg
  case "tier", "tiers":
    return tierMsg
  case "change":
    return changeMsg
  case "attack":
    return attackMsg
  case "combat":
    return combatMsg
  case "skill", "skills":
    if len(words) == 2 {
      return skillMsg
    }

    return HelpSkill(strings.Join(words[2:], " "))
  case "class":
    if len(words) == 2 {
      return classMsg
    }

    return HelpClass(strings.Join(words[2:], " "))
  case "exit", "quit":
    return exitMsg
  }

  return "Sorry, I'm not sure what you need help with."
}

func HelpSkill(skName string) string {
  sk := skills.GetSkill(skName)
  if sk == nil {
    return "Sorry, I can't find that skill."
  }

  msg := `***************%s***************
%s

Cost: %d %s`

  return fmt.Sprintf(msg, sk.Name, sk.Desc, sk.CostAmt, sk.CostType)
}

func HelpClass(clName string) string {
  cl := classes.Find(clName)
  if cl == nil {
    return "Sorry, I can't find that class."
  }

  msg := `***************%s***************
%s

Attack stats: %v
Defense stats: %v

Skills:
%s`

  skillLines := []string{}
  for _, sk := range cl.GetSkills() {
    skillLines = append(skillLines, fmt.Sprintf("  * %s (%d)", sk.Skill.Name, sk.Level))
  }
  skillsMsg := strings.Join(skillLines, "\n")

  return fmt.Sprintf(msg, cl.GetName(), cl.GetDesc(), cl.GetAtkStats(), cl.GetDefStats(), skillsMsg)
}

const helpMsg = `***************Help***************

Available commands:
'status' to see details about your character including available skills
'say <message>' to communicate with people in your room
'move <exit key>' to move to a new room
'list' to see who is currently in your room
'look' to see where you are
'look <npc name>' to see more details about an NPC
'attack <npc name>' to start combat
'attack <npc name> <skill name>' to start combat by using a skill
'classes' for information on your character's available classes
'change <class name>' to change your class
'exit' or 'quit' to log out
'help' to repeat this message

Most commands have their first letter as a shortcut (e.g. 'l' for look, 'a' for attack)

When you are in combat, prepare a skill for your next action by inputting its name

Shortcut: 'h'`

const sayMsg = `***************Say***************

Type 'say' followed by a message to broadcast that message to the other players in your current location.

Because the message is localized, if you leave your location or if another player leaves your location,
you won't be able to use 'say' to communicate unless you find each other again.

Shortcut: 's'`

const yellMsg = `***************Yell***************

Type 'yell' followed by a message to broadcast that message to the other players in your current location AND
any adjacent locations.

Yell can be helpful if you become separated from someone and are trying to find them.
It can also help guide a nearby player to your location.

Shortcut: 'y'`

const listMsg = `***************List***************

'list' will show you the names of the players and NPCs who are in your current location.

Shortcut: 'ls'`

const lookMsg = `***************Look***************

'look' will describe your current location, show you the available exits, and the names of the other players
and NPCs who are in the same area.

You can use the 'key' for an exit in conjunction with the 'move' command in order to change your location.
The 'key' for an exit is the letter or letters wrapped in parentheses in the exit description.
For instance, if the exit description says '(O)ut the door', the exit key is 'o'.
Use 'help move' for more information on moving with exit keys.

Shortcut: 'l'`

const moveMsg = `***************Move***************

Type 'move' followed by an exit key to change your current location.

Exits are displayed when you enter a new location and when you use the 'look' command.
If you are in a room with an exit that has a key of 'o', you can move to a new location
by typing 'move o'.

Use 'help look' for more information on identifying exit keys

Shortcut: 'm'`

const statusMsg = `**************Status**************

'status' will display information about your character.

The first section includes information on your character's class and progress in that class.

The second section shows how close you are to falling in combat and tracks the resources used to perform skills.

The third section shows your ability in statistics that impact attack, defense, and skill effectiveness.

The fourth section lists the skills currently available to you. Skills are unique to each class and are gained
by increasing your level in a class. If you have none, you should try leveling up!

Type 'help <statistic>' for more information on how each individual metric is used.
You can also use 'help skill <skill name>' for information on an individual skill.

Shortcut: 'st'`

const classesMsg = `**************Classes**************

The 'classes' command is used to list the classes currently available to you.
Each class is accompanied by a brief description of the class.
Use 'help class' followed by the class name for more information on a specific class.

Use 'change' to change your current class to one of those listed by 'classes'. See 'help change' for more infomation.

More classes can be unlocked by playing the game. Don't forget to use this command when you reach a milestone.

Shortcut: 'cl'`

const changeSubclassMsg = `**************Change Subclass**************

Use 'change subclass' followed by the name of a class to update your subclasses.

You can only use the name of a class in the tier below your main class tier.
For instance if your current class is Tier 2, you can only use a Tier 1 class name with this command.

Your subclass can be changed at anytime outside of combat.
Changing your subclass will change your skills.
Use 'classes' to see the classes available to you.

See 'help change' for information on changing your main class.
See 'help class tiers' for information on the class system itself.

Shortcut: 'c sc'`

const tierMsg = `**************Tiers**************

All classes have a Tier based on their power. Characters start with access to Tier 1 classes.

Higher-Tier classes can be unlocked by reaching maximum level with a specific class or classes from lower tiers.
For instance, the Tier 2 class Minder can be unlocked by reaching maximum level in
the Tier 1 classes Conscript and Augur.

After gaining access to Tier 2 classes, a character has the ability to equip multiple classes simultaneously.
Each class that is equipped gives access to that classes skills. The main class (highest Tier class) also controls
what stats are used for attack and defense and the maximum determination of a character.

A character can equip a class from each Tier they've unlocked a class for. For instance, if a character has access
to Tier 3 classes, they can equip a Tier 3, a Tier 2, and a Tier 1 class at the same time.

However, classes can only be equipped "on top" of a class of maximum level from the Tier below. For example, if
you unlocked the Tier 2 class Minder, you could equip it on top of your maxed out Conscript or Augur classes, but
if you wanted to use another Tier 1 class like Athlete, that class would need to be at maximum level to also equip
Minder.

Note: You can always equip classes in a lower Tier than your most powerful class by using the 'change' command.
If you unlocked Minder but want to level up Athlete, use 'change athlete' to set Athlete as your main class.

See 'help change' for information on changing your main class.
See 'help change subclass' for information on changing your subclasses.
`

const changeMsg = `**************Change**************

Use 'change' followed by the name of a class to change your current class to the named one.

Your class can be changed at anytime outside of combat.
Changing your class will change your maximum determination, your level, and your skills.
Use 'classes' to see the classes available to you.

See 'help change subclass' for information on changing subclasses.
See 'help class tiers' for information on the class system itself.

Shortcut: 'c'`

const attackMsg = `**************Attack**************

Use 'attack' followed by the name of an NPC to begin combat with that NPC.
For instance: 'attack slime'

The name of the NPC does not need to be exact. Including just one word or some letters of a word should work.
For instance an NPC named 'Forest Slime Runt' could be attacked by using 'attack runt', 'attack fo', etc.

After the name of an NPC, you can also include the name of a skill to begin combat using that skill.
For instance: 'attack slime charge'
This would begin combat with an NPC named 'slime' using the 'charge' skill.

In combat, you and the NPC will take turns attacking each other or using a skill.
For more information on how combat works, use 'help combat'.

Shortcut: 'a'`

const combatMsg = `**************Combat**************

Combat happens in real time, with turns alternating back and forth between combatants.

During combat you will default to doing a standard attack on your enemy.

Alternatively, you can use a skill by typing the name of the skill. This will prepare the skill to be used
in lieu of an attack on your next turn. You can replace your prepared skill with a new one by typing the
name of a different skill before your turn.

Using skills usually consumes stamina, focus, or both. If you do not have enough of a particular resource
the chosen skill will not be executed.

Combat ends automatically when one of the combatants runs out of determination.
If you are victorious, you'll gain experience which moves you closer to leveling up.
If you are defeated, you'll be respawned momentarily in your last established spawn location.

Defeated enemies respawn after some time has passed.

See 'help attack' for more information on initiating combat.`

const skillMsg = `**************Skills**************

Skills are used in combat for a variety of effects.

Use 'help combat' for more information on using skills in combat.
Use 'help skill' followed by the name of a skill for more information on a specific skill.`

const classMsg = `**************Class**************

Your class defines your maximum determination, level, and skills.
You can change your class with the 'change' command. Use 'help change' for more information.
You can see your currently-available classes by using the 'classes' command. Use 'help classes' for more information.

For more information on a specific class, use 'class' followed by the name of the class.`

const exitMsg = `***************Exit***************

Use 'exit' (or 'quit') to safely disconnect from the game.
Using this command will ensure your progress is saved.`
