package image

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"
)

func ImageExt(filename, mime string) (ext string, err error) {
	switch mime {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	default:
		ext = filepath.Ext(filename)
		switch ext {
		case ".jpg", ".png", ".gif":
		default:
			err = fmt.Errorf("unsupport image format `%s`", filename)
		}
	}
	return
}

func Decord(r io.Reader, ext string) (img image.Image, err error) {
	switch ext {
	case ".jpg":
		img, err = jpeg.Decode(r)
	case ".png":
		img, err = png.Decode(r)
	case ".gif":
		img, err = gif.Decode(r)
	default:
		err = fmt.Errorf("unsupport docord format `%s`", ext)
	}
	return
}
