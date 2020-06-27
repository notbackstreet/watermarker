package cmd

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	"github.com/disintegration/imaging"
)

// FileImage represents an image file
type FileImage struct {
	Name  string
	Image image.Image
}

// WatermarkImage overlays `watermarkFile` on top of `originalFile` with the specified `opacity` and saves
// the resulting image to `outputFile`.
func WatermarkImage(originalFile, watermarkFile, outputFile string, opacity float64) error {
	// Open original image.
	src, err := imaging.Open(originalFile)
	if err != nil {
		return errors.Wrap(err, "failed to open original image")
	}

	// Open the image to use as the watermark
	watermark, err := imaging.Open(watermarkFile)
	if err != nil {
		return errors.Wrap(err, "failed to open watermark image")
	}
	// Overlay the watermark image on the original image with the specified opacity
	marked := imaging.OverlayCenter(src, watermark, opacity)

	// Save the watermarked image
	err = imaging.Save(marked, outputFile)
	if err != nil {
		return errors.Wrap(err, "failed to save new watermarked image")
	}
	return nil
}

func BulkWatermark(srcFolder, watermarkFile, outputFolder string, opacity float64) error {
	// Open the image to use as the watermark
	watermark, err := imaging.Open(watermarkFile)
	if err != nil {
		return errors.Wrap(err, "failed to open watermark image")
	}

	// Get the image files from the srcFolder
	files, err := ioutil.ReadDir(srcFolder)
	if err != nil {
		return errors.Wrap(err, "failed to get files from source folder")
	}

	var images []FileImage
	for _, file := range files {
		image, err := imaging.Open(srcFolder + file.Name())
		if err != nil {
			fmt.Printf("Skipping file %s ...\n", file.Name())
			continue
		} else {
			fmt.Printf("Found image file %s\n", file.Name())
			images = append(images, FileImage{Name: file.Name(), Image: image})
		}
	}

	if len(images) == 0 {
		fmt.Println("No images found in source folder. Did you forget to add a trailing '/' to the directory name?")
		return nil
	}

	// Convert the images
	fmt.Println("Watermarking images...")
	var watermarked []FileImage
	for _, img := range images {
		watermarked = append(watermarked, FileImage{Name: img.Name, Image: imaging.OverlayCenter(img.Image, watermark, opacity)})
	}

	// Save the images
	fmt.Println("Saving images...")
	os.Mkdir(outputFolder, 0700)
	for _, marked := range watermarked {
		err = imaging.Save(marked.Image, outputFolder+marked.Name)
		if err != nil {
			fmt.Printf("Failed to save watermarked image %s, received error %s\n", marked.Name, err.Error())
			continue
		}
	}

	return nil
}
