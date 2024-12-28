package textfield

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type TextField struct {
	Bounds             rl.Rectangle
	Text               string
	MaxLength          int
	BackgroundColor    rl.Color
	BorderColor        rl.Color
	TextColor          rl.Color
	FontSize           int32
	IsActive           bool
	CursorPosition     int
	CursorBlinkTimer   float32
	BackspaceHoldTimer float32
}

func NewTextField(x, y, width, height float32, maxLength int) *TextField {
	return &TextField{
		Bounds:             rl.NewRectangle(x, y, width, height),
		Text:               "",
		MaxLength:          maxLength,
		BackgroundColor:    rl.White,
		BorderColor:        rl.Black,
		TextColor:          rl.Black,
		FontSize:           20,
		IsActive:           false,
		CursorPosition:     0,
		CursorBlinkTimer:   0.0,
		BackspaceHoldTimer: 0.0,
	}
}

func (tf *TextField) SetColors(background, border, text rl.Color) {
	tf.BackgroundColor = background
	tf.BorderColor = border
	tf.TextColor = text
}

func (tf *TextField) SetFontSize(fontSize int32) {
	tf.FontSize = fontSize
}

func (tf *TextField) Update() {
	tf.CursorBlinkTimer += rl.GetFrameTime()
	if tf.CursorBlinkTimer >= 1.0 {
		tf.CursorBlinkTimer = 0.0
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePos := rl.GetMousePosition()
		tf.IsActive = rl.CheckCollisionPointRec(mousePos, tf.Bounds)
	}

	if tf.IsActive {
		for key := rl.GetCharPressed(); key > 0; key = rl.GetCharPressed() {
			if len(tf.Text) < tf.MaxLength {
				tf.Text = tf.Text[:tf.CursorPosition] + string(key) + tf.Text[tf.CursorPosition:]
				tf.CursorPosition++
			}
		}

		if rl.IsKeyPressed(rl.KeyBackspace) {
			if tf.CursorPosition > 0 {
				tf.Text = tf.Text[:tf.CursorPosition-1] + tf.Text[tf.CursorPosition:]
				tf.CursorPosition--
			}
		}

		if rl.IsKeyPressed(rl.KeyLeft) && tf.CursorPosition > 0 {
			tf.CursorPosition--
		}

		if rl.IsKeyPressed(rl.KeyRight) && tf.CursorPosition < len(tf.Text) {
			tf.CursorPosition++
		}

		if rl.IsKeyPressed(rl.KeyHome) {
			tf.CursorPosition = 0
		}

		if rl.IsKeyPressed(rl.KeyEnd) {
			tf.CursorPosition = len(tf.Text)
		}

		if rl.IsKeyDown(rl.KeyBackspace) {
			tf.BackspaceHoldTimer += rl.GetFrameTime()
			if tf.BackspaceHoldTimer > 0.5 {
				tf.BackspaceHoldTimer = 1.0
				if tf.CursorPosition > 0 {
					tf.Text = tf.Text[:tf.CursorPosition-1] + tf.Text[tf.CursorPosition:]
					tf.CursorPosition--
				}
			}
		} else {
			tf.BackspaceHoldTimer = 0.0
		}
	}
}

func (tf *TextField) Draw() {
	rl.DrawRectangleRec(tf.Bounds, tf.BackgroundColor)

	borderColor := tf.BorderColor
	if tf.IsActive {
		borderColor = rl.Red
	}

	rl.DrawRectangleLinesEx(tf.Bounds, 2.0, borderColor)

	textY := int32(tf.Bounds.Y + (tf.Bounds.Height-float32(tf.FontSize))/2.0)
	rl.DrawText(tf.Text, int32(tf.Bounds.X)+5, textY, tf.FontSize, tf.TextColor)

	if tf.IsActive && tf.CursorBlinkTimer < 0.5 {
		textWidth := float32(rl.MeasureText(tf.Text[:tf.CursorPosition], tf.FontSize))
		rl.DrawLine(
			int32(tf.Bounds.X+5+textWidth),
			textY,
			int32(tf.Bounds.X+5+textWidth),
			textY+tf.FontSize,
			tf.TextColor,
		)
	}
}

func (tf *TextField) GetText() string {
	return tf.Text
}

func (tf *TextField) SetValue(value string) {
	tf.Text = value[:min(len(value), tf.MaxLength)]
	tf.CursorPosition = len(tf.Text)
}

func (tf *TextField) Activate() {
	tf.IsActive = true
	tf.CursorPosition = len(tf.Text)
}

func (tf *TextField) Deactivate() {
	tf.IsActive = false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
