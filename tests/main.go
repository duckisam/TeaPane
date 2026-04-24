package main

import (
	"fmt"

	pane "github.com/duckisam/TeaPane"
)

func main(){
	p1 := pane.NewPane(100, 50, true)
	p1.DisplayString = "hello\ntext\nThis ██xt should wrap around as its is too long"
	p1.Style.Basis = pane.Percent(70)
	p2 := pane.NewPane(100, 18, true)
	p2.Style.Basis = pane.Percent(30)
	p2.DisplayString = "hello\ntext\nThis text should wrap around as its is too long 2"

	c := pane.NewContainer(pane.NewContainerStyle(), p1, p2)

	fmt.Print(pane.RenderContainer(c, 100, 25))

	

}
