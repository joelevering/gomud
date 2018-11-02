package character

import (
  "math"
  "sync"

  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/skills"
  "github.com/joelevering/gomud/statfx"
  "github.com/joelevering/gomud/stats"
)

type Character struct {
  Name       string            `json:"name"`
  Class      interfaces.ClassI
  Level      int               `json:"level"`
  ExpGiven   int               `json:"exp_given"`
  Exp        int
  NextLvlExp int
  MaxDet     int               `json:"max_health"`
  Det        int               `json:"health"`
  MaxStm     int               `json:"max_stamina"`
  Stm        int               `json:"stamina"`
  MaxFoc     int               `json:"max_focus"`
  Foc        int               `json:"focus"`
  Str        int               `json:"strength"`
  Flo        int               `json:"flow"`
  Ing        int               `json:"ingenuity"`
  Kno        int               `json:"knowledge"`
  Sag        int               `json:"sagacity"`

  Room       interfaces.RoomI
  Spawn      interfaces.RoomI
  InCombat   bool
  CmbSkill   *skills.Skill
  CmbSkillMu sync.Mutex
  Fx         map[statfx.StatusEffect]*statfx.SEInst
  Dots       map[statfx.DotType]*statfx.DotInst
}

func NewCharacter() *Character {
  return &Character{
    Class:      classes.Conscript,
    Level:      1,
    NextLvlExp: 10,
    MaxDet:     200,
    Det:        200,
    MaxStm:     100,
    Stm:        100,
    MaxFoc:     100,
    Foc:        100,
    Str:        10,
    Flo:        10,
    Ing:        10,
    Kno:        10,
    Sag:        10,
    Fx:         make(map[statfx.StatusEffect]*statfx.SEInst),
    Dots:       make(map[statfx.DotType]*statfx.DotInst),
  }
}

func (ch *Character) GetClassName() string {
  return ch.Class.GetName()
}

func (ch *Character) GetName() string {
  return ch.Name
}

func (ch *Character) SetName(name string) {
  ch.Name = name
}

func (ch *Character) GetLevel() int {
  return ch.Level
}

func (ch *Character) GetExp() int {
  return ch.Exp
}

func (ch *Character) GetExpGiven() int {
  return ch.ExpGiven
}

func (ch *Character) GetNextLvlExp() int {
  return ch.NextLvlExp
}

func (ch *Character) GetMaxDet() int {
  return ch.MaxDet
}

func (ch *Character) SetMaxDet(maxDet int) {
  ch.MaxDet = maxDet
}

func (ch *Character) GetDet() int {
  return ch.Det
}

func (ch *Character) SetDet(det int) {
  if det > ch.MaxDet {
    ch.Det = ch.MaxDet
  } else if det < 0 {
    ch.Det = 0
  } else {
    ch.Det = det
  }
}

func (ch *Character) GetMaxStm() int {
  return ch.MaxStm
}

func (ch *Character) SetMaxStm(maxStm int) {
  ch.MaxStm = maxStm
}

func (ch *Character) GetStm() int {
  return ch.Stm
}

func (ch *Character) SetStm(stm int) {
  ch.Stm = stm
}

func (ch *Character) GetMaxFoc() int {
  return ch.MaxFoc
}

func (ch *Character) SetMaxFoc(maxFoc int) {
  ch.MaxFoc = maxFoc
}

func (ch *Character) GetFoc() int {
  return ch.Foc
}

func (ch *Character) SetFoc(foc int) {
  ch.Foc = foc
}

func (ch *Character) GetStr() int {
  return ch.Str
}

func (ch *Character) SetStr(str int) {
  ch.Str = str
}

func (ch *Character) GetFlo() int {
  return ch.Flo
}

func (ch *Character) SetFlo(flo int) {
  ch.Flo = flo
}

func (ch *Character) GetIng() int {
  return ch.Ing
}

func (ch *Character) SetIng(ing int) {
  ch.Ing = ing
}

func (ch *Character) GetKno() int {
  return ch.Kno
}

func (ch *Character) SetKno(kno int) {
  ch.Kno = kno
}

func (ch *Character) GetSag() int {
  return ch.Sag
}

func (ch *Character) SetSag(sag int) {
  ch.Sag = sag
}

func (ch *Character) GetAtk() int {
  atk := 0
  atkStats := ch.Class.GetAtkStats()

  for _, stat := range atkStats {
    switch stat {
    case stats.Str:
      atk += ch.Str
    case stats.Flo:
      atk += ch.Flo
    case stats.Ing:
      atk += ch.Ing
    case stats.Kno:
      atk += ch.Kno
    case stats.Sag:
      atk += ch.Sag
    }
  }

  return atk
}

func (ch *Character) GetDef() int {
  def := 0
  defStats := ch.Class.GetDefStats()

  for _, stat := range defStats {
    switch stat {
    case stats.Str:
      def += ch.Str
    case stats.Flo:
      def += ch.Flo
    case stats.Ing:
      def += ch.Ing
    case stats.Kno:
      def += ch.Kno
    case stats.Sag:
      def += ch.Sag
    }
  }

  return def
}

func (ch *Character) GetSkills() []*skills.Skill {
  return ch.Class.SkillsForLvl(ch.Level)
}

func (ch *Character) SetCmbSkill(sk *skills.Skill) {
  ch.CmbSkillMu.Lock()
  ch.CmbSkill = sk
  ch.CmbSkillMu.Unlock()
}

func (ch *Character) GetSpawn() interfaces.RoomI {
  return ch.Spawn
}

func (ch *Character) SetSpawn(spawn interfaces.RoomI) {
  ch.Spawn = spawn
}

func (ch *Character) HealPct(pct float64) {
  amt := float64(ch.GetMaxDet()) * pct
  ch.Heal(int(amt))
}

func (ch *Character) Heal(amt int) {
  newDet := ch.GetDet() + amt
  if newDet > ch.GetMaxDet() {
    ch.FullHeal()
  } else {
    ch.Det = newDet
  }
}

func (ch *Character) FullHeal() {
  ch.Det = ch.MaxDet
}

func (ch *Character) RefocusPct(pct float64) {
  amt := float64(ch.GetMaxFoc()) * pct
  ch.Refocus(int(amt))
}

func (ch *Character) Refocus(amt int) {
  newFoc := ch.GetFoc() + amt
  if newFoc > ch.GetMaxFoc() {
    ch.SetFoc(ch.GetMaxFoc())
  } else {
    ch.SetFoc(newFoc)
  }
}

func (ch *Character) RecuperatePct(pct float64) {
  amt := float64(ch.GetMaxStm()) * pct
  ch.Recuperate(int(amt))
}

func (ch *Character) Recuperate(amt int) {
  newStm := ch.GetStm() + amt
  if newStm > ch.GetMaxStm() {
    ch.SetStm(ch.GetMaxStm())
  } else {
    ch.SetStm(newStm)
  }
}

func (ch *Character) IsDefeated() bool {
  if ch.GetDet() <= 0 {
    return true
  }

  return false
}

func (ch *Character) GainExp(exp int) (leveledUp bool) {
  ch.Exp += exp

  if ch.Exp >= ch.NextLvlExp {
    ch.LevelUp()
    return true
  }

  return false
}

func (ch *Character) LevelUp() {
  // Increase stats based on Class
  statGrowth := ch.Class.GetStatGrowth()
  ch.SetMaxDet(ch.MaxDet + statGrowth.Det)
  ch.SetStr(ch.Str + statGrowth.Str)
  ch.SetFlo(ch.Flo + statGrowth.Flo)
  ch.SetIng(ch.Ing + statGrowth.Ing)
  ch.SetKno(ch.Kno + statGrowth.Kno)
  ch.SetSag(ch.Sag + statGrowth.Sag)

  // Level up and carryover EXP
  ch.Level += 1
  ch.Exp = ch.Exp - ch.NextLvlExp

  // Set new EXP to level
  newNextLvlExp := float64(ch.NextLvlExp) * 1.25
  ch.NextLvlExp = int(math.Round(newNextLvlExp))

  ch.FullHeal()
}

func (ch *Character) ExpToLvl() int {
  return ch.NextLvlExp - ch.Exp
}

func (ch *Character) TickFx() {
  for _, fx := range ch.Fx {
    if fx.Duration == 0 {
      delete(ch.Fx, fx.Effect)
    } else {
      fx.Duration -= 1
    }
  }

  for _, dot := range ch.Dots {
    if dot.Duration == 0 {
      delete(ch.Dots, dot.Type)
    } else {
      dot.Duration -= 1
    }
  }
}

// private

func (ch *Character) getAndClearCmbSkill() *skills.Skill {
  ch.CmbSkillMu.Lock()
  defer ch.CmbSkillMu.Unlock()
  sk := ch.CmbSkill
  ch.CmbSkill = nil

  return sk
}

func (ch *Character) payForSkill(sk skills.Skill) bool {
  if sk.CostType == stats.Stm {
    cost := sk.CostAmt
    if ch.isConserving() {
      cost = int(float64(cost) * 0.5)
    }

    newStm := ch.GetStm() - cost
    if newStm < 0 {
      return false
    }

    ch.SetStm(newStm)
    return true
  }

  if sk.CostType == stats.Foc {
    newFoc := ch.GetFoc() - sk.CostAmt
    if newFoc < 0 {
      return false
    }

    ch.SetFoc(newFoc)
    return true
  }

  return false
}

func (ch *Character) addFx(i statfx.SEInst) {
  existing := ch.Fx[i.Effect]
  if existing != nil && existing.Duration > i.Duration {
    return
  }

  ch.Fx[i.Effect] = &i
}

func (ch *Character) addDot(i statfx.DotInst) {
  ch.Dots[i.Type] = &i
}

func (ch *Character) hasEffect(e statfx.StatusEffect) bool {
  return ch.Fx[e] != nil
}

func (ch *Character) isStunned() bool {
  return ch.Fx[statfx.Stun] != nil
}

func (ch *Character) isWeak() bool {
  return ch.Fx[statfx.Weak] != nil
}

func (ch *Character) isVulnerable() bool {
  return ch.Fx[statfx.Vulnerable] != nil
}

func (ch *Character) isConcentrating() bool {
  return ch.Fx[statfx.Concentration] != nil
}

func (ch *Character) isConserving() bool {
  return ch.Fx[statfx.Conserve] != nil
}

func (ch *Character) isDodging() bool {
  return ch.Fx[statfx.Dodging] != nil
}

func (ch *Character) isBleeding() bool {
  return ch.Dots[statfx.Bleed] != nil
}
