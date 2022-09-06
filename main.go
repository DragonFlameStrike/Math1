package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var a, b, c, delta, epsilon float64

func main() {
	if len(os.Args) != 6 {
		fmt.Printf("usage:   a b c delta epsilon\n")
		fmt.Printf("example: 1 1 1 2 0.0001\n")
		return
	}
	var err bool
	//a = 1
	//b = 1
	//c = 1
	//delta = 2
	//epsilon = 0.01
	a, b, c, delta, epsilon, err = readData(os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	if err {
		fmt.Print("bad input")
		return
	}
	answer := solution()
	fmt.Printf("answer %v", answer)
	return
}

func solution() []float64 {
	// x^3+ax^2+bx+c=0
	// 3x^2+2ax+b=0
	var output []float64
	var alpha float64
	var beta float64
	d := desc()
	if d > 0 {
		alpha = (-2*a - math.Sqrt(d)) / (2 * 3)
		beta = (-2*a + math.Sqrt(d)) / (2 * 3)
		fa := f(alpha)
		fb := f(beta)
		if fa > epsilon && fb > epsilon {
			temp := iter(alpha, math.MinInt)
			return append(output, temp)
		}
		if fa < (-epsilon) && fb < (-epsilon) {
			temp := iter(beta, math.MaxInt)
			return append(output, temp)
		}
		if (fa < (-epsilon) && fb > (epsilon)) || (fa > (epsilon) && fb < (-epsilon)) {
			temp := iter(beta, math.MaxInt)
			output = append(output, temp)
			temp2 := iter(alpha, math.MinInt)
			output = append(output, temp2)
			root := findRoot(alpha, beta)
			output = append(output, root)
			return output
		}
		if math.Abs(fa) > (epsilon) && math.Abs(fb) < (epsilon) {
			output = append(output, beta)
			temp := iter(alpha, math.MinInt)
			return append(output, temp)
		}
		if math.Abs(fa) < (epsilon) && math.Abs(fb) > (epsilon) {
			output = append(output, alpha)
			temp := iter(beta, math.MaxInt)
			return append(output, temp)
		}
		if math.Abs(fa) < (epsilon) && math.Abs(fb) < (epsilon) {
			return append(output, (alpha+beta)/2)
		}

	} else {
		r := f(0)
		if math.Abs(r) < epsilon {
			return append(output, 0)
		} else if r > epsilon {
			temp := iter(0, math.MinInt)
			return append(output, temp)
		} else {
			temp := iter(0, math.MaxInt)
			return append(output, temp)
		}
	}
	return nil
}

func findRoot(alpha float64, beta float64) float64 {
	for {
		fa := f(alpha)
		fb := f(beta)
		if math.Abs(fa) < epsilon {
			return alpha
		}
		if math.Abs(fb) < epsilon {
			return beta
		}
		z := (alpha + beta) / 2
		fz := f(z)
		if fz*fa < 0 {
			beta = z
		} else if fz*fb < 0 {
			alpha = z
		} else {
			return z
		}
	}
}

func iter(start float64, direction int) float64 {
	fa := f(start)
	alpha := start
	var cdelta float64
	if direction == math.MaxInt {
		cdelta = delta
	} else {
		cdelta = -delta
	}
	beta := start + cdelta
	fb := f(start + cdelta)
	var i float64
	for i = 2; fa*fb > 0; i++ {
		fa = fb
		alpha = beta
		fb = f(start + i*cdelta)
		beta = start + i*cdelta
	}
	return findRoot(alpha, beta)
}

func f(point float64) float64 {
	return math.Pow(point, 3) + a*math.Pow(point, 2) + b*point + c
}

func desc() float64 {
	return 4*a*a - 4*3*b
}

func readData(a string, b string, c string, delta string, epsilon string) (float64, float64, float64, float64, float64, bool) {
	na, e1 := strconv.ParseFloat(a, 32)
	nb, e2 := strconv.ParseFloat(b, 32)
	nc, e3 := strconv.ParseFloat(c, 32)
	ndelta, e4 := strconv.ParseFloat(delta, 32)
	nepsilon, e5 := strconv.ParseFloat(epsilon, 32)
	if e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil {
		return 0, 0, 0, 0, 0, true
	}
	return na, nb, nc, ndelta, nepsilon, false
}
