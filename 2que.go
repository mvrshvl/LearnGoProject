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

func que17(in_1 []int, in_2 []int) (out []int) {
	out = append(in_1[:], in_2[:]...)
	return
}

func que18(in_1 []int, in_2 []int) (out []int) {
	for i := 0; i < len(in_2); i++ {
		for j := 0; j < len(in_1); j++ {
			if in_2[i] != in_1[j] {
				out = append(out, in_1[j])
			}
		}
		in_1 = []int{}
		in_1 = append(in_1, out[:]...)
		out = []int{}
	}
	return in_1
}

func que19(in []int) (out []int) {
	tmp := in[0]
	for i := 0; i < len(in)-1; i++ {
		in[i] = in[i+1]
	}
	in[len(in)-1] = tmp
	return in
}
func que110(in []int, count int) (out []int) {
	for j := 0; j < count; j++ {
		tmp := in[0]
		for i := 0; i < len(in)-1; i++ {
			in[i] = in[i+1]
		}
		in[len(in)-1] = tmp
	}
	return in
}

func que111(in []int) (out []int) {
	tmp := in[len(in)-1]
	for i := len(in) - 1; i > 0; i-- {
		in[i] = in[i-1]
	}
	in[0] = tmp
	return in
}

func que112(in []int, count int) (out []int) {
	for j := 0; j < count; j++ {
		tmp := in[len(in)-1]
		for i := len(in) - 1; i > 0; i-- {
			in[i] = in[i-1]
		}
		in[0] = tmp
	}
	return in
}
