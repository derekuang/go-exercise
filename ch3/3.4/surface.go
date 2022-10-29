// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var width, height int       // canvas size in pixels
var xyscale, zscale float64 // pixels per x, y, z unit

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

// !+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	var err error
	width, err = strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil {
		width = 600
	}

	height, err = strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		height = 320
	}

	xyscale = float64(width / 2 / xyrange) // pixels per x or y unit
	zscale = float64(height) * 0.4         // pixels per z unit

	svg(w)
}

//!-handler

func svg(out http.ResponseWriter) {
	out.Header().Set("Content-Type", "image/svg+xml")

	out.Write([]byte(fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			out.Write([]byte(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)))
		}
	}
	out.Write([]byte("</svg>"))
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if r == 0 {
		return math.MaxFloat64
	} else {
		return math.Sin(r) / r
	}
}
