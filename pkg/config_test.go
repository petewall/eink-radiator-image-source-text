package pkg_test

import (
	"encoding/json"
	"image"
	"os"

	"github.com/fogleman/gg"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-text/internal"
	"github.com/petewall/eink-radiator-image-source-text/internal/internalfakes"
	"github.com/petewall/eink-radiator-image-source-text/pkg"
)

var _ = Describe("Config", func() {
	var (
		backgroundImage *image.RGBA
		context         *internalfakes.FakeContext
		returnedImage   *image.RGBA
		makeBackground  *internalfakes.FakeBackgroundMaker
		newContext      *internalfakes.FakeContextMaker
	)

	BeforeEach(func() {
		backgroundImage = image.NewRGBA(image.Rect(0, 0, 300, 200))
		makeBackground = &internalfakes.FakeBackgroundMaker{}
		makeBackground.Returns(backgroundImage)
		internal.MakeBackground = makeBackground.Spy

		context = &internalfakes.FakeContext{}
		newContext = &internalfakes.FakeContextMaker{}
		newContext.Returns(context)
		internal.NewContext = newContext.Spy

		returnedImage = image.NewRGBA(image.Rect(0, 0, 300, 200))
		context.ImageReturns(returnedImage)
	})

	Describe("GenerateImage", func() {
		It("generates an image", func() {
			config := &pkg.Config{
				Text: "It is now safe to turn off your computer",
			}
			img, err := config.GenerateImage(300, 200)
			Expect(err).ToNot(HaveOccurred())

			By("setting the background", func() {
				Expect(makeBackground.CallCount()).To(Equal(1))
				width, height, color := makeBackground.ArgsForCall(0)
				Expect(width).To(Equal(300))
				Expect(height).To(Equal(200))
				Expect(color).To(Equal("white"))
			})

			By("writing the text", func() {
				Expect(newContext.CallCount()).To(Equal(1))
				Expect(newContext.ArgsForCall(0)).To(Equal(backgroundImage))
				Expect(context.SetColorCallCount()).To(Equal(1))
				Expect(context.SetColorArgsForCall(0)).To(Equal(colornames.Map["black"]))
				Expect(context.DrawStringWrappedCallCount()).To(Equal(1))
				text, x, y, ax, ay, width, lineSpacing, align := context.DrawStringWrappedArgsForCall(0)
				Expect(text).To(Equal("It is now safe to turn off your computer"))
				Expect(x).To(Equal(150.0))
				Expect(y).To(Equal(100.0))
				Expect(ax).To(Equal(0.5))
				Expect(ay).To(Equal(0.5))
				Expect(width).To(Equal(300.0))
				Expect(lineSpacing).To(Equal(1.0))
				Expect(align).To(Equal(gg.AlignCenter))
			})

			By("returning the image", func() {
				Expect(context.ImageCallCount()).To(Equal(1))
				Expect(img).To(Equal(returnedImage))
			})
		})
	})
})

var _ = Describe("ParseConfig", func() {
	var (
		configFile         *os.File
		configFileContents []byte
	)

	JustBeforeEach(func() {
		var err error
		configFile, err = os.CreateTemp("", "blank-config.yaml")
		Expect(err).ToNot(HaveOccurred())
		_, err = configFile.Write(configFileContents)
		Expect(err).ToNot(HaveOccurred())
	})

	BeforeEach(func() {
		config := pkg.Config{
			Text: "It is now safe to turn off your computer",
		}
		var err error
		configFileContents, err = yaml.Marshal(config)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.Remove(configFile.Name())).To(Succeed())
	})

	It("parses the image config file", func() {
		config, err := pkg.ParseConfig(configFile.Name())
		Expect(err).ToNot(HaveOccurred())
		Expect(config.Text).To(Equal("It is now safe to turn off your computer"))
	})

	Context("config file is json formatted", func() {
		BeforeEach(func() {
			config := pkg.Config{
				Text: "It is now safe to turn off your computer",
			}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("parses just fine", func() {
			config, err := pkg.ParseConfig(configFile.Name())
			Expect(err).ToNot(HaveOccurred())
			Expect(config.Text).To(Equal("It is now safe to turn off your computer"))
		})
	})

	When("reading the config file fails", func() {
		It("returns an error", func() {
			_, err := pkg.ParseConfig("this file does not exist")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to read image config file: open this file does not exist: no such file or directory"))
		})
	})

	When("parsing the config file fails", func() {
		BeforeEach(func() {
			configFileContents = []byte("this is invalid yaml!")
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("failed to parse image config file: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `this is...` into pkg.Config"))
		})
	})

	// Commented out because there's no invalid format at the moment
	// When("the config file has invalid data", func() {
	// 	BeforeEach(func() {
	// 		config := pkg.Config{
	// 			Text: "",
	// 		}
	// 		var err error
	// 		configFileContents, err = json.Marshal(config)
	// 		Expect(err).ToNot(HaveOccurred())
	// 	})

	// 	It("returns an error", func() {
	// 		_, err := pkg.ParseConfig(configFile.Name())
	// 		Expect(err).To(HaveOccurred())
	// 		Expect(err.Error()).To(Equal("config file is not valid: scale value is invalid: \"zelda\", must be one of resize, contain, cover"))
	// 	})
	// })
})
