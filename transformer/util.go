package transformer

import "image"

// Rectangle creates a image.Rectangle by width/height
func Rectangle(x, y, w, h int) image.Rectangle {
	return image.Rect(x, y, x+w, y+h)
}

// RectangleFrom returns a rectangle from the specified position
func RectangleFrom(x, y int, r image.Rectangle) image.Rectangle {
	return Rectangle(x, y, r.Dx(), r.Dy())
}
