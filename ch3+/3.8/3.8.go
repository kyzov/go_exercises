package main

import (
	"complexfloat"
	"complexrat"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

const (
	width, height = 1024, 1024
	contrast      = 15
)

var iterationsFlag = flag.Int("iterations", 200, "iterations: int")
var verboseFlag = flag.Bool("verbose", false, "verbose: bool")
var xFlag = flag.Int("x", 0, "center x: int")
var yFlag = flag.Int("y", 0, "center y: int")
var precFlag = flag.Uint("prec", 1024, "bigfloat prec: uint")
var zoomFlag = flag.Uint("zoom", 0, "zoom bit: uint")

func mainBigFloat() {
	xcenter := float64(*xFlag)
	ycenter := float64(*yFlag)
	invzoom := 2 / float64(int64(1)<<*zoomFlag)
	prec := *precFlag

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dy := big.NewFloat(invzoom * 2 / height)
	dx := big.NewFloat(invzoom * 2 / width)
	y := big.NewFloat(-invzoom + ycenter).SetPrec(prec)
	for py := 0; py < height; py++ {
		x := big.NewFloat(-invzoom + xcenter).SetPrec(prec)
		for px := 0; px < width; px++ {
			z := &complexfloat.ComplexFloat{
				Re:   x,
				Im:   y,
				Prec: prec,
			}
			img.Set(px, py, mandelbrotBigFloat(z))
			x.Add(x, dx)
		}
		y.Add(y, dy)
		if *verboseFlag {
			fmt.Fprintf(os.Stderr, "bflo (%d/%d)\n", py, height)
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrotBigFloat(z *complexfloat.ComplexFloat) color.RGBA {
	prec := *precFlag
	iterations := uint8(*iterationsFlag)
	two := big.NewFloat(2).SetPrec(prec)
	v := &complexfloat.ComplexFloat{
		Re:   big.NewFloat(0).SetPrec(prec),
		Im:   big.NewFloat(0).SetPrec(prec),
		Prec: prec,
	}
	for n := uint8(0); n < iterations; n++ {
		v.Square().Add(z)
		if v.AbsCompare(two) > 0 {
			return color.RGBA{contrast * (n + 20), contrast * (n + 21), n * contrast, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

func maiBigRat() {
	xcenter := int64(*xFlag)
	ycenter := int64(*yFlag)
	zoom := int64(1) << *zoomFlag

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dy := big.NewRat(4, height*zoom)
	dx := big.NewRat(4, width*zoom)
	y := big.NewRat(ycenter*zoom-2, zoom)
	for py := 0; py < height; py++ {
		x := big.NewRat(xcenter*zoom-2, zoom)
		for px := 0; px < width; px++ {
			z := &complexrat.ComplexRat{
				Re: x,
				Im: y,
			}
			img.Set(px, py, mandelbrotBigRat(z))
			x.Add(x, dx)
		}
		if *verboseFlag {
			fmt.Fprintf(os.Stderr, "brat (%d/%d)\n", py, height)
		}
		y.Add(y, dy)
	}
	png.Encode(os.Stdout, img)
}

func mandelbrotBigRat(z *complexrat.ComplexRat) color.Color {
	iterations := uint8(*iterationsFlag)
	two := big.NewRat(2, 1)
	v := &complexrat.ComplexRat{
		Re: big.NewRat(0, 1),
		Im: big.NewRat(0, 1),
	}
	for n := uint8(0); n < iterations; n++ {
		v.Square().Add(z)
		if v.AbsCompare(two) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func maincomplex64() {
	const (
		xmin, ymin, xmax, ymax = -1.5, -1.5, 1.5, 1.5
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot64(z))
		}
	}

	png.Encode(os.Stdout, img)

}

func mandelbrot64(z complex64) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cplx64.Abs(v) > 2 {
			return color.RGBA{contrast * (n + 20), contrast * (n + 21), n * contrast, 255} //<-Упражнение 3.5
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
