package filter

import "bitmap/models"

func Negative(img *models.BMPImg) {
	height := len(img.PixelData)
	width := len(img.PixelData[0])

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			pixel := &img.PixelData[row][col]
			pixel.Green = 255 - pixel.Green
			pixel.Red = 255 - pixel.Red
			pixel.Blue = 255 - pixel.Blue
		}
	}
}
