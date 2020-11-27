package driver

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type defaultDriver struct {
	primary ebiten.MouseButton
}

var _ Driver = &defaultDriver{}

func (d *defaultDriver) CursorPosition() image.Point {
	return image.Pt(ebiten.CursorPosition())
}

func (d *defaultDriver) IsMouseButtonPressed(key ebiten.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(key)
}

func (d *defaultDriver) IsMouseButtonJustPressed(key ebiten.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(key)
}

func (d *defaultDriver) MousePrimaryButton() ebiten.MouseButton {
	btn := d.primary
	if d.primary == 0 {
		d.primary = ebiten.MouseButtonLeft
	}

	return btn
}

func (d *defaultDriver) IsKeyPressed(key ebiten.Key) bool {
	return ebiten.IsKeyPressed(key)
}

func (d *defaultDriver) IsKeyJustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}

func (d *defaultDriver) touchPositions(ids []ebiten.TouchID) []image.Point {
	positions := make([]image.Point, len(ids))

	for i := range positions {
		positions[i] = image.Pt(ebiten.TouchPosition(ids[i]))
	}

	return positions
}

func (d *defaultDriver) PressedTouchPositions() []image.Point {
	return d.touchPositions(ebiten.TouchIDs())
}

func (d *defaultDriver) JustPressedTouchPositions() []image.Point {
	return d.touchPositions(inpututil.JustPressedTouchIDs())
}
