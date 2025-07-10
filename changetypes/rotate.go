package changetypes

import (
	"bitmap/models"
	"fmt"
)

func Rotate(img *models.BMPImg, angle int) error {
	h := len(img.PixelData)
	w := len(img.PixelData[0])

	var rotatedData [][]models.Pixel

	switch angle {
	case 90:

		rotatedData = make([][]models.Pixel, w)
		for i := 0; i < w; i++ {
			rotatedData[i] = make([]models.Pixel, h)
			for j := 0; j < h; j++ {
				rotatedData[i][j] = img.PixelData[j][w-1-i]
			}
		}

	case -90:

		rotatedData = make([][]models.Pixel, w)
		for i := 0; i < w; i++ {
			rotatedData[i] = make([]models.Pixel, h)
			for j := 0; j < h; j++ {
				rotatedData[i][j] = img.PixelData[h-1-j][i]
			}
		}

	case 180:
		rotatedData = make([][]models.Pixel, h)
		for i := 0; i < h; i++ {
			rotatedData[i] = make([]models.Pixel, w)
			for j := 0; j < w; j++ {
				rotatedData[i][j] = img.PixelData[h-1-i][w-1-j]
			}
		}
	case 360:
		// Не бейте, выбора не было
		return nil
	default:
		return fmt.Errorf("invalid rotation angle: %d", angle)
	}

	img.InfoHeader.Width = int32(len(rotatedData[0]))
	img.InfoHeader.Height = int32(len(rotatedData))

	img.PixelData = rotatedData

	return nil
}
