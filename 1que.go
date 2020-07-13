package main

import "strings"

func Sqrt(x float64) float64 {
	y := 1.0
	for i := 0; i < 10; i++ {
		y = y - ((y*y - x) / (2 * y))
	}
	return y
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
