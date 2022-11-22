package internal_test

import (
	"image"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/petewall/eink-radiator-image-source-text/internal"
)

var _ = Describe("NewContext", func() {
	It("returns a new image context", func() {
		background := image.NewRGBA(image.Rect(0, 0, 300, 200))
		context := internal.NewContext(background)
		Expect(context).ToNot(BeNil())
	})
})
