package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const ImageTypeName = "text"

var rootCmd = &cobra.Command{
	Use:   ImageTypeName,
	Short: "Generate an image from text",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
