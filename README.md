# teapane

> ⚠️ This README was generated with the assistance of AI (Claude by Anthropic) based on the current source code and design discussions with the author.

---

**teapane** is a Go library for building terminal UI layouts using a flexbox-inspired model. It is designed to work alongside [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss), giving you a familiar CSS-like mental model for arranging panes in a terminal window.

> **Status: Beta** — The core rendering pipeline is working. The full flexbox feature set (grow/shrink, justify-content, align-items, wrapping, and gap) is actively being developed.

---

## Why teapane?

Building complex terminal layouts with raw Lip Gloss requires a lot of manual size math. teapane abstracts that away by letting you describe *how* panes should relate to each other — much like CSS flexbox — rather than calculating exact pixel (cell) positions yourself. If you know how to write a flexbox layout in CSS, the concepts here will feel immediately familiar.

---

## Core Concepts

### Pane

A `Pane` is the basic building block. It holds a string of content to display, a resolved width and height, and a `PaneStyle` that controls how it looks and how it behaves inside a container.

```go
pane := teapane.NewPane(40, 10, true) // width, height, hasBorder
pane.DisplayString = "Hello, terminal!"
```

### PaneContainer

A `PaneContainer` holds a slice of `Pane`s and a `ContainerStyle` that controls how they are laid out — the flex direction, wrapping, alignment, and so on.

```go
container := teapane.NewContainer(
    teapane.NewContainerStyle(),
    paneA,
    paneB,
    paneC,
)
```

### Rendering

To render a container at a given size, call `RenderContainer`. To render a standalone pane, call `RenderPane` or use the `View()` method on a `Pane` directly.

```go
output := teapane.RenderContainer(container, totalWidth, totalHeight)
fmt.Print(output)
```

---

## Sizing System

teapane uses a `Size` struct with three unit modes, mirroring CSS concepts:

`Fixed(n)` gives the pane an exact cell count regardless of the container size. `Percent(n)` gives the pane a percentage of the available space. `Auto()` means the pane will share whatever space is left over equally with other auto-sized panes.

```go
// This pane always takes up exactly 20 columns.
pane.Style.Basis = teapane.Fixed(20)

// This pane takes up 50% of the available width.
pane.Style.Basis = teapane.Percent(50)

// This pane takes up an equal share of whatever is left.
pane.Style.Basis = teapane.Auto()
```

You can also set `MinBasis` and `MaxBasis` on a pane to clamp its resolved size, similar to `min-width` and `max-width` in CSS.

---

## Borders

Borders are configured through `PaneBorder` and are included in the pane's total size — the content area shrinks accordingly, just like `box-sizing: border-box` in CSS. A default border style is provided:

```go
// Uses the built-in rounded box-drawing border in amber.
pane.Style.Border = teapane.DefaultBorder
```

You can customise the corner and edge characters and the Lip Gloss color to create any border style you like.

---

## Flex Direction

The container's `FlexDirection` controls whether panes flow left-to-right (`DirectionRow`) or top-to-bottom (`DirectionColumn`).

```go
style := teapane.NewContainerStyle()
style.FlexDirection = teapane.DirectionColumn
```

---

## Ordering

Panes are sorted by their `Style.Order` field before rendering, just like the CSS `order` property. Lower values appear first. This lets you reorder panes in the layout without changing their position in the slice.

---

## Styles Reference

### PaneStyle fields (current)

`Grow` and `Shrink` are defined but not yet active in the layout engine — they will control how panes expand or contract to fill leftover space. `Basis` and `CrossBasis` control main-axis and cross-axis sizing. `MinBasis`, `MaxBasis`, `MinCrossBasis`, and `MaxCrossBasis` clamp those sizes. `Order` controls render order. `WrapText` enables soft word-wrapping of the pane's content. `AlignSelf` will override the container's `AlignItems` for this specific pane. `Border` holds the pane's border configuration.

### ContainerStyle fields (current)

`FlexDirection` sets the flow axis. `FlexWrap`, `Justify`, `AlignItems`, `AlignContent`, `GapRow`, and `GapColumn` are defined and will be fully active once the layout engine is extended (see Roadmap below).

---

## Roadmap

The following features are defined in the type system but not yet implemented in the layout engine. They are planned for the next development phase, following the CSS flexbox algorithm closely.

**flex-grow / flex-shrink** will distribute leftover space (or absorb overflow) proportionally among panes that opt in, using a weighted shrink algorithm identical to how browsers handle it.

**justify-content** will control spacing along the main axis — centering items, pushing them to the ends, or distributing space between and around them.

**align-items / align-self** will position panes along the cross axis within their flex line, with per-pane overrides via `AlignSelf`.

**flex-wrap** will allow panes to break onto multiple lines when they overflow the container, with `WrapReverse` for reverse line stacking.

**gap** (row and column) will insert fixed space between panes without affecting `justify-content` free-space calculations.

**Nested containers** are a longer-term goal. The plan is to introduce a `FlexItem` interface and a separate `FlexItemStyle` struct so that a `PaneContainer` can itself be a child of another container, holding both a `ContainerStyle` (for its children) and a `FlexItemStyle` (for its own parent).

---

## Dependencies

teapane depends on [Lip Gloss](https://github.com/charmbracelet/lipgloss) for color rendering and for `JoinHorizontal` / `JoinVertical` composition. It is intended to be used with [Bubble Tea](https://github.com/charmbracelet/bubbletea), though it has no hard dependency on the Bubble Tea runtime itself — you can call `RenderContainer` from any context.

---

## License

teapane is released under the [MIT License](https://opensource.org/licenses/MIT). You are free to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the software, provided the original copyright notice and this permission notice are included in all copies or substantial portions of the software.
