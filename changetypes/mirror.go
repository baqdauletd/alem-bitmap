package changetypes

import "bitmap/models"

func Mirror(img *models.BMPImg, direction string) {
	h := len(img.PixelData)
	w := len(img.PixelData[0])

	switch direction {
	case "horizontal", "h", "hor", "horizontally":
		for i := 0; i < h; i++ {
			for j := 0; j < w/2; j++ {
				img.PixelData[i][j], img.PixelData[i][w-j-1] = img.PixelData[i][w-j-1], img.PixelData[i][j]
			}
		}
	case "vertical", "v", "ver", "vertically":
		for i := 0; i < h/2; i++ {
			for j := 0; j < w; j++ {
				img.PixelData[i][j], img.PixelData[h-i-1][j] = img.PixelData[h-i-1][j], img.PixelData[i][j]
			}
		}
	}
}
