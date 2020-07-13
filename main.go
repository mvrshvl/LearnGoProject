package main

import "fmt"

func main() {
	/*--------------ДЗ1----------------------------------*/
	//1.1
	//fmt.Println(Sqrt(2))
	//1.2
	//pic.Show(Pic)
	//1.3
	//wc.Test( WordCount)
	//1.4
	//f := fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(f())
	//}
	/*---------------------------------------------------*/
	//2.1.1
	fmt.Println(que11([]int{1, 2, 3, 4, 5}))
	//2.1.2
	fmt.Println(que12([]int{1, 2, 3, 4, 5}))
	//2.1.3
	fmt.Println(que13([]int{1, 2, 3, 4, 5}))
	//2.1.4
	var (
		a []int = []int{1, 2, 3, 4, 5}
		b int
	)
	a, b = que14(a)
	fmt.Println(a, b)
	//2.1.5
	a = []int{1, 2, 3, 4, 5}
	a, b = que15(a)
	fmt.Println(a, b)
	//2.1.6
	// fmt.Scan(&i) Нужен ли?
	a = []int{1, 2, 3, 4, 5}
	a, b = que16(a, 1)
	fmt.Println(a, b)
}
