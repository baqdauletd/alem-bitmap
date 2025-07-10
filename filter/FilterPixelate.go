package filter

import (
	"bitmap/models"
)

func Pixelate(img *models.BMPImg, blocksize int) {
	h := len(img.PixelData)
	w := len(img.PixelData[0])

	for y := 0; y < h; y += blocksize {
		for x := 0; x < w; x += blocksize {
			colorPixel := avgColor(img, x, y, blocksize)
			fillBlock(img, x, y, blocksize, colorPixel)
		}
	}
}

func avgColor(img *models.BMPImg, startX, startY, blocksize int) models.Pixel {
	var redSum, greenSum, blueSum, current uint32
	h := len(img.PixelData)
	w := len(img.PixelData[0])

	for y := startY; y < startY+blocksize && y < h; y++ {
		for x := startX; x < startX+blocksize && x < w; x++ {
			redSum += uint32(img.PixelData[y][x].Red)
			greenSum += uint32(img.PixelData[y][x].Green)
			blueSum += uint32(img.PixelData[y][x].Blue)
			current++
		}
	}

	return models.Pixel{
		Red:   byte(redSum / current),
		Green: byte(greenSum / current),
		Blue:  byte(blueSum / current),
	}
}

func fillBlock(img *models.BMPImg, startX, startY, blocksize int, colorPixel models.Pixel) {
	h := len(img.PixelData)
	w := len(img.PixelData[0])

	for y := startY; y < startY+blocksize && y < h; y++ {
		for x := startX; x < startX+blocksize && x < w; x++ {
			img.PixelData[y][x] = colorPixel
		}
	}
}
