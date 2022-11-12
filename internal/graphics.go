package internal

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

//counterfeiter:generate . Context
type Context interface {
	Image() image.Image
	DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align gg.Align)
	LoadFontFace(path string, points float64) error
	SetColor(c color.Color)
}

//counterfeiter:generate . ContextMaker
type ContextMaker func(i image.Image) Context

var NewContext ContextMaker = func(i image.Image) Context {
	return gg.NewContextForImage(i)
}
