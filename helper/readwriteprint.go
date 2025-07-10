package helper

import (
	"encoding/binary"
	"fmt"
	"os"

	"bitmap/models"
)

func ReadBMP(name string) (*models.BMPImg, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}
	defer file.Close()

	var h models.FileHeader
	err = binary.Read(file, binary.LittleEndian, &h)
	if err != nil {
		return nil, fmt.Errorf("error reading file header")
	}

	if h.FileType != models.BMPfileType {
		return nil, fmt.Errorf("wrong BMP file type")
	}

	var i models.InfoHeader
	err = binary.Read(file, binary.LittleEndian, &i)
	if err != nil {
		return nil, fmt.Errorf("error reading info header")
	}

	rowSize := (int(i.Width)*3 + 3) & ^3
	pixels := make([][]models.Pixel, i.Height)
	for y := 0; y < abs(int(i.Height)); y++ {
		row := make([]byte, rowSize)
		if _, err := file.Read(row); err != nil {
			return nil, fmt.Errorf("error reading row %d: %v", y, err)
		}
		pixels[y] = make([]models.Pixel, i.Width)
		for x := 0; x < int(i.Width); x++ {
			offset := x * 3
			pixels[y][x] = models.Pixel{
				Blue:  row[offset],
				Green: row[offset+1],
				Red:   row[offset+2],
			}
		}
	}

	return &models.BMPImg{
		FileHeader: h,
		InfoHeader: i,
		PixelData:  pixels,
	}, nil
}

func WriteBMP(name string, img *models.BMPImg) error {
	if img == nil || len(img.PixelData) == 0 {
		return fmt.Errorf("invalid image data")
	}
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("can't create file: %v", err)
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, img.FileHeader)
	if err != nil {
		return fmt.Errorf("can't write file header: %v", err)
	}

	err = binary.Write(file, binary.LittleEndian, img.InfoHeader)
	if err != nil {
		return fmt.Errorf("can't write info header: %v", err)
	}

	rowSize := (int(img.InfoHeader.Width)*3 + 3) & ^3
	for y := 0; y < len(img.PixelData); y++ {
		rowBytes := make([]byte, rowSize)
		for x := 0; x < len(img.PixelData[y]); x++ {
			offset := x * 3
			rowBytes[offset] = img.PixelData[y][x].Blue
			rowBytes[offset+1] = img.PixelData[y][x].Green
			rowBytes[offset+2] = img.PixelData[y][x].Red
		}

		if err := binary.Write(file, binary.LittleEndian, rowBytes); err != nil {
			return fmt.Errorf("can't write row %d: %v", y, err)
		}
	}

	return nil
}

func PrintBMP(name string) {
	img, err := ReadBMP(name)
	if err != nil {
		fmt.Println("error\n", err)
		os.Exit(1)
	}

	h := img.FileHeader
	i := img.InfoHeader

	fmt.Println("BMP Header:")
	fmt.Printf("- FileType %c%c\n", byte(h.FileType), byte(h.FileType>>8))
	fmt.Printf("- FileSizeInBytes %d\n", h.FileSize)
	fmt.Printf("- HeaderSize %d\n", 54)

	fmt.Println("DIB Header:")
	fmt.Printf("- DibHeaderSize %d\n", i.Size)
	fmt.Printf("- WidthInPixels %d\n", i.Width)
	fmt.Printf("- HeightInPixels %d\n", i.Height)
	fmt.Printf("- PixelSizeInBits %d\n", i.BitsPerPixel)
	fmt.Printf("- ImageSizeInBytes %d\n", i.ImageSize)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
