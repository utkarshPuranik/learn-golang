package main

const delta = 0.00001

func isNearEnough(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	return d < delta
}

func Sqrt(x float64) (float64, int) {
	temp := 0.0
	z := 1.0
	count := 0
	for {
		count += 1
		temp = z - (z*z-x)/2*z
		if d := temp - z; isNearEnough(d) {
			return temp, count
		}
		z = temp
	}
}

//Value of my sqrt is 1.414713296482291 with 499975 iterations (delta=0.001) and main sqrt is 1.4142135623730951
//Value of my sqrt is 1.4142635597210849 with 49999971 iterations (delta=0.0001) and main sqrt is 1.4142135623730951
