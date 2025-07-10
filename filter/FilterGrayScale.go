package filter

import "bitmap/models"

func Grayscale(img *models.BMPImg) {
	height := len(img.PixelData)
	width := len(img.PixelData[0])

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			pixel := &img.PixelData[row][col]

			gray := byte(float64(pixel.Red)*0.3 + float64(pixel.Green)*0.59 + float64(pixel.Blue)*0.11)
			pixel.Green = gray
			pixel.Red = gray
			pixel.Blue = gray
		}
	}
}
