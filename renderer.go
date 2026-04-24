package teapane

import (
	"cmp"
	"slices"
	"strings"
	"github.com/charmbracelet/lipgloss"
)

func RenderPane(p Pane, width, height int) string{ 
	if width <= 0 || height <= 0{
		return ""
	}

	var paneString strings.Builder 
	textWidth  := width
	textHeight := height

	borderStyle := lipgloss.NewStyle().Foreground(DefaultBorder.Color)
	
	if p.Style.Border.Enabled{
		textWidth  -= 2
		textHeight -= 2
		paneString.WriteString(borderStyle.Render(p.Style.Border.TopLeft))
		paneString.WriteString(borderStyle.Render(strings.Repeat(p.Style.Border.Horizontal, width - 2)))
		paneString.WriteString(borderStyle.Render(p.Style.Border.TopRight))
		paneString.WriteRune('\n')
	}

	var toWrite []string
	if p.Style.WrapText{
		p.DisplayString = lipgloss.NewStyle().Width(textWidth).Render(p.DisplayString)
	}
	toWrite = strings.Split(p.DisplayString, "\n")


	for i := 0; i < textHeight; i++{
		if p.Style.Border.Enabled{ paneString.WriteString(borderStyle.Render(p.Style.Border.Vertical)) }
		line := ""
		if i < len(toWrite){
			line = toWrite[i]
		}

		s :=lipgloss.NewStyle().MaxWidth(textWidth).Render(line) 
		paneString.WriteString(s)
		visibleWidth := lipgloss.Width(s)
		if visibleWidth < textWidth{
			for range textWidth - visibleWidth{
				paneString.WriteRune(' ')
			}
		}
		

		if p.Style.Border.Enabled{ paneString.WriteString(borderStyle.Render(p.Style.Border.Vertical)) }
		paneString.WriteRune('\n')
	}

	if p.Style.Border.Enabled{
		paneString.WriteString(borderStyle.Render(p.Style.Border.BottomLeft))
		paneString.WriteString(borderStyle.Render(strings.Repeat(p.Style.Border.Horizontal, width - 2)))
		paneString.WriteString(borderStyle.Render(p.Style.Border.BottomRight))
	}

	return paneString.String()
}

func clampI(mn, mx, val int) int {
	if mn > 0 && val < mn { return mn }
	if mx > 0 && val > mx { return mx }
	return val
}


func ResolvePaneSizes(panes []Pane, main, cross int) []Pane {
	if main <= 0 || cross <= 0{
		return nil
	}

	usedMain  := 0

	//first pass for fixed width
	for i, pane := range panes{
		if pane.Style.Basis.Unit == UnitFixed{
			size := clampI(pane.Style.MinBasis, pane.Style.MaxBasis, ResolveSize(pane.Style.Basis, main))
			usedMain += size
			panes[i].Width = size
		}
	}

	usedSize := main - usedMain

	for i, pane := range panes{
		if pane.Style.Basis.Unit == UnitPercent{
			size := clampI(pane.Style.MinBasis, pane.Style.MaxBasis, ResolveSize(pane.Style.Basis, usedSize))
			usedMain += size
			panes[i].Width = size
		}
	}

	usedSize = main - usedMain
	autoCount := 0
	for _, pane := range panes{ if pane.Style.Basis.Unit == UnitAuto{ autoCount++ } }

	if autoCount != 0{
		autoSize := usedSize / autoCount
		for i, pane := range panes{
			if pane.Style.Basis.Unit == UnitAuto{
				size := clampI(pane.Style.MinBasis, pane.Style.MaxBasis,  autoSize)
				usedMain += size
				panes[i].Width = size
			}
		}

	}

	leftOver := main - usedMain

	if leftOver > 0{
		panes[len(panes)-1].Width += leftOver
	}

	for i, pane := range panes{
		size := clampI(pane.Style.MinCrossBasis, pane.Style.MaxCrossBasis, ResolveSize(pane.Style.CrossBasis, cross))
		if size == 0{
			size = cross
		}
		panes[i].Height = size
	}

	return panes
}

func RenderContainer(c PaneContainer, main, cross int) string{
	if c.Panes == nil{
		return ""
	}
	//order
	slices.SortStableFunc(c.Panes, func(i, j Pane) int{
		return cmp.Compare(i.Style.Order, j.Style.Order)
	})
	//getSizes

	strs := []string{}
	c.Panes = ResolvePaneSizes(c.Panes, main, cross)
	for _, pane := range c.Panes{
		strs = append(strs, RenderPane(pane, pane.Width, pane.Height))
	}
	str := ""
	switch c.Style.FlexDirection{
	case DirectionColumn:
		str = lipgloss.JoinVertical(0, strs...)
	default:
		str = lipgloss.JoinHorizontal(0, strs...)
	}

	return str
}

