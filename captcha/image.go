package captcha

import (
	"image"
	"io"
)

type Image interface {
	image.PalettedImage
	// WriteTo writes captcha image in PNG format into the given writer.
	WriteTo(w io.Writer) (int64, error)
}
