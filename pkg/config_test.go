package pkg_test

import (
	"encoding/json"
	"errors"
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
		findFont        *internalfakes.FakeFontFinder
		textFitter      *internalfakes.FakeTextFitter
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

		findFont = &internalfakes.FakeFontFinder{}
		findFont.Returns("/path/to/your/font.ttf", nil)
		internal.FindFont = findFont.Spy

		textFitter = &internalfakes.FakeTextFitter{}
		textFitter.Returns(64, nil)
		internal.FitText = textFitter.Spy

		returnedImage = image.NewRGBA(image.Rect(0, 0, 300, 200))
		context.ImageReturns(returnedImage)
	})

	Describe("GenerateImage", func() {
		It("generates an image", func() {
			config := &pkg.Config{
				Text:  "It is now safe to turn off your computer",
				Color: "black",
				Background: pkg.BackgroundType{
					Color: "white",
				},
				Font: "comic sans",
				Size: 128,
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

				Expect(findFont.CallCount()).To(Equal(1))
				Expect(findFont.ArgsForCall(0)).To(Equal("comic sans"))

				Expect(context.LoadFontFaceCallCount()).To(Equal(1))
				fontPath, fontSize := context.LoadFontFaceArgsForCall(0)
				Expect(fontPath).To(Equal("/path/to/your/font.ttf"))
				Expect(fontSize).To(Equal(128.0))

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

		Context("font size is set to 0", func() {
			It("generates an image with the text fit to the size", func() {
				config := &pkg.Config{
					Text: "It is now safe to turn off your computer",
					Font: "charcoal",
				}
				_, err := config.GenerateImage(300, 200)
				Expect(err).ToNot(HaveOccurred())

				By("calling the text fitting code to get the right font size", func() {
					Expect(textFitter.CallCount()).To(Equal(1))
					ctx, text, font, width, height := textFitter.ArgsForCall(0)
					Expect(ctx).To(Equal(context))
					Expect(text).To(Equal("It is now safe to turn off your computer"))
					Expect(font).To(Equal("/path/to/your/font.ttf"))
					Expect(width).To(Equal(300))
					Expect(height).To(Equal(200))
				})
			})

			When("fitting the text fails", func() {
				BeforeEach(func() {
					textFitter.Returns(0, errors.New("fit text failed"))
				})
				It("returns an error", func() {
					config := &pkg.Config{
						Text: "It is now safe to turn off your computer",
						Font: "charcoal",
					}
					_, err := config.GenerateImage(300, 200)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("could not fit font \"charcoal\": fit text failed"))
				})
			})
		})

		When("finding the font fails", func() {
			BeforeEach(func() {
				findFont.Returns("", errors.New("find font failed"))
			})
			It("returns an error", func() {
				config := &pkg.Config{
					Text:  "It is now safe to turn off your computer",
					Color: "black",
					Background: pkg.BackgroundType{
						Color: "white",
					},
					Font: "comic sans",
					Size: 128,
				}
				_, err := config.GenerateImage(300, 200)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("could not find font \"comic sans\": find font failed"))
			})
		})

		When("setting the font fails", func() {
			BeforeEach(func() {
				context.LoadFontFaceReturns(errors.New("load font face failed"))
			})
			It("returns an error", func() {
				config := &pkg.Config{
					Text:  "It is now safe to turn off your computer",
					Color: "black",
					Background: pkg.BackgroundType{
						Color: "white",
					},
					Font: "comic sans",
					Size: 128,
				}
				_, err := config.GenerateImage(300, 200)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to set font \"comic sans\" 128.0: load font face failed"))
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

	When("the config file has missing data", func() {
		BeforeEach(func() {
			config := pkg.Config{
				Text: "",
			}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("config file is not valid: missing text"))
		})
	})

	When("the config file has invalid color", func() {
		BeforeEach(func() {
			config := pkg.Config{
				Text:  "what",
				Color: "the golf",
			}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("config file is not valid: unknown color: \"the golf\""))
		})
	})

	When("the config file has invalid background color", func() {
		BeforeEach(func() {
			config := pkg.Config{
				Text:  "shoes",
				Color: "blue",
				Background: pkg.BackgroundType{
					Color: "suede",
				},
			}
			var err error
			configFileContents, err = json.Marshal(config)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error", func() {
			_, err := pkg.ParseConfig(configFile.Name())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("config file is not valid: invalid background: unknown color: \"suede\""))
		})
	})
})
