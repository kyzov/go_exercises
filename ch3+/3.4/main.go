package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"colors"
	"surface"

	"image/color"

)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		width := parseFirstIntOrDefault(r.Form["width"], 600)
		height := parseFirstIntOrDefault(r.Form["height"], 320)
		cells := parseFirstIntOrDefault(r.Form["size"], 100)
		xyrange := parseFirstFloat64OrDefault(r.Form["xyrange"], 30)
		xyscale := parseFirstFloat64OrDefault(r.Form["xyscale"], float64(width/2)/xyrange)
		zscale := parseFirstFloat64OrDefault(r.Form["zscale"], float64(height)*0.4)
		angle := parseFirstFloat64OrDefault(r.Form["angle"], math.Pi/6)
		topColor := parseFirstColorOrDefault(r.Form["topColor"], color.RGBA{0xff, 0x00, 0x00, 0xff})
		bottomColor := parseFirstColorOrDefault(r.Form["bottomColor"], color.RGBA{0x00, 0x00, 0xff, 0xff})
		w.Header().Set("Content-Type", "image/svg+xml")
		surface.Render(w, width, height, cells, xyrange, xyscale, zscale, angle, topColor, bottomColor)
	}
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func parseFirstIntOrDefault(array []string, defaultValue int) int {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.Atoi(array[0])
	if err != nil {
		return defaultValue
	}
	return value
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


func parseFirstColorOrDefault(array []string, defaultValue color.Color) color.Color {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := colors.ColorFromString(array[0])
	if err != nil {
		return defaultValue
	}
	return value
}