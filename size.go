package teapane

type SizeUnit int

const (
	UnitAuto SizeUnit = iota
	UnitFixed
	UnitPercent
	
)

type Size struct{
	value int
	unit SizeUnit
}

func Auto()         Size {return Size{unit: UnitPercent}}
func Fixed(n int)   Size {return Size{value: n, unit: UnitFixed}}
func Percent(n int) Size {return Size{value: n, unit: UnitPercent}}
