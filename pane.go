package teapane

type Pane struct{
	height        int
	width         int
	displayHeight int
	displayWidth  int
	displayString string
	style         PaneStyle
}


func (p Pane) View() string{
	return ""
}
