package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var opacity float64
var outFileName string
var defaultOpacity = 0.25

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "image is used to watermark a single image",
	Long: `The "image" command is used to watermark a single image. It takes
	two image files as its arguments, and will overlay the second image over
	the first image with the default opacity. The level of opacity can be specified with the
	"--opacity" or "-o" flag. By default, the first image file will be overwritten with
	the new "watermarked" image, to output to a new file use the "--name" or "-n" flag and
	provide a name for the new file.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if outFileName == "" {
			outFileName = args[0]
		}
		if opacity > 1.0 || opacity < 0.0 {
			fmt.Printf("Invalid opacity entered (opacity must be between 0.0 and 1.0), defaulting to %v ...\n", defaultOpacity)
			opacity = defaultOpacity
		}
		if err := WatermarkImage(args[0], args[1], outFileName, opacity); err != nil {
			log.Fatalf("Unable to watermark image, received the following error: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.Flags().Float64VarP(&opacity, "opacity", "o", defaultOpacity, fmt.Sprintf("Opacity can be used to overwrite the default opacity (%v) of the image being used as the watermark. It expects a float between 0.0 and 1.0.", defaultOpacity))
	imageCmd.Flags().StringVarP(&outFileName, "name", "n", "", "Name can be used to output the watermarked image to a new file with the provided name.")
}
