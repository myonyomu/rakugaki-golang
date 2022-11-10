package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
	}

	for y := range result {
		for x := range result[y] {
			result[y][x] = uint8((x ^ y*3) / 2)
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
