package client

import (
	"fmt"
	"time"

	"github.com/joelevering/gomud/interfaces"
)

const tick = 1500 * time.Millisecond

type CombatResults struct {
  pcDmg     int
  npcDmg    int
  pcHeal    int
  npcHeal   int
  pcEffect  string
  npcEffect string
}

type CombatInstance struct {
	cli  interfaces.CliI
  npc  interfaces.NPCI
  pc   interfaces.CharI
	npcC interfaces.CharI
}

func (ci *CombatInstance) Start() {
  ci.pc.EnterCombat()

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

  if ci.pcIsDead() {
    ci.pc.LeaveCombat()
    ci.cli.LoseCombat(ci.npcC)
    return true
  }

  if ci.npcIsDead() {
    ci.pc.LeaveCombat()
    ci.npc.LoseCombat(ci.pc)
    ci.cli.WinCombat(ci.npcC)
    return true
  }

  return false
}

func (ci *CombatInstance) getPCResults() *CombatResults {
  res := &CombatResults{}
  combatCmd := ci.cli.GetCombatCmd()

  if len(combatCmd) == 0 {
    res.npcDmg = CalcAtkDmg(ci.pc.GetAtk(), ci.npcC.GetDef())
    return res
  }

  // TODO turn this into a map loaded once on app load and accessible by any CI
	switch combatCmd[0] {
	case "bash":
    smiteStr := int(float64(ci.pc.GetAtk()) * 1.1)
    res.npcDmg = CalcAtkDmg(smiteStr, ci.npcC.GetDef())
  default: // attack
    res.npcDmg = CalcAtkDmg(ci.pc.GetAtk(), ci.npcC.GetDef())
  }

  return res
}

func (ci *CombatInstance) getNPCResults() *CombatResults {
  return &CombatResults{
    pcDmg: CalcAtkDmg(ci.npcC.GetAtk(), ci.pc.GetDef()),
  }
}

// Block 1/1000th of the damage per point of Endurance
func CalcAtkDmg(atkStr int, defEnd int) int {
	endPercentDamageBlocked := float64(defEnd) * 0.001
	dmg := (1.0 - endPercentDamageBlocked) * float64(atkStr)

	return int(dmg)
}

func (ci *CombatInstance) applyDamage(npcDmg, pcDmg int) {
  npcDet := ci.npcC.GetDet()
  pcDet := ci.pc.GetDet()

  if pcDet-pcDmg < 0 {
    ci.pc.SetDet(0)
  } else {
    ci.pc.SetDet(pcDet - pcDmg)
  }

  if npcDet-npcDmg < 0 {
    ci.npcC.SetDet(0)
  } else {
    ci.npcC.SetDet(npcDet - npcDmg)
  }
}

func (ci *CombatInstance) pcIsDead() bool {
  return ci.pc.GetDet() <= 0
}

func (ci *CombatInstance) npcIsDead() bool {
  return ci.npcC.GetDet() <= 0
}

func (ci *CombatInstance) report(npcDmg, pcDmg int) {
  ci.cli.SendMsg(fmt.Sprintf("%s took %d damage!", ci.npcC.GetName(), npcDmg))
  ci.cli.SendMsg(fmt.Sprintf("You took %d damage!", pcDmg))
  ci.cli.SendMsg("")

  if ci.pcIsDead() {
    return // handled in Start()
  }

  var npcMsg string

  if ci.npcIsDead() {
    npcMsg = fmt.Sprintf("%s is defeated!", ci.npcC.GetName())
  } else {
    npcMsg = fmt.Sprintf("%s has %d/%d", ci.npcC.GetName(), ci.npcC.GetDet(), ci.npcC.GetMaxDet())
  }

  ci.cli.SendMsg(fmt.Sprintf("You have %d/%d health left. %s", ci.pc.GetDet(), ci.pc.GetMaxDet(), npcMsg))
}
