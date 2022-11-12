package internal_test

import (
	"github.com/fogleman/gg"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/petewall/eink-radiator-image-source-text/internal"
)

var _ = Describe("FitText", func() {
	var font string

	BeforeEach(func() {
		var err error
		font, err = internal.FindFont("Ubuntu")
		Expect(err).ToNot(HaveOccurred())
	})

	It("finds the largest font size that will fit the area", func() {
		context := gg.NewContext(200, 200)
		size, err := internal.FitText(context, "small", font, 200, 200)
		Expect(err).ToNot(HaveOccurred())
		Expect(size).To(Equal(79.0))

		size, err = internal.FitText(context, "this is a longer amount of text", font, 200, 200)
		Expect(err).ToNot(HaveOccurred())
		Expect(size).To(Equal(13.0))

		size, err = internal.FitText(context, "this is also\nlonger, but\nsplit over\na few lines", font, 200, 200)
		Expect(err).ToNot(HaveOccurred())
		Expect(size).To(Equal(38.0))
	})
})
