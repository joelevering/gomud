package character

import (
  "math"

  "github.com/joelevering/gomud/classes"
  "github.com/joelevering/gomud/interfaces"
  "github.com/joelevering/gomud/stats"
)

type Character struct {
  Name       string            `json:"name"`
  ClassName  string            `json:"class"`
  Class      interfaces.ClassI
  Level      int               `json:"level"`
  Exp        int               `json:"experience"`
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
  InCombat   bool
  Spawn      interfaces.RoomI
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
  }
}

func (ch *Character) ResetStats() {
  lvl := ch.Level
  exp := ch.Exp

  ch.Level = 0
  for i := 0; i < lvl; i++ {
    ch.levelUp()
  }

  ch.Exp = exp
}

func (ch *Character) SetClass() {
  ch.Class = classes.ByName[ch.ClassName]
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
  ch.Det = det
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

func (ch *Character) GetSpawn() interfaces.RoomI {
  return ch.Spawn
}

func (ch *Character) SetSpawn(spawn interfaces.RoomI) {
  ch.Spawn = spawn
}

func (ch *Character) Heal() {
  ch.Det = ch.MaxDet
}

func (ch *Character) EnterCombat() {
  ch.InCombat = true
}

func (ch *Character) LeaveCombat() {
  ch.InCombat = false
}

func (ch *Character) IsInCombat() bool {
  return ch.InCombat
}

func (ch *Character) GainExp(exp int) (leveledUp bool) {
  ch.Exp += exp

  if ch.Exp >= ch.NextLvlExp {
    ch.levelUp()
    return true
  }

  return false
}

func (ch *Character) levelUp() {
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

  ch.Heal()
}

func (ch *Character) ExpToLvl() int {
  return ch.NextLvlExp - ch.Exp
}
