package changetypes

import (
	"bitmap/models"
	"fmt"
)

func Crop(image *models.BMPImg, opts models.CropOpts) error {
	originalWidth := int(image.InfoHeader.Width)
	originalHeight := abs(int(image.InfoHeader.Height))

	if opts.OffsetX >= originalWidth || opts.OffsetY >= originalHeight {
		return fmt.Errorf("crop area exceeds image boundaries")
	}

	if opts.Width == 0 {
		opts.Width = originalWidth - opts.OffsetX
	}
	if opts.Height == 0 {
		opts.Height = originalHeight - opts.OffsetY
	}

	if opts.OffsetX+opts.Width > originalWidth || opts.OffsetY+opts.Height > originalHeight {
		return fmt.Errorf("crop area exceeds image boundaries")
	}

	croppedData := make([][]models.Pixel, opts.Height)
	for i := range croppedData {
		croppedData[i] = make([]models.Pixel, opts.Width)
	}

	for i := 0; i < opts.Height; i++ {
		for j := 0; j < opts.Width; j++ {
			croppedData[i][j] = image.PixelData[originalHeight-opts.OffsetY-1-i][opts.OffsetX+j]
		}
	}

	for i := 0; i < opts.Height/2; i++ {
		croppedData[i], croppedData[opts.Height-1-i] = croppedData[opts.Height-1-i], croppedData[i]
	}

	image.PixelData = croppedData
	image.InfoHeader.Width = int32(opts.Width)
	image.InfoHeader.Height = int32(opts.Height)

	bytesPerPixel := 3
	rowSize := (opts.Width*bytesPerPixel + 3) & ^3
	imageSize := rowSize * opts.Height
	image.FileHeader.FileSize = uint32(imageSize + 54)

	image.InfoHeader.ImageSize = uint32(imageSize)

	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
