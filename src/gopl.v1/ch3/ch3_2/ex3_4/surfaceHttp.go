package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

var height, width float64 = 300, 600
var cells float64 = 100           // number of grid cells
var xyrange = 30.0                // axis ranges (-xyrange..+xyrange)
var xyscale = width / 2 / xyrange // pixels per x or y unit
var zscale = height * 0.4         // pixels per z unit
var angle = math.Pi / 6           // angle of x, y axes (=30°)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	//http服务
	http.HandleFunc("/", handle)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if err := r.ParseForm(); err != nil {
		return
	}
	//var height int
	//var width  int
	for k, v := range r.Form {
		if k == "height" {
			h, _ := strconv.ParseFloat(v[0], 64)
			if h > 0 {
				height = h
			}
		}
		if k == "width" {
			w, _ := strconv.ParseFloat(v[0], 64)
			if w > 0 {
				width = w
			}
		}
	}
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: #ff0000; fill: #0000ff; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")

}
func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Cos(r)
}
