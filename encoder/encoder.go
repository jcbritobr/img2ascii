package encoder

import (
	"bytes"
	"image"
	"image/color"

	// Needs decode jpeg and png files
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/nfnt/resize"
)

var (
	levels = []string{" ", "░", "▒", "▓", "█"}
)

// Encode encods png and jpeg files into ascii representation
func Encode(reader io.Reader, scale uint) (string, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}

	var data string
	buffer := bytes.NewBufferString(data)
	n := resize.Resize(scale, 0, img, resize.Lanczos3)

	for y := n.Bounds().Min.Y; y < n.Bounds().Max.Y; y++ {
		for x := n.Bounds().Min.X; x < n.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(n.At(x, y)).(color.Gray)
			level := c.Y / uint8((255 / len(levels)))
			if level >= uint8(len(levels)) {
				level = uint8(len(levels)) - 1
			}
			buffer.WriteString(levels[level])
		}
		buffer.WriteByte('\n')
	}
	return buffer.String(), nil
}
