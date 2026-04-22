package teapane

type PaneContainer struct{
	Panes []Pane
	Style ContainerStyle
}

func NewContainer(style ContainerStyle, panes ...Pane) PaneContainer{
	return PaneContainer{
		Style: style,
		Panes: panes,
	}
}
