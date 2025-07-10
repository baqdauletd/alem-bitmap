package models

const (
	BMPfileType = 0x4D42
)

type FileHeader struct {
	FileType   uint16
	FileSize   uint32
	Reserve1   uint16
	Reserve2   uint16
	OffsetBits uint32
}

type InfoHeader struct {
	Size            uint32
	Width           int32
	Height          int32
	Planes          uint16
	BitsPerPixel    uint16
	Compression     uint32
	ImageSize       uint32
	XPixelsPerM     int32
	YPixelsPerM     int32
	ColorsUsed      uint32
	ColorsImportant uint32
}

type BMPImg struct {
	FileHeader FileHeader
	InfoHeader InfoHeader
	PixelData  [][]Pixel
}

type Pixel struct {
	Blue  byte
	Green byte
	Red   byte
}

type ChangeType int

const (
	ChangeTypeMirror ChangeType = iota
	ChangeTypeRotate
	ChangeTypeFilter
	ChangeTypeCrop
)

type Change struct {
	Type ChangeType
	Args []string
}

type MirrorOpts struct {
	Horizontal bool
	Vertical   bool
}

type CropOpts struct {
	OffsetX int
	OffsetY int
	Width   int
	Height  int
}
