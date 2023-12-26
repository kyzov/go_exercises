package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	var dotx, doty, zoom float64
	fmt.Println("Введите точку x:")
	fmt.Scan(&dotx)
	fmt.Println("Введите точку y:")
	fmt.Scan(&doty)
	fmt.Println("Введите приблежение: ")
	fmt.Scan(&zoom)

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		x := parseFirstFloat64OrDefault(r.Form["x"], dotx)
		y := parseFirstFloat64OrDefault(r.Form["y"], doty)
		zoom := parseFirstFloat64OrDefault(r.Form["zoom"], zoom)
		renderPNG(w, x, y, zoom)
	}
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func renderPNG(out io.Writer, x, y, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	m := math.Exp2(1 - zoom)
	xmin, xmax := x-m, x+m
	ymin, ymax := y-m, y+m

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func parseFirstFloat64OrDefault(array []string, defaultValue float64) float64 {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.ParseFloat(array[0], 64)
	if err != nil {
		return defaultValue
	}
	return value
}
