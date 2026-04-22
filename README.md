# TeaPane

TeaPane is a small Go library for rendering terminal panes with a flexbox-inspired layout model.  
It is designed to pair well with the [Bubble Tea](https://github.com/charmbracelet/bubbletea) ecosystem.

## Features

- Render bordered text panes in the terminal
- Build pane containers with row/column direction
- Control pane sizing with:
  - fixed values
  - percentages
  - automatic distribution
- Basic pane ordering and wrapping behavior
- Lip Gloss-compatible styling primitives

## Installation

```bash
go get github.com/duckisam/TeaPane
```

## Quick Start

```go
package main

import (
	"fmt"

	pane "github.com/duckisam/TeaPane"
)

func main() {
	p1 := pane.NewPane(100, 50, true)
	p1.Style.WrapText = true
	p1.Style.Basis = pane.Percent(70)
	p1.DisplayString = "Hello from pane one."

	p2 := pane.NewPane(100, 18, true)
	p2.Style.WrapText = true
	p2.Style.Basis = pane.Percent(30)
	p2.DisplayString = "Hello from pane two."

	container := pane.NewContainer(pane.NewContainerStyle(), p1, p2)
	fmt.Print(pane.RenderContainer(container, 100, 25))
}
```

## Core Types

- `Pane`: single drawable pane with content and style
- `PaneContainer`: groups panes and renders them together
- `PaneStyle`: per-pane layout and border behavior
- `ContainerStyle`: container-level flex direction and layout options
- `Size`: helper type for `Fixed`, `Percent`, and `Auto` sizing

## Development

Run checks:

```bash
go test ./...
```

## License

This project is licensed under the terms in the [LICENSE](./LICENSE) file.
