# RayInteract-Go

A simple and lightweight UI component library for Raylib-Go, providing essential interactive elements with customizable themes and smooth animations.

## Features

- Easy-to-use UI components:
    - Buttons with hover and click animations
    - Checkboxes with smooth check/uncheck animations
    - Text fields with cursor support
    - Dropdowns with scrollable lists
- Customizable themes
- Smooth animations and state transitions
- Written in pure Go using Raylib-Go

## Installation

```bash
go get -u github.com/yourusername/rayinteract_go
```

## Quick Start

```go
package main

import (
        "github.com/gen2brain/raylib-go/raylib"
        "rayinteract_go/interact"
)

func main() {
        rl.InitWindow(800, 600, "RayInteract Demo")
        defer rl.CloseWindow()
        
        theme := interact.DefaultTheme()
        
        // Create UI components
        button := interact.NewButton(100, 100, 200, 40, "Click Me!", theme)
        checkbox := interact.NewCheckbox(100, 160, 20, "Enable Option", theme)
        textfield := interact.NewTextField(100, 200, 200, 40, 32, theme)
        
        for !rl.WindowShouldClose() {
                // Update
                interact.UpdateAll(button, checkbox, textfield)
                
                // Draw
                rl.BeginDrawing()
                rl.ClearBackground(rl.RayWhite)
                
                interact.DrawAll(button, checkbox, textfield)
                
                rl.EndDrawing()
        }
}
```

## Documentation

### Theme

The `Theme` struct defines the appearance of all UI components:

```go
type Theme struct {
        Background rl.Color // Background color
        Border     rl.Color // Border color
        Text       rl.Color // Text color
        Hover      rl.Color // Hover state color
        Pressed    rl.Color // Pressed state color
        Check      rl.Color // Checkbox check color
        FontSize   int32    // Font size for text
}
```

### Button

```go
// Create a new button
button := interact.NewButton(x, y, width, height, label, theme)

// Check if button was clicked
if button.IsClicked() {
        // Handle click
}
```

### Checkbox

```go
// Create a new checkbox
checkbox := interact.NewCheckbox(x, y, size, label, theme)

// Get checkbox state
if checkbox.IsChecked() {
        // Handle checked state
}
```

### TextField

```go
// Create a new text field
textfield := interact.NewTextField(x, y, width, height, maxLength, theme)

// Get text content
text := textfield.GetText()

// Set text content
textfield.SetValue("New text")
```

### Dropdown

```go
// Create a new dropdown
items := []string{"Option 1", "Option 2", "Option 3"}
dropdown := interact.NewDropdown(x, y, width, height, items, theme)

// Get selected index
if dropdown.SelectedIndex != nil {
        selectedItem := items[*dropdown.SelectedIndex]
}
```

## License

MIT License - feel free to use this library in your own projects.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.