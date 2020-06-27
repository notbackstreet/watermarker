package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var opacityBulk float64
var outputFolder string
var defaultOpacityBulk = 0.25

// bulkCmd represents the bulk command
var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "Bulk is used to watermark all images in a folder",
	Long: `The "bulk" command is used to watermark multiple images at once. It takes
	two arguments, first a folder containing images to be watermarked, and second
	an image to apply as the watermark. The level of opacity for the watermark image
	can be overwritten with the "--opactiy" or "-o" flag. By default, the watermarked
	images will be saved in a new folder named "watermarked" within the orginal folder,
	to override the "watermarked" folder name  use the "--name" or "-n" flag and
	provide a name for the new folder.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if opacityBulk > 1.0 || opacityBulk < 0.0 {
			fmt.Printf("Invalid opacity entered (opacity must be between 0.0 and 1.0), defaulting to %v ...\n", defaultOpacityBulk)
			opacityBulk = defaultOpacityBulk
		}
		if err := BulkWatermark(args[0], args[1], outputFolder, opacityBulk); err != nil {
			log.Fatalf("Unable to bulk watermark images in folder, received the following error: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(bulkCmd)
	bulkCmd.Flags().Float64VarP(&opacityBulk, "opacity", "o", 0.25, fmt.Sprintf("Opacity can be used to overwrite the default opacity (%v) of the image being used as the watermark. It expects a float between 0.0 and 1.0.", defaultOpacityBulk))
	bulkCmd.Flags().StringVarP(&outputFolder, "name", "n", "watermarked/", "Name can be used to customize the name of the generated folder containing the new watermarked images.")
}
