package gradient

import (
	"math"
)


// gradient just presents the gradient axis for the source
func (G *Gradient) gradient(steps int) ([]uint, []uint, []uint) {

	var ( //stores all binds
		Red		[]uint = make([]uint, 0) //red
		Green 	[]uint = make([]uint, 0) //green
		Blue 	[]uint = make([]uint, 0) //blue

		rounded float64 = 0 //amount of addition on each rotation
	)

	delta := (float64(len(G.Colours)) - 1) / float64(steps-1) //upwards force on every rotation

	//loops through all steps given properly
	for step := 0; step < steps; step++ {

		curve, force := Divmod(float64(rounded), 1)
		if curve >= float64(len(G.Colours) - 1) {
			curve = float64(len(G.Colours) - 2)
			force = 1.0
		}

		curvature := int(curve) // Works the current curve out
		Red	  = append(Red, uint(math.Round(float64(G.Colours[curvature].Red) * (1 - force) + float64(G.Colours[curvature + 1].Red) * force)))
		Green = append(Green, uint(math.Round(float64(G.Colours[curvature].Green) * (1 - force) + float64(G.Colours[curvature + 1].Green) * force)))
		Blue  = append(Blue, uint(math.Round(float64(G.Colours[curvature].Blue) * (1 - force) + float64(G.Colours[curvature + 1].Blue) * force)))
		rounded += delta
	}

	return Red, Green, Blue
}


//the divmod equivalent in golang
func Divmod(x, y float64) (float64, float64) {
	return math.Floor(float64(x / y)), x - y*math.Floor(float64(x / y))
}
