package main

import (
	"math"
	"strings"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var t float64
	for {
		z, t = z-(z*z-x)/(2*z), z
		if math.Abs(t-z) < 1e-6 {
			break
		}
	}
	return z
}

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		a[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			a[i][j] = uint8(i * j)
		}
	}
	return a
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		m[v]++
	}
	return m
}

func fibonacci() func() int {
	n := 0
	a := 0
	b := 1
	c := a + b
	return func() int {
		var ret int
		switch {
		case n == 0:
			n++
			ret = 0
		case n == 1:
			n++
			ret = 1
		default:
			ret = c
			a = b
			b = c
			c = a + b
		}
		return ret
	}
}
