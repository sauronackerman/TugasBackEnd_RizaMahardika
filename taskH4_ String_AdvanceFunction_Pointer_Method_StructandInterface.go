package main

import (
	"fmt"
	"strings"
)

//Problem 1 - Compare String
func Compare(a, b string) string {
	if strings.Contains(a, b) {
		return b
	} else if strings.Contains(b, a){
		return a
	} else {
		return "Salah"
	}


}
func testCompare()  {
	fmt.Println("Hasil test: ")
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA

	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG

	fmt.Println(Compare("KI", "KIJANG"))      // KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}

//Problem 2 - Caesar Cipher
func caesar(offset int, input string) string  {
	var inputRune = []rune(input)
	var newRune []rune
	if offset > 26 {
		offset = offset % 26
	}
	for _, value := range inputRune {
		//fmt.Println(value)
		newValue := int(value) + offset
		//fmt.Println(newValue)
		if newValue > 122 {
			//fmt.Println(newValue)
			newValue = (newValue % 122) + 96
			//fmt.Println(newValue)
		}
		newRune = append(newRune, rune(newValue))
	}
	return string(newRune)

//	cara 2 (hasil browsing)
//	transform := func(r rune) rune {
//		if offset > 26 {
//			offset = offset % 26
//		}
//		s := int(r) + offset
//		if s > 'z' {
//			return rune(s - 26)
//		} else if s < 'a' {
//			return rune(s + 26)
//		}
//		return rune(s)
//	}
//	return strings.Map(transform, input)
//
}
func testCaesar()  {
	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy"))  // kvdobbkkmknowi
	//fmt.Println(caesar(10, "at"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}

//Problem 3 -  Swap Two Number Using Pointer
func swap(num1, num2 *int)  {
	temp1 := *num1
	//temp2 := *num2
	//*num1 = temp2
	*num1 = *num2
	*num2 = temp1
}
func testSwap()  {
	a := 10
	b := 20
	fmt.Println("Sebelum swap: ", a, b)
	swap(&a, &b)
	fmt.Println("setelah swap: ", a, b)
}

//Problem 4 - Min and Max Using Pointer
func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]
	for _, value := range numbers {
			if *value < min {
				min = *value
			}
			if *value > max {
				max = *value
			}
		}
	return min, max

}

func testGetMinMax()  {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Masukkan nilai satu persatu:")
	fmt.Scan(&a1)

	fmt.Scan(&a2)

	fmt.Scan(&a3)

	fmt.Scan(&a4)

	fmt.Scan(&a5)

	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai min ", min)

	fmt.Println("Nilai max ", max)
}

//Problem 5 - Student Score
type Student struct {
	name  []string
	score []int
}
func (s Student) Average() float64 {
	sum := make(map[string]int)
	for _, value := range s.score {
		sum[""] += value
	}

	return float64(sum[""] / len(s.score))


}

func (s Student) Min() (min int, name string) {
	minMap := map[string]int{
		"": s.score[0],
	}
	var iMin int
	for i, value := range s.score {
		if minMap[""] > value {
			minMap[""] = value
			iMin = i
		}
	}
	return minMap[""], s.name[iMin]
}



func (s Student) Max() (max int, name string) {
	maxMap := map[string]int{
		"": s.score[0],
	}

	var iMax int

	for i, value := range s.score {
		if maxMap[""] < value {
			maxMap[""] = value
			iMax = i
		}
	}

	return maxMap[""], s.name[iMax]


}

func testStudentScore(){
	var a = Student{}



	for i := 0; i < 6; i++ {

		var name string

		fmt.Print("Input " + string(i) + " Student’s Name : ")

		fmt.Scan(&name)

		a.name = append(a.name, name)

		var score int

		fmt.Print("Input " + name + " Score : ")

		fmt.Scan(&score)

		a.score = append(a.score, score)

	}



	fmt.Println("\n\nAvarage Score Students is", a.Average())

	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}

//Problem 6 - Substitution Chiper
type student struct {
	name string
	nameEncode string
	score int
}
type Chiper interface {
	Encode() string
	Decode() string
}
func (s *student) Encode() string {

	var inputRune = []rune(s.name)
	var newRune []rune
	for _, value := range inputRune{
		newValue := int(value) + 5
		if newValue > 122 {
			//fmt.Println(newValue)
			newValue = (newValue % 122) + 96
			//fmt.Println(newValue)
		}
		newRune = append(newRune, rune(newValue))
	}
	return string(newRune)
}
func (s *student) Decode() string {

	Decoding := func(r rune) rune {
		s := int(r) - 5
		if s > 'z' {
			return rune(s - 26)
		} else if s < 'a' {
			return rune(s + 26)
		}
		return rune(s)
	}
	return strings.Map(Decoding, s.nameEncode)


}

func testStudentChiper(){
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student’s Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode() + "\n")
	} else if menu == 2 {
		fmt.Print("\nInput Student’s Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode() + "\n")
	} else {
		fmt.Println("Wrong input name menu")
	}

}

func listProgramTaskH4()  {
	fmt.Println("=======================")
	fmt.Println("List Program Task Hari ke 4 Backend - Riza Mahardika")
	fmt.Println("1. Compare String")
	fmt.Println("2. Caesar Chiper")
	fmt.Println("3. Swap Two Number Using Pointer")
	fmt.Println("4. Min and Max Using Pointer")
	fmt.Println("5. Student Score")
	fmt.Println("6. Substitution Chiper")
	fmt.Println("=======================")
	fmt.Println("Pilih Program yang akan dijakankan:")
}
func pilihProgramTaskH4()  {
	var pilih int
	fmt.Scanf("%d", &pilih)
	switch pilih{
	case 1:
		testCompare()
	case 2:
		testCaesar()
	case 3:
		testSwap()
	case 4:
		testGetMinMax()
	case 5:
		testStudentScore()
	case 6:
		testStudentChiper()
	default:
		fmt.Println("Input tidak dikenali")
	}
}

func main() {
listProgramTaskH4()
pilihProgramTaskH4()
}

