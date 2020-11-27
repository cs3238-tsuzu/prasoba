package transformer

import (
	"image"

	"github.com/cs3238-tsuzu/prasoba/driver"
	"github.com/hajimehoshi/ebiten/v2"
)

// Rect is a image.Rectangle with the capability to translate
type Rect struct {
	rect image.Rectangle
}

// NewRect returns a new Rect
func NewRect(r image.Rectangle) *Rect {
	return &Rect{
		rect: r,
	}
}

// Clone returns a cloned Rect
func (r *Rect) Clone() *Rect {
	return &Rect{
		rect: r.rect,
	}
}

// Center move the center of the rect to the position
func (r *Rect) Center(x, y int) *Rect {
	return r.From(x, y, Center)
}

// From move the rect to the position in the mode
func (r *Rect) From(x, y int, mode PositionMode) *Rect {
	x, y = CalcRenderedPosFromRectangle(
		x, y, r.rect, mode,
	)

	r.rect = RectangleFrom(x, y, r.rect)

	return r
}

// Size returns Rectangle size
func (r *Rect) Size() image.Point {
	return r.rect.Size()
}

// Min returns the minimum position of the rectangle
func (r *Rect) Min() image.Point {
	return r.rect.Min
}

// Max returns the maximum position of the rectangle
func (r *Rect) Max() image.Point {
	return r.rect.Max
}

func (r *Rect) Hover() bool {
	pos := driver.DefaultDriver.CursorPosition()

	if pos.In(r.rect) {
		return true
	}

	for _, pos := range driver.DefaultDriver.PressedTouchPositions() {
		if pos.In(r.rect) {
			return true
		}
	}

	return false
}

func (r *Rect) ClickedBy(key ebiten.MouseButton) bool {
	pos := driver.DefaultDriver.CursorPosition()

	if pos.In(r.rect) && driver.DefaultDriver.IsMouseButtonJustPressed(key) {
		return true
	}

	for _, pos := range driver.DefaultDriver.JustPressedTouchPositions() {
		if pos.In(r.rect) {
			return true
		}
	}

	return false
}

func (r *Rect) Clicked() bool {
	return r.ClickedBy(driver.DefaultDriver.MousePrimaryButton())
}

func (r *Rect) PressedBy(key ebiten.MouseButton) bool {
	pos := driver.DefaultDriver.CursorPosition()

	if pos.In(r.rect) && driver.DefaultDriver.IsMouseButtonPressed(key) {
		return true
	}

	for _, pos := range driver.DefaultDriver.PressedTouchPositions() {
		if pos.In(r.rect) {
			return true
		}
	}

	return false
}

func (r *Rect) Pressed() bool {
	return r.PressedBy(driver.DefaultDriver.MousePrimaryButton())
}
