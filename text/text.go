package text

import (
	"image"
	"image/color"

	"github.com/cs3238-tsuzu/prasoba/transformer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

// Text is a ebiten's Text
type Text struct {
	text string
	f    font.Face

	*transformer.Rect
	offset image.Point
	col    color.Color
}

// NewText returns a new Text
func NewText(t string, f font.Face) *Text {
	bound := text.BoundString(f, t)

	return &Text{
		text:   t,
		f:      f,
		Rect:   transformer.NewRect(image.Rect(0, 0, bound.Dx(), bound.Dy())),
		offset: bound.Min,
	}
}

// Clone returns a cloned Text
func (d *Text) Clone() *Text {
	return &Text{
		text:   d.text,
		f:      d.f,
		Rect:   d.Rect.Clone(),
		offset: d.offset,
		col:    d.col,
	}
}

// WithColor sets a text color
func (d *Text) WithColor(col color.Color) *Text {
	d.col = col

	return d
}

// Center move the center of the rect to the position
func (d *Text) Center(x, y int) *Text {
	d.Rect.Center(x, y)

	return d
}

// From move the rect to the position in the mode
func (d *Text) From(x, y int, mode transformer.PositionMode) *Text {
	d.Rect.From(x, y, mode)

	return d
}

// Draw draws a given text on a given destination image dst.
func (d *Text) Draw(dst *ebiten.Image) {
	pos := d.Rect.Min().Sub(d.offset)

	text.Draw(dst, d.text, d.f, pos.X, pos.Y, d.col)
}
