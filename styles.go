package teapane

type PaneStyle struct{
	grow      int
	shrink    int
	basis     Size
	minWidth  Size
	maxWidth  Size
	minHeight Size 
	maxHeight Size 
	order     int
	wrapText  bool
	alignSelf AlignSelf
}

type ContainerStyle struct{
	flexDirection FlexDirection 
	flexWrap      FlexWrap
	justify       Justify
	alignItems    AlignItems
	alignContent  AlignContent
	gapRow        int
	gapColumn     int
}

type FlexDirection string
type FlexWrap      string
type Justify       string
type AlignItems    string
type AlignContent  string
type AlignSelf    string

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
	topLeft     string
	topRIght    string
	bottomLeft  string
	bottomRight string
	horizontal  string
	vertical    string
}
