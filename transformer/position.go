package transformer

import "image"

type PositionMode int

const (
	TopLeft PositionMode = iota
	TopCenter
	TopRight

	CenterLeft
	Center
	CenterRight

	BottomLeft
	BottomCenter
	BottomRight
)

// CalcRenderedPos calculates the bottom-left position to render at from the transformed position in that specified mode
func CalcRenderedPos(fx, fy, width, height int, mode PositionMode) (nx, ny int) {
	dx, dy := int(mode%3), int(mode/3)

	y := fy - height*dy/2
	x := fx - width*dx/2

	return x, y
}

// CalcRenderedPosFromRectangle calculates the bottom-left position to render at from the transformed position in that specified mode
func CalcRenderedPosFromRectangle(fx, fy int, r image.Rectangle, mode PositionMode) (nx, ny int) {
	return CalcRenderedPos(fx, fy, r.Dx(), r.Dy(), mode)
}
