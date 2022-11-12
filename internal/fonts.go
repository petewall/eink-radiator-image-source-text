package internal

import (
	"github.com/flopp/go-findfont"
)

//counterfeiter:generate . FontFinder
type FontFinder func(font string) (filePath string, err error)

var FindFont FontFinder = findfont.Find
