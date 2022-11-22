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

	DescribeTable("generated images",
		func(configPath, expectedPath string) {
			outputFile = "text.png"
			Run("generate --config " + configPath + " --output " + outputFile + " --height 300 --width 400")
			Eventually(CommandSession).WithTimeout(time.Second * 5).Should(Exit(0))

			actualData, err := os.ReadFile(outputFile)
			Expect(err).ToNot(HaveOccurred())
			expectedData, err := os.ReadFile(expectedPath)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualData).To(Equal(expectedData))
		},
		Entry("text only", "inputs/text.yaml", "outputs/text.png"),
		Entry("text with newlines", "inputs/text_withnewlines.yaml", "outputs/text_withnewlines.png"),
		Entry("text with word wrapping enabled", "inputs/text_wrapped.yaml", "outputs/text_wrapped.png"),
		Entry("text with newlines and word wrapping enabled", "inputs/text_wrapped_withnewlines.yaml", "outputs/text_wrapped_withnewlines.png"),

		Entry("text only, sized to fit", "inputs/text_fit.yaml", "outputs/text_fit.png"),
		Entry("text with newlines, sized to fit", "inputs/text_fit_withnewlines.yaml", "outputs/text_fit_withnewlines.png"),
		Entry("text with word wrapping enabled, sized to fit", "inputs/text_fit_wrapped.yaml", "outputs/text_fit_wrapped.png"),
		Entry("text with newlines and word wrapping enabled, sized to fit", "inputs/text_fit_wrapped_withnewlines.yaml", "outputs/text_fit_wrapped_withnewlines.png"),
	)

	When("using --to-stdout", func() {
		It("writes the image to stdout", func() {
			Run("generate --config inputs/text.yaml --to-stdout --height 300 --width 400")
			Eventually(CommandSession).WithTimeout(time.Second * 5).Should(Exit(0))

			expectedData, err := os.ReadFile("outputs/text.png")
			Expect(err).ToNot(HaveOccurred())
			Expect(CommandSession.Out.Contents()).To(Equal(expectedData))
		})
	})

	When("using a color that doesn't exist", func() {
		It("returns an error", func() {
			Run("generate --config inputs/invalid_color.json --height 200 --width 200")
			Eventually(CommandSession).WithTimeout(time.Second * 5).Should(Exit(1))
		})
	})
})
