package main

import (
	"fmt"
	"math"
)

type pole []int

var step int

func (p *pole) Push(a int) {
	*p = append(*p, a)
}

func (p *pole) Pop() int {
	r := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return r
}

func move(n int, p *[3]pole, src, tgt, aux int) {
	if n < 1 {
		return
	}

	move(n-1, p, src, aux, tgt)
	p[tgt].Push(p[src].Pop())
	step++
	fmt.Printf("---[ %d ]------\n1: %v\n2: %v\n3: %v\n", step, p[0], p[1], p[2])
	move(n-1, p, aux, tgt, src)
}

func main() {
	var p [3]pole
	var n int

	p[0] = pole{} // source
	p[1] = pole{} // aux
	p[2] = pole{} // target

	println("Towers of Hanoi")
	println("---------------")
	for n < 3 {
		print("Enter number of disks (>=3): ")
		fmt.Scanf("%d", &n)
	}

	for i := n; i > 0; i-- {
		p[0].Push(i)
	}

	fmt.Printf("---[ %d ]------\n1: %v\n2: %v\n3: %v\n", step, p[0], p[1], p[2])
	// move N disks src tgt aux
	move(n, &p, 0, 2, 1)
	fmt.Printf("--> Optimal number of moves 2^n-1 = %.0f\n", math.Pow(2, float64(n))-1)
}
