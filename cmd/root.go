package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "watermarker",
	Short: "A simple tool to watermark images with another image.",
	Long: `Watermarker is a CLI tool written in Go that allows you to "watermark" images.
It does so by overlaying an image as a "watermark" over another image.
You can specify the opacity of the overlay image to have a stronger or softer "watermark".`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
