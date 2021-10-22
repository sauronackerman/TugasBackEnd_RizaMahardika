package main

import (
	"fmt"
	"math"
	"strconv"
)

//Problem 1 Bilangan Prima, hanya ditambah sqrt pada loop, fungsi belum disederhanakan
//func primeNumber(input int) bool {
//	var prima bool
//	if input <= 2 {
//		switch {
//		case input == 2:
//			prima = true
//		default:
//			prima = false
//		}
//	} else {
//
//		//for i := 2; i < int(math.Sqrt(float64(input))); i++ {
//		test := int(math.Sqrt(float64(input)))
//		fmt.Println(test)
//		for i := 2; i < test; i++ {
//			//for i := 2; i < input; i++ {
//			if input%i == 0 {
//
//				prima = false
//				break
//			} else {
//				if i < int(input)-1 {
//					fmt.Println(i,"cek")
//					prima = false
//				}
//				 prima = true
//
//				break
//			}
//		}
//	}
//return prima
//
//}
//penyederhanaan fungsi
func primeNumber(input int) bool {
	if input < 2 {
		return false
	} else {

		for i := 2; i < int(math.Sqrt(float64(input))); i++ {
			if input%i == 0 {
				return false
			}
			}
	}
	return true

}

func mencariBilanganPrima2() {
	fmt.Println(primeNumber(1000000007)) // true

	fmt.Println(primeNumber(1500450271)) // true

	fmt.Println(primeNumber(1000000000)) // false

	fmt.Println(primeNumber(10000000019)) // true

	fmt.Println(primeNumber(10000000033)) // true
	var input int
	fmt.Println("Masukkan bilangan integer")
	fmt.Scanf("%d", &input)
	bilangan := strconv.Itoa(input)
	//fmt.Println(math.Sqrt(float64(input)))
	//fmt.Println(int(math.Sqrt(float64(input))))
	hasil := primeNumber(input)


	if hasil {
		fmt.Println(bilangan, "Adalah bilangan prima")
	} else {
		fmt.Println(bilangan, "Bukan bilangan prima")
	}
}

//Problem 2 - Fast exponentiation
func exponential2(base, pangkat uint64) uint64 {
var result uint64 = 1
	if pangkat == 0 {
		return 0
	} else if pangkat == 1 {
		return 1
	} else {
		for pangkat > 0 {
			if pangkat%2==1{
				result *= base
			}
			pangkat /= 2
			base *= base
		}
	}
	return result
}
func cariPangkat2() {
	fmt.Println(exponential2(2, 3))  // 8

	fmt.Println(exponential2(7, 2))  // 125

	fmt.Println(exponential2(10, 5)) // 100

	fmt.Println(exponential2(17, 6))  // 32

	fmt.Println(exponential2(5, 3))  // 343
	var base uint64
	var pangkat uint64
	fmt.Println("Masukkan base dan pangkat secara berurutan dengan spasi")
	fmt.Scanf("%d %d", &base, &pangkat)
	fmt.Println("hasilnya adalah: ", exponential2(base, pangkat))
}

//Problem 3 - array merge
func ArrayMerge(arrayA, arrayB []string) []string {
	arrayA = append(arrayA, arrayB...)  
	for i, value := range arrayA{

		for j:= i + 1 ; j < len(arrayA); j++ {
			//fmt.Println("ini array pada loop 2 :", arrayA)
			if value == arrayA[j] {
				arrayA[i] = arrayA[len(arrayA)-1] // ganti index i dengan last element.
				//fmt.Println("ini copy last element :", arrayA)
				arrayA = arrayA[:len(arrayA)-1]   // potong slice.
			}
		}
	}
return arrayA
}

func gabungArray(){
	// Test cases
	fmt.Println("Hasil Test cases:")
	fmt.Println("Merger sebelum difilter:  [king devil jin akuma eddie steve geese]")
	fmt.Println("Merger setelah difilter: ", ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"})) // [“king”, “devil jin”, “akuma”, “eddie”, “steve”, “geese”]
	fmt.Println("================================")
	fmt.Println("Merger sebelum difilter: [sergei jin jin steve bryan]")
	fmt.Println("Merger setelah difilter: ", ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"})) // [“sergei”, “jin”, “steve”, “bryan”]
	fmt.Println("================================")
	fmt.Println("Merger sebelum difilter:  [alisa yoshimitsu devil jin yoshimitsu alisa law] ")
	fmt.Println("Merger setelah difilter: ", ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"})) // [“alisa”, “yoshimitsu”, “devil jin”, “law”]
	fmt.Println("================================")
	fmt.Println("Merger sebelum difilter:  [devil jin sergei]")
	fmt.Println("Merger setelah difilter: ", ArrayMerge([]string{}, []string{"devil jin", "sergei"})) // [“devil jin”, “sergei”]
	fmt.Println("================================")
	fmt.Println("Merger sebelum difilter: [hwoarang]")
	fmt.Println("Merger setelah difilter: ", ArrayMerge([]string{"hwoarang"}, []string{})) // [“hwoarang”]
	fmt.Println("================================")
	fmt.Println("Merger sebelum difilter: []")
	fmt.Println("Merger setelah difilter: ", ArrayMerge([]string{}, []string{}))  // []
	fmt.Println("================================")
}

//Problem 4 - Angka Muncul sekali
func munculSekali(angka string) []int {
	array := []rune(angka)
	var arraySekali []int
	var unik bool
	//fmt.Println(string(array))
	for i, value := range array {
		unik = true
		for j:= 0 ; j < len(array); j++ {
			value2 := array[j]
			if value == value2 && i != j{
				unik = false
				break
			}
		}
		if unik == true {
			i, _ := strconv.Atoi(string(value))
			arraySekali = append(arraySekali, i)
		}

	}

return arraySekali

}

func runMunculSekali()  {
	fmt.Println("Hasil test cases:")
	fmt.Println("Input: 1234123")
	fmt.Println("Output: ", munculSekali("1234123"))    // [4]
	fmt.Println("================================")
fmt.Println("Input: ")
	fmt.Println("Output: 76523752", munculSekali("76523752"))   // [6, 3]
	fmt.Println("================================")
	fmt.Println("Input: 12345")
	fmt.Println("Output: ", munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println("================================")
	fmt.Println("Input: 1122334455")
	fmt.Println("Output: ", munculSekali("1122334455")) // []
	fmt.Println("================================")
	fmt.Println("Input: 0872504")
	fmt.Println("Output: ", munculSekali("0872504"))    // [8 7 2 5 4]
	fmt.Println("================================")
	fmt.Println("coba masukkan angka: ")
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println("Angka yang unik adalah:")
	fmt.Println(munculSekali(input))

}

//Problem 5
func PairSum(arr []int, target int) []int {

	//cara 1 dengan o(n^2)
//	var arrSum []int
//	for i, value := range arr {
//		for j:= i + 1 ; j < len(arr); j++ {
//			value2 := arr[j]
//			if (value + value2) == target {
//				arrSum = append(arrSum, i, j)
//			}
//		}
//	}
//return arrSum

//	cara 2 dengan 0(n)
	index1 := 0
	index2 := len(arr) - 1
	sum := arr[index1] + arr[index2]

	for sum != target {
		if sum > target {
			index2 -= 1
		} else {
			index1 += 1
		}
		sum = arr[index1] + arr[index2]
	}

	return []int{index1, index2}

}
func runPairSum (){
	fmt.Println("Hasil test cases: ")
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}

func listProgramTaskH3()  {
	fmt.Println("=======================")
	fmt.Println("List Program Task Hari ke 3 Backend")
	fmt.Println("1. Mencari Bilangan Prima")
	fmt.Println("2. Fast Exponentiation")
	fmt.Println("3. Merge Array")
	fmt.Println("4. Filter Angka Unik")
	fmt.Println("5. Pair with Target Sum")
	fmt.Println("=======================")
	fmt.Println("Pilih Program yang akan dijakankan:")
}
func pilihProgramTaskH3()  {
	var pilih int
	fmt.Scanf("%d", &pilih)
	switch pilih{
	case 1:
		mencariBilanganPrima2()
	case 2:
		cariPangkat2()
	case 3:
		gabungArray()
	case 4:
		runMunculSekali()
	case 5:
		runPairSum()
	default:
		fmt.Println("Input tidak dikenali")
	}
}

func main() {
listProgramTaskH3()
pilihProgramTaskH3()
}
