package classes

var ConscriptGrowth = StatGrowth{
  Health: 25,
  Str: 2,
  End: 3,
}

type Conscript struct {}

func (c Conscript) GetName() string {
  return "Conscript"
}

func (c Conscript) GetStatGrowth() StatGrowth {
  return ConscriptGrowth
}
