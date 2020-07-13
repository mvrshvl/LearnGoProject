package main

func que11(in []int) (out []int) {
	for i := 0; i < len(in); i++ {
		in[i]++
	}
	return in
}
func que12(in []int) (out []int) {
	in = append(in, 5)
	return in
}
