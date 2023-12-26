package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -1.5, -1.5, +1.5, +1.5
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}

	png.Encode(os.Stdout, img)

}

func newton(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15
	//var v complex128
	for n := uint8(0); n < iterations; n++ {
		z -= (z*z*z*z - 1) / (4 * z * z * z)
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{contrast * (n + 1), contrast * (n + 2), n * contrast, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
