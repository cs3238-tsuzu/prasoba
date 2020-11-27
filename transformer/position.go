package transformer

import "image"

type PositionMode int

const (
	BottomLeft PositionMode = iota
	BottomCenter
	BottomRight

	CenterLeft
	Center
	CenterRight

	TopLeft
	TopCenter
	TopRight
)

// CalcRenderedPos calculates the bottom-left position to render at from the transformed position in that specified mode
func CalcRenderedPos(fx, fy, width, height int, mode PositionMode) (nx, ny int) {
	dx, dy := int(mode%3), int(mode/3)

	y := fy - height*dy/2
	x := fx - width*dx/2

	return y, x
}

// CalcRenderedPosFromRectangle calculates the bottom-left position to render at from the transformed position in that specified mode
func CalcRenderedPosFromRectangle(fx, fy int, r image.Rectangle, mode PositionMode) (nx, ny int) {
	return CalcRenderedPos(fx, fy, r.Dx(), r.Dy(), mode)
}
