package pkg

import (
	"fmt"
	"image"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-text/internal"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . ImageGenerator
type ImageGenerator interface {
	GenerateImage(width, height int) (image.Image, error)
}

// type BackgroundType struct {
// 	Color string `json:"color" yaml:"color"`
// }

type Config struct {
	Text string `json:"text" yaml:"text"`
}

func (c *Config) GenerateImage(width, height int) (image.Image, error) {
	background := internal.MakeBackground(width, height, "white")
	context := internal.NewContext(background)
	context.SetColor(colornames.Map["black"])
	context.DrawStringWrapped(c.Text, float64(width/2), float64(height/2), 0.5, 0.5, float64(width), 1.0, gg.AlignCenter)
	return context.Image(), nil
}

func (c *Config) Validate() error {
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

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("config file is not valid: %w", err)
	}

	return config, nil
}
