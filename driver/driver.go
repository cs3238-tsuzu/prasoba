package driver

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// DefaultDriver is a default driver for ebiten
var DefaultDriver = &defaultDriver{}

// Driver is the middleware for ebiten
type Driver interface {
	CursorPosition() image.Point

	IsMouseButtonPressed(key ebiten.MouseButton) bool

	IsMouseButtonJustPressed(key ebiten.MouseButton) bool

	MousePrimaryButton() ebiten.MouseButton

	IsKeyPressed(ebiten.Key) bool

	IsKeyJustPressed(key ebiten.Key) bool

	PressedTouchPositions() []image.Point

	JustPressedTouchPositions() []image.Point
}
