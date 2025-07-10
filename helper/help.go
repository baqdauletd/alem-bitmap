package helper

import (
	"fmt"
)

func PrintHelp(arg ...string) {
	if len(arg) < 1 {

		fmt.Println(helpMain)
		return
	}

	switch arg[0] {
	case "header":

		fmt.Println(helpHeader)
	case "apply":

		fmt.Println(helpApply)
	default:

		fmt.Println(helpMain)
	}
}

const (
	helpMain = `Usage:
    bitmap <command> [arguments]
  
The commands are:
    header    prints bitmap file header information
    apply     applies processing to the image and saves it to the file

Use "bitmap <command> --help" for more information about a command.
`

	helpHeader = `Usage:
    bitmap header <source_file>
  
Description:
    Prints bitmap file header information
  
Arguments:
    <source_file>    Path to the source bitmap (.bmp) file
`

	helpApply = `Usage:
    bitmap apply [options] <source_file> <output_file>

Description:
    Applies processing to the image and saves it to the file

Arguments:
    <source_file>    Path to the source bitmap file
    <output_file>    Path to save the processed bitmap file

Options:
    --mirror=<value>        Mirror the image. Values: horizontal, h, horizontally, hor, vertical, v, vertically, ver
    --filter=<value>        Apply a filter. Can be used multiple times. Values: blue, red, green, grayscale, negative, pixelate, blur
                            Default values: pixelate = 20px, blur = 20px
    --rotate=<value>        Rotate the image. Can be used multiple times. Values: right, 90, 180, 270, left, -90, -180, -270
    --crop=<value>          Crop the image. Format: OffsetX-OffsetY-Width-Height. Width and Height are optional

Examples:
    bitmap apply --mirror=horizontal --filter=grayscale input.bmp output.bmp
    bitmap apply --rotate=right --rotate=right --crop=20-20-100-100 input.bmp output.bmp
`
)
