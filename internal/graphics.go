package internal

import (
	"fmt"
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

//counterfeiter:generate . Context
type Context interface {
	Image() image.Image
	DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align gg.Align)
	LoadFontFace(path string, points float64) error
	MeasureString(s string) (w, h float64)
	MeasureMultilineString(s string, lineSpacing float64) (w, h float64)
	SetColor(c color.Color)
}

//counterfeiter:generate . ContextMaker
type ContextMaker func(i image.Image) Context

var NewContext ContextMaker = func(i image.Image) Context {
	return gg.NewContextForImage(i)
}

//counterfeiter:generate . TextFitter
type TextFitter func(context Context, text, font string, width, height int) (float64, error)

var FitText TextFitter = func(context Context, text, font string, width, height int) (float64, error) {
	var size float64
	for size = 0; size < 1000; size += 1 {
		err := context.LoadFontFace(font, size)
		if err != nil {
			return 0, fmt.Errorf("failed to test font \"%s\" %.1f: %w", font, size, err)
		}

		w, h := context.MeasureMultilineString(text, 1)
		// fmt.Printf("Is %.0fx%.0f (%.0f) larger than %dx%d?\n", w, h, size, width, height)
		if w > float64(width) || h > float64(height) {
			return size - 1, nil
		}
	}

	return 1000, nil
}
