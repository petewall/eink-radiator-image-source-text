package pkg

import (
	"fmt"
	"image"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v2"

	blank "github.com/petewall/eink-radiator-image-source-blank/pkg"
	"github.com/petewall/eink-radiator-image-source-text/internal"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . ImageGenerator
type ImageGenerator interface {
	GenerateImage(width, height int) (image.Image, error)
}

type BackgroundType struct {
	Color string `json:"color" yaml:"color"`
}

type Config struct {
	Text       string         `json:"text" yaml:"text"`
	Color      string         `json:"color" yaml:"color"`
	Background BackgroundType `json:"background" yaml:"background"`
	Font       string         `json:"font,omitempty" yaml:"font,omitempty"`
	Size       float64        `json:"size,omitempty" yaml:"size,omitempty"`
}

func (c *Config) GenerateImage(width, height int) (image.Image, error) {
	background := internal.MakeBackground(width, height, c.Background.Color)
	context := internal.NewContext(background)
	context.SetColor(colornames.Map[c.Color])

	font, err := internal.FindFont(c.Font)
	if err != nil {
		return nil, fmt.Errorf("could not find font \"%s\": %w", c.Font, err)
	}

	err = context.LoadFontFace(font, c.Size)
	if err != nil {
		return nil, fmt.Errorf("failed to set font \"%s\" %.1f: %w", c.Font, c.Size, err)
	}
	context.DrawStringWrapped(c.Text, float64(width/2), float64(height/2), 0.5, 0.5, float64(width), 1.0, gg.AlignCenter)
	return context.Image(), nil
}

func (c *Config) Validate() error {
	if c.Text == "" {
		return fmt.Errorf("missing text")
	}

	if _, isPresent := colornames.Map[c.Color]; !isPresent {
		return fmt.Errorf("unknown color: \"%s\"", c.Color)
	}

	backgroundConfig := blank.Config{Color: c.Background.Color}
	if err := backgroundConfig.Validate(); err != nil {
		return fmt.Errorf("invalid background: %w", err)
	}

	return nil
}

func ParseConfig(path string) (*Config, error) {
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read image config file: %w", err)
	}

	var config *Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse image config file: %w", err)
	}

	if config.Color == "" {
		config.Color = "black"
	}

	if config.Background.Color == "" {
		config.Background.Color = "white"
	}

	if config.Font == "" {
		config.Font = "Arial"
	}

	if config.Size == 0 {
		config.Size = 48
	}

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("config file is not valid: %w", err)
	}

	return config, nil
}
