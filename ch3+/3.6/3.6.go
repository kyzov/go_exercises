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
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img1 := image.NewRGBA(image.Rect(0, 0, width*2, height*2))
	for py := 0; py < height*2; py++ {
		y := float64(py)/(height*2)*(ymax-ymin) + ymin
		for px := 0; px < width*2; px++ {
			x := float64(px)/(width*2)*(xmax-xmin) + xmin
			z := complex(x, y)
			img1.SetRGBA(px, py, mandelbrot(z))
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height*2; py = py + 2 {
		for px := 0; px < width*2; px = px + 2 {
			var c = []color.RGBA{
				img1.RGBAAt(px, py),
				img1.RGBAAt(px+1, py),
				img1.RGBAAt(px, py+1),
				img1.RGBAAt(px+1, py+1),
			}
			var pr, pg, pb, pa int
			for n := 0; n < 4; n++ {
				pr += int(c[n].R)
				pg += int(c[n].G)
				pb += int(c[n].B)
				pa += int(c[n].A)
			}
			img.SetRGBA(px/2, py/2, color.RGBA{uint8(pr / 4), uint8(pg / 4), uint8(pb / 4), uint8(pa / 4)})
		}
	}

	png.Encode(os.Stdout, img)

}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{contrast * (n + 20), contrast * (n + 21), n * contrast, 255} //<-Упражнение 3.5
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
