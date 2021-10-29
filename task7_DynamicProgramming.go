package main

import (
	"fmt"
	"math"
)

//Problem 1. Fibonacci Number Top-Down
var memo = map[int]int{}
func fibo(n int) int {
	if value, isCounted:= memo[n]; isCounted {
		return value
	} else if n <= 1 {
		memo[n] = n
	} else {
		memo[n] = fibo(n-1) + fibo(n-2)
	}
	return memo[n]
}
func testFiboTopDown()  {
	fmt.Println("Hasil Test Fibo Top Down:")
	fmt.Println(fibo(0))  // 0

	fmt.Println(fibo(1))  // 1

	fmt.Println(fibo(2))  // 1

	fmt.Println(fibo(3))  // 2

	fmt.Println(fibo(5))  // 5

	fmt.Println(fibo(6))  // 8

	fmt.Println(fibo(7))  // 13

	fmt.Println(fibo(9))  // 13

	fmt.Println(fibo(10)) // 55
	fmt.Println(fibo(50))
}

//Problem 2. Fibonacci Number Bottom-Up
func fiboBottomUp(n int) int {
	if n <= 1 {
		return n
	}
	fn1, fn2, fn := 1, 0, 0
	for i := 2; i <= n; i++ {
		fn = fn1 + fn2
		fn2 = fn1
		fn1 = fn
	}
	return fn
}

func testFiboBottomUp()  {
	fmt.Println("Hasil Test Fibo Bottom Up:")
	fmt.Println(fiboBottomUp(0))  // 0

	fmt.Println(fiboBottomUp(1))  // 1

	fmt.Println(fiboBottomUp(2))  // 1

	fmt.Println(fiboBottomUp(3))  // 2

	fmt.Println(fiboBottomUp(5))  // 5

	fmt.Println(fiboBottomUp(6))  // 8

	fmt.Println(fiboBottomUp(7))  // 13
 fmt.Println(fibo(0))  // 0

   fmt.Println(fibo(1))  // 1

   fmt.Println(fibo(2))  // 1

   fmt.Println(fibo(3))  // 2

   fmt.Println(fibo(5))  // 5

   fmt.Println(fibo(6))  // 8

   fmt.Println(fibo(7))  // 13

   fmt.Println(fibo(9))  // 13

   fmt.Println(fibo(10)) // 55
	fmt.Println(fiboBottomUp(9))  // 13

	fmt.Println(fiboBottomUp(10)) // 55
	fmt.Println(fiboBottomUp(50))
}

//Problem 3. Frog
func Frog(jumps []int) int {
	var minCost = map[int]int{}
	//2 val awal untuk kalkulasi jump 1 dan jump 2
	minCost[0] = 0
	minCost[1] = int(math.Abs(float64(jumps[1] - jumps[0])))
	//fmt.Println("Min awal",minCost)
	for i := 2; i < len(jumps); i++ {
		//fmt.Println("jum[i] i-1, i-2", jumps[i], jumps[i-1], jumps[i-2])
		//fmt.Println("mincost before", minCost[i-1], minCost[i-2])
		jump1 := int(math.Abs(float64(jumps[i] - jumps[i-1]))) + minCost[i-1]
		jump2 := int(math.Abs(float64(jumps[i] - jumps[i-2]))) + minCost[i-2]
		//fmt.Println("jump12", jump1, jump2)
		if jump1 < jump2 {
			minCost[i] = jump1
		} else {
			minCost[i] = jump2
		}
		//fmt.Println("test min",minCost)
	}
	return minCost[len(minCost)-1]
}

func testFrog()  {
	fmt.Println("Hasil test Frog:")
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

func listProgramTask7()  {
	fmt.Println("=======================")
	fmt.Println("List Program Task - Riza Mahardika")
	fmt.Println("1. Fibonacci Number Top-down")
	fmt.Println("2. Fibonacci Number Bottom-up")
	fmt.Println("3. Frog")
	fmt.Println("=======================")
	fmt.Println("Pilih Program yang ajumpskan dijakankan:")
}
func pilihProgramTask7()  {
	var pilih int
	fmt.Scanf("%d", &pilih)
	switch pilih{
	case 1:
		testFiboTopDown()
	case 2:
		testFiboBottomUp()
	case 3:
		testFrog()
	default:
		fmt.Println("Input tidak dikenali")
	}
}
func main() {
//testFiboTopDown()
//	testFiboBottomUp()
//	testFrog()
	listProgramTask7()
	pilihProgramTask7()
}
