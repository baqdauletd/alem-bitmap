package main

import (
	"fmt"
	"os"
	"strings"

	"bitmap/helper"
)

var fileName string

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(helper.ErrInvalidArgs)
		helper.PrintHelp()
		os.Exit(1)
	}
	command := args[1]
	if len(os.Args) == 2 {
		for i := range args {
			if strings.HasPrefix(args[i], "--help") || strings.HasPrefix(args[i], "--h") {
				helper.PrintHelp()
			}
		}
	}
	if command == "header" && (len(args) == 3 && (args[2] == "--help" || args[2] == "-h")) {
		helper.PrintHelp("header")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println(helper.ErrNoFilename)
		os.Exit(1)
	}

	for i := range args {
		if strings.HasSuffix(args[i], ".bmp") {
			fileName = args[i]
		}
	}

	switch command {
	case "header":
		if len(os.Args) < 3 || !strings.HasSuffix(args[2], ".bmp") {
			fmt.Println("error: file '' is not a bmp file (.bmp)")
			os.Exit(1)
		}
		helper.PrintBMP(fileName)

	case "apply":
		if len(os.Args) < 4 {
			if command == "apply" && (len(args) == 3 && (args[2] == "--help" || args[2] == "--h")) {
				helper.PrintHelp("apply")
				os.Exit(1)
			} else {

				fmt.Println("error\n", helper.ErrInvalidArgs)
				os.Exit(1)
			}
		}
		var inputFile string
		var outputFile string
		if strings.HasSuffix(args[len(args)-1], ".bmp") && strings.HasSuffix(args[len(args)-2], ".bmp") {
			inputFile = args[len(args)-2]
			outputFile = args[len(args)-1]
		} else {
			fmt.Println("error: file '' is not a bmp file (.bmp)")
			os.Exit(1)
		}
		changes, err := helper.ParseChanges(args[2 : len(args)-2])
		if err != nil {
			fmt.Println("error\n", err)
			os.Exit(1)
		}

		img, err := helper.ReadBMP(inputFile)
		if err != nil {
			fmt.Println("error\n", err)
			os.Exit(1)
		}

		err = helper.ApplyChanges(img, changes)
		if err != nil {
			fmt.Println("error\n", err)
			os.Exit(1)
		}

		err = helper.WriteBMP(outputFile, img)
		if err != nil {
			fmt.Println("error\n", err)
			os.Exit(1)
		}

		fmt.Println("success")
	default:
		helper.PrintHelp()
	}
}
