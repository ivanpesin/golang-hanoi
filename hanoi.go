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

func (p *pole) Peek() int {
	return (*p)[len(*p)-1]
}

func fromTo(poles [3]pole, a, b int) (int, int, bool) {
	if len(poles[a]) == 0 && len(poles[b]) == 0 {
		return -1, -1, true
	}
	if len(poles[a]) == 0 {
		return b, a, false
	}
	if len(poles[b]) == 0 {
		return a, b, false
	}
	if poles[a].Peek() > poles[b].Peek() {
		return b, a, false
	}
	return a, b, false
}

func main() {

	var n int // number of disks

	var poles [3]pole // 3 poles, each pole is a stack

	println("Towers of Hanoi")
	println("---------------")
	for n < 3 {
		print("Enter number of disks (>=3): ")
		fmt.Scanf("%d", &n)
	}

	finish := int(math.Pow(2, float64(n)) - 1)
	solution := [3][2]int{ // for even number of disks
		{0, 1},
		{0, 2},
		{1, 2},
	}
	if n%2 == 1 { // adjust for odd nuber of disks
		solution[0][1] = 2
		solution[1][1] = 1
	}

	for i := n; i > 0; i-- {
		poles[0].Push(i)
	}

	fmt.Printf("---[ %d ]------\n1: %v\n2: %v\n3: %v\n", step, poles[0], poles[1], poles[2])
	step++

MainLoop:
	for {
		for i := 0; i < 3; i++ {
			if src, tgt, skip := fromTo(poles, solution[i][0], solution[i][1]); !skip {
				poles[tgt].Push(poles[src].Pop())
				fmt.Printf("---[ %d ]------\n1: %v\n2: %v\n3: %v\n", step, poles[0], poles[1], poles[2])
				step++
				if step > finish {
					break MainLoop
				}
			}
		}
	}
	fmt.Printf("--> Optimal number of moves 2^n-1 = %.0f\n", math.Pow(2, float64(n))-1)
}
