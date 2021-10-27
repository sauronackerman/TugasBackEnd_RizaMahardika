package main

import (
	"fmt"
	"math"
	"strconv"
)

//Problem 1 - Prima ke X
func primeX(number int) int {
	var indexPrime int
	var i int
	for i= 2; i >= 2; i++ {
		isPrime := true
		for j:= 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				j = i
			}

		}
		if isPrime {
			indexPrime++
			fmt.Println("test i", i)
		}
		if indexPrime == number {
			return i
		}

	}
	return i
}
func testPrimeX() {
	fmt.Println("Hasil test:")
	fmt.Println(primeX(1))  // 2

	fmt.Println(primeX(5))  // 11

	fmt.Println(primeX(8))  // 19

	fmt.Println(primeX(9))  // 23

	fmt.Println(primeX(10)) // 29
}

//Problem 2. Fibonacci Recursive
func fibonacci(number int) int {
	if number <=1 {
		return number
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}
}
func testFibonacci()  {
	fmt.Println("Hasil Test:")
	fmt.Println(fibonacci(0))  // 0

	fmt.Println(fibonacci(2))  // 1

	fmt.Println(fibonacci(9))  // 34

	fmt.Println(fibonacci(10)) // 55

	fmt.Println(fibonacci(12)) // 144
}

//Problem 3. Prima Segi Empat
func primaSegiEmpat(high, wide, start int) {
	var indexPrime int
	var i int
	var slicePrime []int
	for i= start + 1; i >= start + 1; i++ {
		isPrime := true
		for j:= 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				j = i
			}

		}
		if isPrime {
			indexPrime++
			slicePrime = append(slicePrime, i)
		}

		if indexPrime == high*wide {
			var j int
			var sum int
			for _,v := range slicePrime {
				sum += v
			}
			for i := 0; i < len(slicePrime); i+= high {
				j += high
				if j > len(slicePrime){
					j = len(slicePrime)
				}
				x := slicePrime[i:j]
				fmt.Printf("%+v \n", x )

			}
			fmt.Println(sum)
			break
		}
	}
}
func testPrimaSegiEmpat()  {
	primaSegiEmpat(2, 3, 13)
	/*

	   17 19

	   23 29

	   31 37

	   156

	*/

	primaSegiEmpat(5, 2, 1)

	/*

	   2  3  5  7 11

	   13 17 19 23 29

	   129

	*/
}

//Problem 4. Total Maksimum Dari Deret Bilangan
func MaxSequence(arr []int) int {
	i := 0
	j := len(arr) - 1
	var max int
	for i != j {
		var sum int
		for k := i; k <= j; k++ {
			//fmt.Println("ini arr[k]",arr[k])
			sum += arr[k]
			//fmt.Println("sum after +arrk", sum)
		}
		if sum > max {
			//fmt.Println("max before sum", max)
			max = sum
			//fmt.Println("max after sum", max)
		}
		if arr[i] < arr[j] {
			i++
		} else {
			j--
		}
	}
	return max

}
func testMaxSequence()  {
	fmt.Println(MaxSequence([]int{-2,1,-3,4,-1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7

	fmt.Println(MaxSequence([]int{-2,-3, 4, -1, -2, 1, 5, -3}))   // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2,-3, 1, 6, -6}))   // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}

//Problem 5. Find Min and Max Number
func FindMinAndMax(arr []int) string {
	minVal := arr[0]
	maxVal := arr[0]
	var minIndex, maxIndex int
	for i := 0; i < (len(arr) - 1); i++ {
		if arr[i+1] < minVal {
			minVal = arr[i+1]
			minIndex = i + 1
		}
		if arr[i+1] > maxVal {
			maxVal = arr[i+1]
			maxIndex = i + 1
		}
	}
	result := "min: " + fmt.Sprint(minVal) + " index: " + fmt.Sprint(minIndex) + " max: " + fmt.Sprint(maxVal) + " index: " + fmt.Sprint(maxIndex)
	return result

}
func testFindMinAndMax() {
	fmt.Println("Hasil test:")
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))

	// min: -2 index: 3 max: 8 index: 5

	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))

	// min: -5 index: 1 max: 22 index: 3

	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))

	// min: -21 index: 4 max: 9 index: 2

	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))

	// min: -1 index: 0 max: 18 index: 5

	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))

	// min: -20 index: 5 max: 7 index: 4
}
//Problem 6. Max Buy Product
func MaximumBuyProduct(money int, productPrice []int) {
	totalProduct := 0
	totalBuy := 0
	for k := range productPrice {
		min := k
		for j := k + 1; j < len(productPrice); j++ {
			if productPrice[j] < productPrice[min] {
				min = j
			}
		}
		productPrice[k], productPrice[min] = productPrice[min], productPrice[k]
	}

	for i := range productPrice {
		if totalBuy+productPrice[i] > money {
			fmt.Println(totalProduct)
			break
		} else if totalBuy+productPrice[i] == money {
			totalProduct++
			fmt.Println(totalProduct)
			break
		} else {
			totalBuy += productPrice[i]
			totalProduct++
		}
	}

}
func testMaxBuyProduct()  {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3

	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4

	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4

	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1

	MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}
//Problem 7. Playing Domino
func playingDomino(cards [][]int, deck []int) (result interface{}) {
	var max int
	for i := 0; i < len(cards); i++ {
		for j := 0; j < 2; j++ {
			if cards[i][j] == deck[0] || cards[i][j] == deck[1] {
				if cards[i][0]+cards[i][1] > max {
					max = cards[i][0] + cards[i][1]
					result = []int{cards[i][0], cards[i][1]}
				}
			}
		}
	}
	if result == nil {
		return "tutup kartu"
	} else {
		return result
	}

}

func testPlayingDomino()  {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))

	// [3, 4]

	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))

	// [6 5]

	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))

	// “tutup kartu”
}

//Problem 8. Most Appear Item
type pair interface{}

type appear struct {
	item  string
	appear int
}
func MostAppearItem(items []string) []pair {
	itemAppear := make(map[string]int)
	for _, value := range items {
		itemAppear[value] += 1
		//fmt.Println("ini test", itemAppear)
	}
	sortedResult := make([]appear, len(itemAppear))
	index := 0
	for k, value := range itemAppear {
		sortedResult[index].item = k
		sortedResult[index].appear = value
		index++
	}
	for i := range sortedResult {
		for j := i + 1; j < len(sortedResult); j++ {
			if sortedResult[i].appear > sortedResult[j].appear {
				temp := sortedResult[i]
				sortedResult[i] = sortedResult[j]
				sortedResult[j] = temp
			}
		}
	}
	var result []pair
	for _, v := range sortedResult {
		result = append(result, v.item+"->"+strconv.Itoa(v.appear))
	}

	return result

}
func testMostAppear()  {
	fmt.Println("Hasil test:")
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}

func listProgramTaskRecursive()  {
	fmt.Println("=======================")
	fmt.Println("List Program Task - Riza Mahardika")
	fmt.Println("1. Prima ke X")
	fmt.Println("2. Fibonacci (Recursive)")
	fmt.Println("3.  Prima Segi Empat")
	fmt.Println("4. Total Maksimum Dari Deret Bilangan")
	fmt.Println("5. Find Min and Max Number")
	fmt.Println("6. Maximum Buy Product")
	fmt.Println("7. Playing Domino")
	fmt.Println("8. Most Appear Item")
	fmt.Println("=======================")
	fmt.Println("Pilih Program yang akan dijakankan:")
}
func pilihProgramTaskRecursive()  {
	var pilih int
	fmt.Scanf("%d", &pilih)
	switch pilih{
	case 1:
		testPrimeX()
	case 2:
		testFibonacci()
	case 3:
		testPrimaSegiEmpat()
	case 4:
		testMaxSequence()
	case 5:
		testFindMinAndMax()
	case 6:
		testMaxBuyProduct()
	case 7:
		testPlayingDomino()
	case 8:
		testMostAppear()
	default:
		fmt.Println("Input tidak dikenali")
	}
}


func main() {
//testPrimeX()
//testFibonacci()
//testPrimaSegiEmpat()
//	testMaxSequence()
//	testFindMinAndMax()
//	testMaxBuyProduct()
//	testPlayingDomino()
//	testMostAppear()
	listProgramTaskRecursive()
	pilihProgramTaskRecursive()
}
