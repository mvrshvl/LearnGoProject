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

func que13(in []int) (out []int) {
	out = append(out, 5)
	out = append(out, in[:]...)
	return
}

func que14(in []int) (out []int, last int) {
	last = in[len(in)-1]
	out = append(in[0 : len(in)-1])
	return
}

func que15(in []int) (out []int, first int) {
	first = in[0]
	out = append(in[1:])
	return
}

func que16(in []int, i_in int) (out []int, i_out int) {
	i_out = in[i_in]
	out = append(in[0:i_in], in[i_in+1:]...)
	return
}
