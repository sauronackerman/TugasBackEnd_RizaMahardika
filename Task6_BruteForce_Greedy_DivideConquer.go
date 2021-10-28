package main

import (
	"fmt"
	"math"
	"sort"
)

// SimpleEquations Problem 1. Simple Equation
func SimpleEquations(a, b, c int) {

	var sortedInput = []int{a, b, c}
	sort.Ints(sortedInput)
	x, y, z := 0, 0, 0
	for i := 1; i <= sortedInput[2]; i++ {
		for j := 1; j <= sortedInput[2]; j++ {
			for k := 1; k <= sortedInput[2]; k++ {
				if i+j+k == a && i*j*k == b && math.Pow(float64(i), 2)+math.Pow(float64(j), 2)+math.Pow(float64(k), 2) == float64(c) {
					x, y, z = k, j, i
				}
			}
		}
	}
	if x == 0 && y == 0 && z == 0 {
		fmt.Println("No Solutions")
	} else {
		fmt.Println(x, y, z)
	}
}

func testSimpleEquation(){
	fmt.Println("Hasil test: ")
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
	//SimpleEquations(7,9,5)
}

//Problem 2. Money Coins
func moneyCoins(money int) []int {
	pecahan := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	var output []int
	for i := len(pecahan) - 1; i >= 0; i-- {
		if money >= pecahan[i] {
			output = append(output, pecahan[i])
			money -= pecahan[i]
			i++
		}
	}
	return output
}

func testMoneyCoins()  {
	fmt.Println("Hasil test: ")
	fmt.Println(moneyCoins(123))   // [100 20 1 1 1]

	fmt.Println(moneyCoins(432))   // [200 200 20 10 1 1]

	fmt.Println(moneyCoins(543))   // [500, 20, 20, 1, 1, 1]

	fmt.Println(moneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]

	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]
}

//Problem 3. Dragon of Loowater
func DragonOfLoowater(dragonHead, knightHeight []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)
	count := 0
	for i := 0; i < len(dragonHead); i++ {
		for j := 0; j < len(knightHeight); j++ {
			result := knightHeight[j] - dragonHead[i]
			if result >= 0 {
				count += knightHeight[j]
				break
			}
		}
	}
	if count > knightHeight[len(knightHeight)-1] {
		fmt.Println(count)
	} else {
		fmt.Println("knight fall")
	}
}
func testDragonOfLoowater() {
	fmt.Println("Hasil test: ")
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})    // 11

	DragonOfLoowater([]int{5, 10}, []int{5})         // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}

//Problem 4. Binary Search Algorithm

func BinarySearch(arr []int, x int)  {
	//memastikan data sorted
	sort.Ints(arr)

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right+left)/2
		if x == arr[mid] {
			fmt.Println(mid)
			return
		}

		if x > arr[mid] {
			left = mid + 1
		} else {
			right = mid -1
		}
	}

	fmt.Println(-1)
}

func testBinarySearch() {
	fmt.Println("Hasil test: ")
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
func listProgramTask6()  {
	fmt.Println("=======================")
	fmt.Println("List Program Task - Riza Mahardika")
	fmt.Println("1. Simple Equation")
	fmt.Println("2. Money Coins")
	fmt.Println("3. Dragon of Loowater")
	fmt.Println("4. Binary Search Algorithm")
	fmt.Println("=======================")
	fmt.Println("Pilih Program yang akan dijakankan:")
}
func pilihProgramTask6()  {
	var pilih int
	fmt.Scanf("%d", &pilih)
	switch pilih{
	case 1:
		testSimpleEquation()
	case 2:
		testMoneyCoins()
	case 3:
		testDragonOfLoowater()
	case 4:
		testBinarySearch()
	default:
		fmt.Println("Input tidak dikenali")
	}
}

func main() {
//testSimpleEquation()
//	testMoneyCoins()
//	testDragonOfLoowater()
//	testBinarySearch()
	listProgramTask6()
	pilihProgramTask6()
}


