package test_test

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Generate", func() {
	var outputFile string

	AfterEach(func() {
		if outputFile != "" {
			Expect(os.Remove(outputFile)).To(Succeed())
			outputFile = ""
		}
	})

	It("generates an image that is resized to fit the desired resolution", func() {
		outputFile = "text.png"
		Run("generate --config text_only_config.json --output text.png --height 300 --width 400")
		Eventually(CommandSession).WithTimeout(time.Second * 5).Should(Exit(0))

		By("saving the image to a file", func() {
			actualData, err := os.ReadFile(outputFile)
			Expect(err).ToNot(HaveOccurred())
			expectedData, err := os.ReadFile("text_only_expected.png")
			Expect(err).ToNot(HaveOccurred())
			Expect(actualData).To(Equal(expectedData))
		})
	})

	When("using --to-stdout", func() {
		It("writes the image to stdout", func() {
			Run("generate --config text_only_config.json --to-stdout --height 300 --width 400")
			Eventually(CommandSession).WithTimeout(time.Second * 5).Should(Exit(0))

			expectedData, err := os.ReadFile("text_only_expected.png")
			Expect(err).ToNot(HaveOccurred())
			Expect(CommandSession.Out.Contents()).To(Equal(expectedData))
		})
	})

	Context("providing colors", func() {
		It("generates an image using those colors", func() {
			Run("generate --config colors_config.yaml --to-stdout --height 200 --width 200")
			Eventually(CommandSession).WithTimeout(time.Second * 5).Should(Exit(0))

			expectedData, err := os.ReadFile("colors_expected.png")
			Expect(err).ToNot(HaveOccurred())
			Expect(CommandSession.Out.Contents()).To(Equal(expectedData))
		})

		When("using a color that doesn't exist", func() {
			It("returns an error", func() {
				Run("generate --config invalid_color_config.yaml --height 200 --width 200")
				Eventually(CommandSession).WithTimeout(time.Second * 5).Should(Exit(1))
			})
		})
	})
})
