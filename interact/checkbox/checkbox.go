// SPDX-License-Identifier: MIT
package checkbox

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Checkbox struct {
	Bounds            rl.Rectangle
	checked           bool
	BackgroundColor   rl.Color
	CheckColor        rl.Color
	BorderColor       rl.Color
	HoverColor        rl.Color
	Label             string
	LabelColor        rl.Color
	FontSize          int32
	IsHovered         bool
	AnimationProgress float32
	IsClicked         bool
}

func NewCheckbox(x, y, size float32, label string) *Checkbox {
	return &Checkbox{
		Bounds:            rl.NewRectangle(x, y, size, size),
		checked:           false,
		BackgroundColor:   rl.White,
		CheckColor:        rl.Green,
		BorderColor:       rl.Black,
		HoverColor:        rl.NewColor(245, 245, 245, 255),
		Label:             label,
		LabelColor:        rl.Black,
		FontSize:          20,
		IsHovered:         false,
		AnimationProgress: 0.0,
		IsClicked:         false,
	}
}

func (cb *Checkbox) SetColors(background, check, border, hover, label rl.Color) {
	cb.BackgroundColor = background
	cb.CheckColor = check
	cb.BorderColor = border
	cb.HoverColor = hover
	cb.LabelColor = label
}

func (cb *Checkbox) SetFontSize(size int32) {
	cb.FontSize = size
}

func (cb *Checkbox) Update() {
	mousePos := rl.GetMousePosition()
	cb.IsHovered = rl.CheckCollisionPointRec(mousePos, cb.Bounds)

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && cb.IsHovered {
		cb.IsClicked = true
		cb.checked = !cb.checked
		cb.AnimationProgress = 0.0
	}

	animationSpeed := float32(4.0)
	frameTime := rl.GetFrameTime()

	if cb.checked && cb.AnimationProgress < 1.0 {
		cb.AnimationProgress += frameTime * animationSpeed
		if cb.AnimationProgress > 1.0 {
			cb.AnimationProgress = 1.0
		}
	} else if !cb.checked && cb.AnimationProgress > 0.0 {
		cb.AnimationProgress -= frameTime * animationSpeed
		if cb.AnimationProgress < 0.0 {
			cb.AnimationProgress = 0.0
		}
	}

	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		cb.IsClicked = false
	}
}

func (cb *Checkbox) Draw() {
	backgroundColor := cb.BackgroundColor
	if cb.IsHovered {
		backgroundColor = cb.HoverColor
	}

	rl.DrawRectangleRec(cb.Bounds, backgroundColor)

	borderThickness := float32(2.0)
	if cb.IsClicked {
		borderThickness = 3.0
	}

	rl.DrawRectangleLinesEx(cb.Bounds, borderThickness, cb.BorderColor)

	if cb.AnimationProgress > 0.0 {
		padding := cb.Bounds.Width * 0.2
		checkBounds := rl.NewRectangle(
			cb.Bounds.X+padding,
			cb.Bounds.Y+padding,
			cb.Bounds.Width-padding*2.0,
			cb.Bounds.Height-padding*2.0,
		)

		centerX := checkBounds.X + checkBounds.Width/2.0
		centerY := checkBounds.Y + checkBounds.Height/2.0
		size := checkBounds.Width / 2.0 * cb.AnimationProgress

		points := [3]rl.Vector2{
			rl.NewVector2(centerX-size, centerY),
			rl.NewVector2(centerX, centerY+size),
			rl.NewVector2(centerX+size, centerY-size),
		}

		checkColor := rl.NewColor(
			cb.CheckColor.R,
			cb.CheckColor.G,
			cb.CheckColor.B,
			uint8(float32(cb.CheckColor.A) * cb.AnimationProgress),
		)

		rl.DrawLineEx(points[0], points[1], 2.0, checkColor)
		rl.DrawLineEx(points[1], points[2], 2.0, checkColor)
	}

	labelX := cb.Bounds.X + cb.Bounds.Width + 10.0
	labelY := cb.Bounds.Y + (cb.Bounds.Height-float32(cb.FontSize))/2.0

	rl.DrawText(cb.Label, int32(labelX), int32(labelY), cb.FontSize, cb.LabelColor)
}

func (cb *Checkbox) IsChecked() bool {
	return cb.checked
}

func (cb *Checkbox) SetChecked(checked bool) {
	if checked != cb.checked {
		cb.checked = checked
		cb.AnimationProgress = 0.0
	}
}

func (cb *Checkbox) Toggle() {
	cb.SetChecked(!cb.IsChecked())
}
