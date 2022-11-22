package internal

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

//counterfeiter:generate . Context
type Context interface {
	Image() image.Image
	DrawString(s string, x, y float64)
	DrawStringAnchored(s string, x, y, ax, ay float64)
	DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align gg.Align)
	LoadFontFace(path string, points float64) error
	MeasureString(s string) (w, h float64)
	MeasureMultilineString(s string, lineSpacing float64) (w, h float64)
	SetColor(c color.Color)
	WordWrap(s string, w float64) []string
}

//counterfeiter:generate . ContextMaker
type ContextMaker func(i image.Image) Context

var NewContext ContextMaker = func(i image.Image) Context {
	return gg.NewContextForImage(i)
}
