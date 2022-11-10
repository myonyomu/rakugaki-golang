package main

import "fmt"

func main() {
	/*
		slice appendって実際には元のarrayと同じ値を持ったarrayがallocateされてない？と思ったので検証。
		されてるし公式ブログに書いてある。
		https://go.dev/blog/slices

		> What if we want to grow the slice beyond its capacity? You can’t!
		> By definition, the capacity is the limit to growth.
		> But you can achieve an equivalent result by allocating a new array, copying the data over,
		> and modifying the slice to describe the new array.
	*/

	arr := [3]int{1, 2, 3}
	fmt.Println(&arr[0]) // 0xc00001a180
	sli := arr[:]
	fmt.Println(&sli[0]) // 0xc00001a180

	sli_reslice := arr[:1]
	fmt.Println(&sli_reslice[0]) // 0xc00001a180

	sli = append(sli, 4)
	fmt.Println(&sli[0])         // 0xc00000e450
	fmt.Println(&sli_reslice[0]) // 0xc00001a180
}
