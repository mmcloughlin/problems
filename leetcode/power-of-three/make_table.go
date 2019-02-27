// +build ignore

package main

import (
	"fmt"
)

func main() {
	powers := []int{}
	x := 1
	for {
		next := x * 3
		if next < x {
			break
		}
		powers = append(powers, x)
		x = next
	}

	fmt.Printf("package solution\n\nvar powers = map[int]bool{\n")
	for _, p := range powers {
		fmt.Printf("\t%d: true,\n", p)
	}
	fmt.Printf("}\n")
}
