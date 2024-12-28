package dropdown

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Dropdown struct {
	Bounds          rl.Rectangle
	Items           []string
	SelectedIndex   *int
	IsOpen          bool
	BackgroundColor rl.Color
	BorderColor     rl.Color
	TextColor       rl.Color
	HoverColor      rl.Color
	FontSize        int32
	HoverIndex      *int
	MaxVisibleItems int
	ScrollOffset    int
}

func NewDropdown(x, y, width, height float32, items []string) *Dropdown {
	return &Dropdown{
		Bounds:          rl.NewRectangle(x, y, width, height),
		Items:           items,
		SelectedIndex:   nil,
		IsOpen:          false,
		BackgroundColor: rl.White,
		BorderColor:     rl.Black,
		TextColor:       rl.Black,
		HoverColor:      rl.LightGray,
		FontSize:        20,
		HoverIndex:      nil,
		MaxVisibleItems: 5,
		ScrollOffset:    0,
	}
}

func (dd *Dropdown) SetColors(background, border, text, hover rl.Color) {
	dd.BackgroundColor = background
	dd.BorderColor = border
	dd.TextColor = text
	dd.HoverColor = hover
}

func (dd *Dropdown) SetFontSize(size int32) {
	dd.FontSize = size
}

func (dd *Dropdown) SetMaxVisibleItems(count int) {
	dd.MaxVisibleItems = count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func (dd *Dropdown) getItemBounds(index int) rl.Rectangle {
	return rl.NewRectangle(
		dd.Bounds.X,
		dd.Bounds.Y+dd.Bounds.Height+float32(index)*dd.Bounds.Height,
		dd.Bounds.Width,
		dd.Bounds.Height,
	)
}

func (dd *Dropdown) Update() {
	mousePos := rl.GetMousePosition()
	dd.HoverIndex = nil

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.CheckCollisionPointRec(mousePos, dd.Bounds) {
			dd.IsOpen = !dd.IsOpen
		} else if dd.IsOpen {
			visibleItems := min(len(dd.Items), dd.MaxVisibleItems)
			for i := 0; i < visibleItems; i++ {
				itemBounds := dd.getItemBounds(i)
				if rl.CheckCollisionPointRec(mousePos, itemBounds) {
					selected := i + dd.ScrollOffset
					dd.SelectedIndex = &selected
					dd.IsOpen = false
				}
			}
		}
	}
}

func (dd *Dropdown) Draw() {
	rl.DrawRectangleRec(dd.Bounds, dd.BackgroundColor)
	rl.DrawRectangleLinesEx(dd.Bounds, 2.0, dd.BorderColor)

	if dd.SelectedIndex != nil && *dd.SelectedIndex < len(dd.Items) {
		rl.DrawText(
			dd.Items[*dd.SelectedIndex],
			int32(dd.Bounds.X)+5,
			int32(dd.Bounds.Y+(dd.Bounds.Height-float32(dd.FontSize))/2.0),
			dd.FontSize,
			dd.TextColor,
		)
	}

	arrowSize := float32(dd.FontSize) * 0.5
	arrowX := dd.Bounds.X + dd.Bounds.Width - arrowSize - 5.0
	arrowY := dd.Bounds.Y + (dd.Bounds.Height-arrowSize)/2.0
	rl.DrawTriangle(
		rl.NewVector2(arrowX, arrowY),
		rl.NewVector2(arrowX+arrowSize, arrowY),
		rl.NewVector2(arrowX+arrowSize/2.0, arrowY+arrowSize),
		dd.TextColor,
	)

	if dd.IsOpen {
		visibleItems := min(len(dd.Items), dd.MaxVisibleItems)
		for i := 0; i < visibleItems; i++ {
			itemIndex := i + dd.ScrollOffset
			if itemIndex >= len(dd.Items) {
				break
			}

			itemBounds := dd.getItemBounds(i)
			backgroundColor := dd.BackgroundColor
			if dd.HoverIndex != nil && *dd.HoverIndex == itemIndex {
				backgroundColor = dd.HoverColor
			}

			rl.DrawRectangleRec(itemBounds, backgroundColor)
			rl.DrawRectangleLinesEx(itemBounds, 2.0, dd.BorderColor)
			rl.DrawText(
				dd.Items[itemIndex],
				int32(itemBounds.X)+5,
				int32(itemBounds.Y+(itemBounds.Height-float32(dd.FontSize))/2.0),
				dd.FontSize,
				dd.TextColor,
			)
		}
	}
}

