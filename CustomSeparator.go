package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// // CustomSeparator is a custom separator widget with a specified color.
// type CustomSeparator struct {
//     *canvas.Line
// }

// // NewCustomSeparator creates a new instance of CustomSeparator with the specified color.
// func NewCustomSeparator(color color.Color) *CustomSeparator {
//     separator := &CustomSeparator{
//         Line: canvas.NewLine(color),
//     }
//     separator.StrokeWidth = 3 // Adjust line thickness as needed
//     return separator
// }




type CustomSeparator struct {
	widget.BaseWidget
	color color.Color
	strokeWidth float32
}

func NewCustomSeparator(color color.Color, strokeWidth float32) *CustomSeparator {
	separator := &CustomSeparator{
		color: color,
		strokeWidth: strokeWidth,
	}
	separator.ExtendBaseWidget(separator)

	return separator
}

func (item *CustomSeparator) CreateRenderer() fyne.WidgetRenderer {
	line := canvas.NewLine(item.color)
    line.StrokeWidth = item.strokeWidth
	return widget.NewSimpleRenderer(line)
}