package cmd

import (
	"github.com/pkg/errors"

	"github.com/disintegration/imaging"
)

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
