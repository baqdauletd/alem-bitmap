package filter

import "bitmap/models"

func Blue(img *models.BMPImg) {
	height := len(img.PixelData)
	width := len(img.PixelData[0])

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			pixel := &img.PixelData[row][col]
			pixel.Green = 0
			pixel.Red = 0
		}
	}
}
