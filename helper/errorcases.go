package helper

import "errors"

var (
	ErrInvalidArgs = errors.New("provided arguments are invalid; verify your input")
	ErrNoFilename  = errors.New("filename for BMP is missing; please provide a valid file")
	ErrWrongBMP    = errors.New("wrong BMP file")
)
