// SPDX-License-Identifier: MIT
package interact

import (
    "github.com/gen2brain/raylib-go/raylib"
    "rayinteract_go/interact/button"
    "rayinteract_go/interact/checkbox"
    "rayinteract_go/interact/dropdown"
    "rayinteract_go/interact/textfield"
)

// Theme holds the common colors and styles for UI components
type Theme struct {
	Background rl.Color
	Border     rl.Color
	Text       rl.Color
	Hover      rl.Color
	Pressed    rl.Color
	Check      rl.Color
	FontSize   int32
}

// DefaultTheme returns a default color theme
func DefaultTheme() Theme {
	return Theme{
		Background: rl.White,
		Border:     rl.Black,
		Text:       rl.Black,
		Hover:      rl.LightGray,
		Pressed:    rl.DarkGray,
		Check:      rl.Green,
		FontSize:   20,
	}
}

// NewButton creates a new button with the theme settings
func NewButton(x, y, width, height float32, label string, theme Theme) *button.Button {
	btn := button.NewButton(x, y, width, height, label)
	btn.SetColors(theme.Background, theme.Hover, theme.Pressed, theme.Border, theme.Text)
	btn.SetFontSize(theme.FontSize)
	return btn
}

// NewTextField creates a new text field with the theme settings
func NewTextField(x, y, width, height float32, maxLength int, theme Theme) *textfield.TextField {
	tf := textfield.NewTextField(x, y, width, height, maxLength)
	tf.SetColors(theme.Background, theme.Border, theme.Text)
	tf.SetFontSize(theme.FontSize)
	return tf
}

// NewCheckbox creates a new checkbox with the theme settings
func NewCheckbox(x, y, size float32, label string, theme Theme) *checkbox.Checkbox {
	cb := checkbox.NewCheckbox(x, y, size, label)
	cb.SetColors(theme.Background, theme.Check, theme.Border, theme.Hover, theme.Text)
	cb.SetFontSize(theme.FontSize)
	return cb
}

// NewDropdown creates a new dropdown with the theme settings
func NewDropdown(x, y, width, height float32, items []string, theme Theme) *dropdown.Dropdown {
	dd := dropdown.NewDropdown(x, y, width, height, items)
	dd.SetColors(theme.Background, theme.Border, theme.Text, theme.Hover)
	dd.SetFontSize(theme.FontSize)
	return dd
}

// UpdateAll updates all provided UI components
func UpdateAll(components ...interface{}) {
	for _, component := range components {
		switch c := component.(type) {
		case *button.Button:
			c.Update()
		case *textfield.TextField:
			c.Update()
		case *checkbox.Checkbox:
			c.Update()
		case *dropdown.Dropdown:
			c.Update()
		}
	}
}

// DrawAll draws all provided UI components
func DrawAll(components ...interface{}) {
	for _, component := range components {
		switch c := component.(type) {
		case *button.Button:
			c.Draw()
		case *textfield.TextField:
			c.Draw()
		case *checkbox.Checkbox:
			c.Draw()
		case *dropdown.Dropdown:
			c.Draw()
		}
	}
}