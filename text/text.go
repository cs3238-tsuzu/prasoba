package text

import (
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
	col color.Color
}

// NewText returns a new Text
func NewText(t string, f font.Face) *Text {
	return &Text{
		text: t,
		f:    f,
		Rect: transformer.NewRect(text.BoundString(f, t)),
	}
}

// Clone returns a cloned Text
func (d *Text) Clone() *Text {
	return &Text{
		text: d.text,
		f:    d.f,
		Rect: d.Rect.Clone(),
		col:  d.col,
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
	text.Draw(dst, d.text, d.f, d.Rect.Min().X, d.Rect.Min().Y, d.col)
}
