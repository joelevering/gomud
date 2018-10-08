package client

import (
	"fmt"
	"time"

	"github.com/joelevering/gomud/interfaces"
)

const tick = 1500 * time.Millisecond

type CombatInstance struct {
	cli  interfaces.CliI
	npc  interfaces.NPCI
}

type CombatResults struct {
  pcDmg     int
  npcDmg    int
  pcHeal    int
  npcHeal   int
  pcEffect  string
  npcEffect string
}

func (ci *CombatInstance) Start() {
	for true {
    combatOver := ci.Loop(true)
    if combatOver { break }

		time.Sleep(tick)
	}
}

func (ci *CombatInstance) Loop(report bool) (combatOver bool) {
  pcResults := ci.getPCResults()
  npcResults := ci.getNPCResults()

  npcDmg := pcResults.npcDmg + npcResults.npcDmg
  pcDmg := pcResults.pcDmg + npcResults.pcDmg

  ci.applyDamage(npcDmg, pcDmg)
  if report {
    ci.report(npcDmg, pcDmg)
  }

  if ci.cliIsDead() {
    ci.cli.Die(ci.npc)
    return true
  }

  if ci.npcIsDead() {
    ci.npc.Die(ci.cli)
    ci.cli.Defeat(ci.npc)
    return true
  }

  return false
}

func (ci *CombatInstance) getPCResults() *CombatResults {
  res := &CombatResults{}
  combatCmd := ci.cli.GetCombatCmd()

  if len(combatCmd) == 0 {
    res.npcDmg = CalcAtkDmg(ci.cli.GetStr(), ci.npc.GetEnd())
    return res
  }

  // TODO turn this into a map loaded once on app load and accessible by any CI
	switch combatCmd[0] {
	case "bash":
    smiteStr := int(float64(ci.cli.GetStr()) * 1.1)
    res.npcDmg = CalcAtkDmg(smiteStr, ci.npc.GetEnd())
  default: // attack
    res.npcDmg = CalcAtkDmg(ci.cli.GetStr(), ci.npc.GetEnd())
  }

  return res
}

func (ci *CombatInstance) getNPCResults() *CombatResults {
  return &CombatResults{
    pcDmg: CalcAtkDmg(ci.npc.GetStr(), ci.cli.GetEnd()),
  }
}

// Block 1/1000th of the damage per point of Endurance
func CalcAtkDmg(atkStr int, defEnd int) int {
	endPercentDamageBlocked := float64(defEnd) * 0.001
	dmg := (1.0 - endPercentDamageBlocked) * float64(atkStr)

	return int(dmg)
}

func (ci *CombatInstance) applyDamage(npcDmg, pcDmg int) {
  npcHealth := ci.npc.GetHealth()
  pcHealth := ci.cli.GetHealth()

  if pcHealth-pcDmg < 0 {
    ci.cli.SetHealth(0)
  } else {
    ci.cli.SetHealth(pcHealth - pcDmg)
  }

  if npcHealth-npcDmg < 0 {
    ci.npc.SetHealth(0)
  } else {
    ci.npc.SetHealth(npcHealth - npcDmg)
  }
}

func (ci *CombatInstance) cliIsDead() bool {
  return ci.cli.GetHealth() <= 0
}

func (ci *CombatInstance) npcIsDead() bool {
  return ci.npc.GetHealth() <= 0
}

func (ci *CombatInstance) report(npcDmg, pcDmg int) {
  ci.cli.SendMsg(fmt.Sprintf("%s took %d damage!", ci.npc.GetName(), npcDmg))
  ci.cli.SendMsg(fmt.Sprintf("You took %d damage!", pcDmg))
  ci.cli.SendMsg("")

  if ci.cliIsDead() {
    return // handled in Start()
  }

  var npcMsg string

  if ci.npcIsDead() {
    npcMsg = fmt.Sprintf("%s is defeated!", ci.npc.GetName())
  } else {
    npcMsg = fmt.Sprintf("%s has %d/%d", ci.npc.GetName(), ci.npc.GetHealth(), ci.npc.GetMaxHealth())
  }

  ci.cli.SendMsg(fmt.Sprintf("You have %d/%d health left. %s", ci.cli.GetHealth(), ci.cli.GetMaxHealth(), npcMsg))
}
