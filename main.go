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
	fmt.Println("1.1\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", que11([]int{1, 2, 3, 4, 5}))

	fmt.Println("1.2\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", que12([]int{1, 2, 3, 4, 5}))

	fmt.Println("1.3\n", "IN", []int{1, 2, 3, 4, 5})
	fmt.Println("OUT", que13([]int{1, 2, 3, 4, 5}))

	var (
		a []int = []int{1, 2, 3, 4, 5}
		b int
		c []int = []int{4, 5, 9, 10, 1}
	)

	fmt.Println("1.4\n", "IN", a)
	a, b = que14(a)
	fmt.Println("OUT", a, b)

	a = []int{1, 2, 3, 4, 5}
	fmt.Println("1.5\n", "IN", a)
	a, b = que15(a)
	fmt.Println("OUT", a, b)

	// fmt.Scan(&i) Нужен ли?
	a = []int{1, 2, 3, 4, 5}
	in_ind := 1
	fmt.Println("1.6\n", "IN", a, in_ind)
	a, b = que16(a, in_ind)
	fmt.Println("OUT", a, b)

	a = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("1.7\n", "IN", a, c)
	fmt.Println("OUT", que17(a, c))

	fmt.Println("1.8\n", "IN", a, c)
	fmt.Println("OUT", que18(a, c))

	fmt.Println("1.9\n", "IN", a)
	fmt.Println("OUT", que19(a))

	fmt.Println("1.10\n", "IN", a, in_ind)
	fmt.Println("OUT", que110(a, in_ind))

	fmt.Println("1.11\n", "IN", a)
	fmt.Println("OUT", que111(a))

	fmt.Println("1.12\n", "IN", a, in_ind)
	fmt.Println("OUT", que112(a, in_ind))

	fmt.Printf("1.13\n IN %v ADDR %p\n", a, a)
	que_113 := que113(a)
	fmt.Printf("OUT %v ADDR %p\n", que_113, que_113)

	fmt.Println("1.14\n", "IN", a)
	fmt.Println("OUT", que114(a))
}
