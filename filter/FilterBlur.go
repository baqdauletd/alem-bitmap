package filter

import "bitmap/models"

func Blur(img *models.BMPImg, blurRadius int) {
	height := len(img.PixelData)
	width := len(img.PixelData[0])

	copyData := make([][]models.Pixel, height)
	for i := range copyData {
		copyData[i] = make([]models.Pixel, width)
		copy(copyData[i], img.PixelData[i])
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			var sumRed, sumGreen, sumBlue, count int

			for dy := -blurRadius; dy <= blurRadius; dy++ {
				for dx := -blurRadius; dx <= blurRadius; dx++ {

					neighborRow := row + dy
					neighborCol := col + dx

					if neighborRow >= 0 && neighborRow < height && neighborCol >= 0 && neighborCol < width {
						neighborPixel := copyData[neighborRow][neighborCol]
						sumRed += int(neighborPixel.Red)
						sumGreen += int(neighborPixel.Green)
						sumBlue += int(neighborPixel.Blue)
						count++
					}
				}
			}

			img.PixelData[row][col].Red = uint8(sumRed / count)
			img.PixelData[row][col].Green = uint8(sumGreen / count)
			img.PixelData[row][col].Blue = uint8(sumBlue / count)
		}
	}
}
