package teapane

type SizeUnit int

const (
	UnitAuto SizeUnit = iota
	UnitFixed
	UnitPercent
	
)

type Size struct{
	Value int
	Unit SizeUnit
}

func Auto()         Size {return Size{Unit: UnitPercent}}
func Fixed(n int)   Size {return Size{Value: n, Unit: UnitFixed}}
func Percent(n int) Size {return Size{Value: n, Unit: UnitPercent}}

func ResolveSize(s Size, total int) int{
	switch s.Unit{
	case UnitFixed:
		return s.Value
	case UnitPercent:
		return (s.Value * total) / 100
	default:
		return 0
	}
}
