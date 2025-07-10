package changetypes

import (
	"bitmap/models"
	"bitmap/filter"
)

func Filter(img *models.BMPImg, filters []string) {
	for _, f := range filters {
		switch f {
		case "red":
			filter.Red(img)
		case "green":
			filter.Green(img)
		case "blue":
			filter.Blue(img)
		case "grayscale":
			filter.Grayscale(img)
		case "negative":
			filter.Negative(img)
		case "pixelate":
			filter.Pixelate(img, 20)
		case "blur":
			filter.Blur(img, 5)
		}
	}
}
