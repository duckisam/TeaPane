package teapane

import (
	"github.com/charmbracelet/lipgloss"
)

type PaneStyle struct{
	Grow          int
	Shrink        int
	Basis         Size
	CrossBasis    Size
	MinBasis      int
	MaxBasis      int
	MinCrossBasis int
	MaxCrossBasis int
	Order         int
	WrapText      bool
	AlignSelf     AlignSelf
	Border        PaneBorder
}

type ContainerStyle struct{
	FlexDirection FlexDirection
	FlexWrap      FlexWrap
	Justify       Justify
	AlignItems    AlignItems
	AlignContent  AlignContent
	GapRow        int
	GapColumn     int
}

func NewContainerStyle() ContainerStyle{
	return ContainerStyle{
		FlexDirection: DirectionRow,
		FlexWrap: WrapNone,
		Justify: JustifyStart,
		AlignItems: AlignItemStrech,
		AlignContent: AlignContentStrech,
	}

}

type FlexDirection string
type FlexWrap      string
type Justify       string
type AlignItems    string
type AlignContent  string
type AlignSelf     string

const (
	DirectionRow     FlexDirection = "row"
	DirectionRowR    FlexDirection = "row-reverse"
	DirectionColumn  FlexDirection = "column"
	DirectionColumnR FlexDirection = "column-reverse"

	WrapNone    FlexWrap = "nowrap"
	WrapWrap    FlexWrap = "wrap"
	WrapReverse FlexWrap = "wrap-reverse"

	JustifyStart   Justify = "flex-start"
	JustifyEnd     Justify = "flex-end"
	JustifyCenter  Justify = "center"
	JustifyBetween Justify = "space-between"
	JustifyAround  Justify = "space-around"
	JustifyEvenly  Justify = "space-evenly"

	AlignItemStrech   AlignItems = "stretch"
	AlignItemStart    AlignItems = "flex-start"
	AlignItemEnd      AlignItems = "flex-end"
	AlignItemCenter   AlignItems = "center"
	AlignItemBaseline AlignItems = "baseline"


	AlignContentStrech  AlignContent = "stretch"
	AlignContentStart   AlignContent = "flex-start"
	AlignContentEnd     AlignContent = "flex-end"
	AlignContentBetween AlignContent = "space-between"
	AlignContentAround  AlignContent = "space-around"
	AlignContentEvenly  AlignContent = "space-evenly"
	AlignContentCenter  AlignContent = "center"

	AlignSelfAuto     AlignSelf = "auto"
	AlignSelfCenter   AlignSelf = "center"
	AlignSlefStart    AlignSelf = "flex-start"
	AlignSelfEnd      AlignSelf = "flex-end"
	AlignSelfStretch  AlignSelf = "stretch"
    AlignSelfBaseline AlignSelf = "baseline"
)

type PaneBorder struct{
	Enabled     bool
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
	Color       lipgloss.Color
}

var (
	DefaultBorder = PaneBorder{Enabled: true, TopLeft: "┌", TopRight: "┐", BottomLeft: "└", BottomRight: "┘", Vertical: "│", Horizontal: "─", Color: lipgloss.Color("#FFFFFF")}
)
