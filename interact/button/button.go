package button

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Bounds            rl.Rectangle
	Label             string
	BackgroundColor   rl.Color
	HoverColor        rl.Color
	PressedColor      rl.Color
	BorderColor       rl.Color
	TextColor         rl.Color
	FontSize          int32
	IsHovered         bool
	IsPressed         bool
	AnimationProgress float32
	Padding           float32
	CornerRadius      float32
	Enabled           bool
}

func NewButton(x, y, width, height float32, label string) *Button {
	return &Button{
		Bounds:            rl.NewRectangle(x, y, width, height),
		Label:             label,
		BackgroundColor:   rl.LightGray,
		HoverColor:        rl.NewColor(200, 200, 200, 255),
		PressedColor:      rl.DarkGray,
		BorderColor:       rl.Black,
		TextColor:         rl.Black,
		FontSize:          20,
		IsHovered:         false,
		IsPressed:         false,
		AnimationProgress: 0.0,
		Padding:           5.0,
		CornerRadius:      5.0,
		Enabled:           true,
	}
}

func (b *Button) SetColors(background, hover, pressed, border, text rl.Color) {
	b.BackgroundColor = background
	b.HoverColor = hover
	b.PressedColor = pressed
	b.BorderColor = border
	b.TextColor = text
}

func (b *Button) SetFontSize(size int32) {
	b.FontSize = size
}

func (b *Button) SetCornerRadius(radius float32) {
	b.CornerRadius = radius
}

func (b *Button) SetPadding(padding float32) {
	b.Padding = padding
}

func (b *Button) SetEnabled(enabled bool) {
	b.Enabled = enabled
}

func (b *Button) Update() {
	if !b.Enabled {
		b.IsHovered = false
		b.IsPressed = false
		return
	}

	mousePos := rl.GetMousePosition()
	b.IsHovered = rl.CheckCollisionPointRec(mousePos, b.Bounds)

	if b.IsHovered && rl.IsMouseButtonDown(rl.MouseLeftButton) {
		b.IsPressed = true
	} else {
		b.IsPressed = false
	}

	targetProgress := float32(0.0)
	if b.IsPressed {
		targetProgress = 1.0
	} else if b.IsHovered {
		targetProgress = 0.5
	}

	animationSpeed := float32(8.0)
	frameTime := rl.GetFrameTime()
	if b.AnimationProgress < targetProgress {
		b.AnimationProgress += frameTime * animationSpeed
		if b.AnimationProgress > targetProgress {
			b.AnimationProgress = targetProgress
		}
	} else if b.AnimationProgress > targetProgress {
		b.AnimationProgress -= frameTime * animationSpeed
		if b.AnimationProgress < targetProgress {
			b.AnimationProgress = targetProgress
		}
	}
}

func (b *Button) Draw() {
	currentColor := b.BackgroundColor
	if !b.Enabled {
		currentColor = rl.Fade(b.BackgroundColor, 0.5)
	} else {
		if b.AnimationProgress <= 0.5 {
			t := b.AnimationProgress * 2.0
			currentColor = rl.NewColor(
				uint8(float32(b.HoverColor.R-b.BackgroundColor.R)*t)+b.BackgroundColor.R,
				uint8(float32(b.HoverColor.G-b.BackgroundColor.G)*t)+b.BackgroundColor.G,
				uint8(float32(b.HoverColor.B-b.BackgroundColor.B)*t)+b.BackgroundColor.B,
				uint8(float32(b.HoverColor.A-b.BackgroundColor.A)*t)+b.BackgroundColor.A,
			)
		} else {
			t := (b.AnimationProgress - 0.5) * 2.0
			currentColor = rl.NewColor(
				uint8(float32(b.PressedColor.R-b.HoverColor.R)*t)+b.HoverColor.R,
				uint8(float32(b.PressedColor.G-b.HoverColor.G)*t)+b.HoverColor.G,
				uint8(float32(b.PressedColor.B-b.HoverColor.B)*t)+b.HoverColor.B,
				uint8(float32(b.PressedColor.A-b.HoverColor.A)*t)+b.HoverColor.A,
			)
		}
	}

	rl.DrawRectangleRounded(b.Bounds, b.CornerRadius, 8, currentColor)

	borderThickness := float32(2.0)
	if b.IsPressed {
		borderThickness = 3.0
	}
	rl.DrawRectangleRoundedLines(b.Bounds, b.CornerRadius, int32(borderThickness), b.BorderColor)

	textWidth := float32(rl.MeasureText(b.Label, b.FontSize))
	textX := b.Bounds.X + (b.Bounds.Width-textWidth)/2.0
	textY := b.Bounds.Y + (b.Bounds.Height-float32(b.FontSize))/2.0

	offsetX, offsetY := float32(0), float32(0)
	if b.IsPressed {
		offsetX, offsetY = 1.0, 1.0
	}

	textColor := b.TextColor
	if !b.Enabled {
		textColor = rl.Fade(b.TextColor, 0.5)
	}

	rl.DrawText(b.Label, int32(textX+offsetX), int32(textY+offsetY), b.FontSize, textColor)
}

func (b *Button) IsClicked() bool {
	return b.Enabled && b.IsHovered && rl.IsMouseButtonReleased(rl.MouseLeftButton)
}
