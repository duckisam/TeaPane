package teapane

type Pane struct{
	Height        int
	Width         int
	DisplayHeight int
	DisplayWidth  int
	DisplayString string
	Style         PaneStyle
}


func NewPane(width, height int, haveBorder bool) Pane{
	dh := height
	dw := width

	if haveBorder{
		dh -= 2
		dw -= 2
	}

	return Pane{
		DisplayHeight: dh,
		DisplayWidth: dw,
		Width: width,
		Height: height,
		Style: PaneStyle{Border: DefaultBorder},
	}

}

func (p Pane) View() string{
	return RenderPane(p, p.Width, p.Height)
}
