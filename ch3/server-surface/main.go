// Server returns Surface SVG
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320             // canvas size in pixels
	cells         = 100                  // number of grid cells
	xyrange       = 30.0                 // axis ranges (-xyrange..+xyrange)
	angle         = math.Pi / 6          // angle of x, y axes (=30°)
	peak, valley  = "#ff0000", "#0000ff" // colorcodes for polygon fill
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(rw http.ResponseWriter, r *http.Request) {
	w, h, cp, cv := getArgs(r)
	rw.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(rw, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", w, h)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, color := corner(i+1, j, w, h, cp, cv)
			bx, by, _ := corner(i, j, w, h, cp, cv)
			cx, cy, _ := corner(i, j+1, w, h, cp, cv)
			dx, dy, _ := corner(i+1, j+1, w, h, cp, cv)
			fmt.Fprintf(rw, "<polygon points='%g,%g %g,%g %g,%g %g,%g'"+
				" style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(rw, "</svg>")
}

func getArgs(r *http.Request) (int, int, string, string) {
	var w, h int
	var cp, cv string
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}
	// width
	if r.Form["width"] != nil {
		w, _ = strconv.Atoi(r.Form["width"][0])
	} else {
		w = width
	}
	// height
	if r.Form["height"] != nil {
		h, _ = strconv.Atoi(r.Form["height"][0])
	} else {
		h = height
	}
	// color: peaks
	if r.Form["peak"] != nil {
		cp = r.Form["peak"][0]
	} else {
		cp = peak
	}
	// color: valleys
	if r.Form["valley"] != nil {
		cv = r.Form["valley"][0]
	} else {
		cv = valley
	}

	return w, h, cp, cv
}

func corner(i, j int, w, h int, cp, cv string) (float64, float64, string) {

	// Find point (x,y) at corner of cell (i, j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	xyscale := float64(w) / 2 / xyrange // pixels per x or y unit
	zscale := float64(h) * 0.4          // pixels per z unit

	// Compute surface height x.
	var z float64
	for z == 0 || math.IsInf(z, 0) {
		z = f(x, y)
	}

	var color string
	if z > 0 {
		color = cp // peak
	} else {
		color = cv // valleys
	}

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(w)/2 + (x-y)*cos30*xyscale
	sy := float64(h)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
