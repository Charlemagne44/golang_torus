package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
)

// draw the circle -> (R2, 0, 0)  + (R1cos(theta), R1sin(theta), 0)
// rotate the circle about y (torus) -> ((R2 + R1cos(theta), R1sin(theta), -(R2 + R1cos(theta)sin(phi))))

// spacing between calculations of the theta angle of rotation
const theta_spacing = 0.07 // angle of rotation about x that creates circle

// the phi axis the torus rotates around
const phi_spacing = 0.02 // angle of rotation about y that creates torus from circle

// y vector from center of circle to edge (all of which rotating about axis)
const R1 = 1

// x vector to origin on circle being rotated about axis
const R2 = 2

// distance of object from the viewer
const K2 = 5

// Initial screen clear function maps
var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	CallClear()

	var screen_width = 25
	flag.IntVar(&screen_width, "w", 25, "The width of the donut in terminal characters")
	flag.Parse()

	// calculate k1 (viewer distance from screen) based on the screen size
	var K1 = screen_width * K2 * 3 / (8 * (R1 + R2))

	var a float64 = 90
	var b float64 = 0
	for {
		render_frame(a, b, screen_width, K1)
		a += 0.001
		b += 0.001
	}
}

// A and B are the angles of which the torus will be rotating about the other two axis
func render_frame(A, B float64, screen_width, K1 int) {
	var cosA float64 = math.Cos(A)
	var sinA float64 = math.Sin(A)
	var cosB float64 = math.Cos(B)
	var sinB float64 = math.Sin(B)

	// initialize output and zbuffer
	var output [][]string = make([][]string, screen_width)
	for i, _ := range output {
		output[i] = make([]string, screen_width)
	}
	for i := 0; i < screen_width; i++ {
		for j := 0; j < screen_width; j++ {
			output[i][j] = " "
		}
	}

	var zbuffer [][]int = make([][]int, screen_width)
	for i, _ := range output {
		zbuffer[i] = make([]int, screen_width)
	}
	for i := 0; i < screen_width; i++ {
		for j := 0; j < screen_width; j++ {
			zbuffer[i][j] = 0
		}
	}

	// calculating the circle that gets rotated to make up the torus
	for theta := float64(0); theta < float64(2)*math.Pi; theta += theta_spacing {
		var costheta float64 = math.Cos(theta)
		var sintheta float64 = math.Sin(theta)

		// calculating phi (rotating the circle about the y axis to make the torus)
		for phi := float64(0); phi < float64(2)*math.Pi; phi += phi_spacing {
			var cosphi float64 = math.Cos(phi)
			var sinphi float64 = math.Sin(phi)

			// calculate the x, y of the circle before revolution
			var circlex float64 = R2 + R1*costheta
			var circley float64 = R1 * sintheta

			// All revolutions for final (x,y,z)
			var x float64 = circlex*(cosB*cosphi+sinA*sinB*sinphi) - circley*cosA*sinB
			var y float64 = circlex*(sinB*cosphi-sinA*cosB*sinphi) + circley*cosA*cosB
			var z float64 = K2 + cosA*circlex*sinphi + circley*sinA
			var oneOverZ float64 = 1 / z

			// calculate x and y projection
			var xp int = int(screen_width/2 + int(float64(K1)*oneOverZ*x))
			var yp int = int(screen_width/2 - int(float64(K1)*oneOverZ*y))

			// calculate the luminance (abridged)
			var luminance float64 = cosphi*costheta*sinB - cosA*costheta*sinphi -
				sinA*sintheta + cosB*(cosA*sintheta-costheta*sinA*sinphi)

			if luminance > 0 {
				if oneOverZ > float64(zbuffer[xp][yp]) {
					zbuffer[xp][yp] = int(oneOverZ)
					var luminance_index float64 = luminance * 8
					var charset string = ".,-~:;=!*#$@"
					output[xp][yp] = string(charset[int(luminance_index)])
				}
			}
		}
	}

	// dump the output vector onto the screen
	fmt.Printf("\x1b[H")
	for j := 0; j < screen_width; j++ {
		for i := 0; i < screen_width; i++ {
			fmt.Printf("%s%s", output[i][j], output[i][j])
		}
		fmt.Println()
	}

}
