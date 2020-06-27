package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bulkCmd represents the bulk command
var bulkCmd = &cobra.Command{
	Use:   "bulk",
	Short: "Bulk is used to watermark all images in a folder",
	Long: `The "bulk" command is used to watermark multiple images at once. It takes
	two arguments, first a folder containing images to be watermarked, and second
	an image to apply as the watermark. The level of opacity for the watermark image
	can be overwritten with the "--opactiy" or "-op" flag. By default, the watermarked
	images will be saved in a new folder named "watermarked" within the orginal folder,
	to override the "watermarked" folder name  use the "--name" or "-n" flag and
	provide a name for the new folder.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bulk called")
	},
}

func init() {
	rootCmd.AddCommand(bulkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	bulkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bulkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
