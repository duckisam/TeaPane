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
		var str strings.Builder
		count := 0
		for i := 0; i < len(p.DisplayString); i++{
			if count == textWidth + 1 ||  p.DisplayString[i] == '\n'{
				count = 0
				str.WriteRune('\n')
			}
			if p.DisplayString[i] == ' ' && count == 0{ continue }
			str.WriteRune(rune(p.DisplayString[i]))
			count++
		}
		p.DisplayString = str.String()
	}
	toWrite = strings.Split(p.DisplayString, "\n")

	runeLines := make([][]rune, len(toWrite))
	for i, line := range toWrite{
		runeLines[i] = []rune(line)
	}

	for i := 0; i < textHeight; i++{
		if p.Style.Border.Enabled{ paneString.WriteString(borderStyle.Render(p.Style.Border.Vertical)) }
		for j := 0; j < textWidth; j++{
			if i < len(toWrite) && j < len(toWrite[i]){
				paneString.WriteRune(runeLines[i][j])
			}else{
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

	for i, pane := range panes{
		size := clampI(pane.Style.MinCrossBasis, pane.Style.MaxBasis, ResolveSize(pane.Style.CrossBasis, cross))
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

