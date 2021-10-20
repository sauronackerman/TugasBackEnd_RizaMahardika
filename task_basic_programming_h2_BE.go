package main

import (
	"fmt"
	"log"
	"strconv"
"github.com/manifoldco/promptui"

)

//Problem 1 luas permukaan tabung
// rumus 2 π r ( r + t )
const π float64 = 3.14

func luasPermukaanTabung(t float64, r float64) float64 {
	return 2 * π * r * (r + t)
}

func hitungLuasPermukaanTabung()  {
	var tinggi float64
	var jariJari float64
	fmt.Println("Masukan tinggi (cm):")
	fmt.Scanf("%f", &tinggi)
	fmt.Println("Masukan jari-jari (cm):")
	fmt.Scanf("%f", &jariJari)
	fmt.Println("luas permukaan tabung tersebut adalah: ", luasPermukaanTabung(tinggi, jariJari), "cm^2")
}

//problem 2 -Grade Nilai

func gradeNilai()  {
	var nilai int
	fmt.Println("Masukkan nilai: ")
	fmt.Scanf("%d", &nilai)
	switch {
	case nilai >= 80 && nilai <= 100:
fmt.Println("Nilai A")
	case nilai >= 65 && nilai <= 79:
		fmt.Println("Nilai B")
	case nilai >= 50 && nilai <= 64	:
		fmt.Println("Nilai C")
	case nilai >= 35 && nilai <= 49:
		fmt.Println("Nilai D")
	case nilai >= 0 && nilai <= 34:
		fmt.Println("Nilai E")
	default:
		fmt.Println("Nilai Invalid")
	}
}

//Problem 3 Faktor Bilangan
func mencariFaktorBilangan() {
	var input int
	fmt.Println("Masukkan Bilangan")
	fmt.Scanf("%d", &input)
	fmt.Println("Faktor bilangan " + strconv.Itoa(input) + " adalah" )
	for i := 1; i <= input; i++ {
		if input % i == 0 {
			fmt.Println(i)
		}
	}
}

//Problem 4 Bilangan Prima
func mencariBilanganPrima() {
	var input int
	fmt.Println("Masukkan bilangan integer")
	fmt.Scanf("%d", &input)
	bilangan := strconv.Itoa(input)
	if input <= 2 {
		switch {
		case input == 2:
			fmt.Println(bilangan , "Adalah bilangan prima")
		default:
			fmt.Println(bilangan, "Bukan bilangan prima")
		}
	} else {
		for i := 2; i < input; i++ {
		if input % i == 0 {
			fmt.Println(bilangan, "Bukan bilangan prima")
			break
		} else {
			if i < input - 1 {
				continue
			}
			fmt.Println(bilangan, "Adalah bilangan prima")
			break
		}
		}
	}
}

//Problem 5 - Palindrome

func palindrome(input string) bool {

	for i := 0; i< len(input)/2; i++ {
		if input[i] != input[len(input)-i-1]{
			return false
		}
	}
	return true

}
func cariPalindrome() {
	fmt.Println(palindrome("civic")) // true

	fmt.Println(palindrome("katak")) // true

	fmt.Println(palindrome("kasur rusak")) // true

	fmt.Println(palindrome("mistar")) // false

	fmt.Println(palindrome("lion")) // false
	var input string
	fmt.Println("Masukkan kata")
	fmt.Scanf("%q", &input)
	hasil := palindrome(input)

	if hasil == true {
		fmt.Println(input, "Adalah Palindrome")

	} else {
	 	fmt.Println(input, "Bukan Palindrom")
	}
}

//Problem 6 - Eksponentiation
func exponential(base, pangkat int) int {
	 result := 1
	for i := 1; i <= pangkat; i++ {
		result *= base
	}
	return result
}
func cariPangkat() {
	fmt.Println(exponential(2, 3))  // 8

	fmt.Println(exponential(5, 3))  // 125

	fmt.Println(exponential(10, 2)) // 100

	fmt.Println(exponential(2, 5))  // 32

	fmt.Println(exponential(7, 3))  // 343
	var base int
	var pangkat int
	fmt.Println("Masukkan base dan pangkat secara berurutan")
fmt.Scanf("%d %d", &base, &pangkat)
fmt.Println("hasilnya adalah: ", exponential(base, pangkat))
}

//Problem 7 - Play With Asterisk
func playWithAsterik(n int) {
	// your code here
	for i := n - 1; i >= 0; i-- {

		for j := 0; j < i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < n-i; k++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}

}

func buatAsterik()  {
	playWithAsterik(5)
	/*

	       *

	      * *

	     * * *

	    * * * *

	   * * * * *

	*/
	fmt.Println("Masukkan angka untuk menyusun jumlah baris segitiga: ")
	var input int
	fmt.Scanf("%d", &input)
	playWithAsterik(input)
}

//Problem 8 – Cetak Tabel Perkalian
func tabelPerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d \t", i*j)
		}
		fmt.Println("\n")
	}

}

func cetakTabelPerkalian()  {

	fmt.Println("Masukkan jumlah maksimal tabel perkalian: ")
	var input int
	fmt.Scanf("%d", &input)
	tabelPerkalian(input)
}


//tambahan : fungsi untuk menjalankan program dalam CLI
func programBasic()  {
	var pilih int
	fmt.Scanf("%d", &pilih)
	switch pilih{
	case 1:
		hitungLuasPermukaanTabung()
	case 2:
		gradeNilai()
	case 3:
		mencariFaktorBilangan()
	case 4:
		mencariBilanganPrima()
	case 5:
		cariPalindrome()
	case 6:
		cariPangkat()
	case 7:
		buatAsterik()
	case 8:
		cetakTabelPerkalian()
	default:
		fmt.Println("Input tidak dikenali")
	}

}
func listProgram()  {
	fmt.Println("1. Menghitung Luas Permukaan Tabung")
	fmt.Println("2. Memasukan Grade Nilai")
	fmt.Println("3. Mencari faktor bilangan")
	fmt.Println("4. Menentukan apakah suatu angka adalah bilangan prima")
	fmt.Println("5. Menentukan apakah suatu kata termasuk paliandrom")
	fmt.Println("6. Mencari hasil eksponensial")
	fmt.Println("7. Membuat pola segitiga")
	fmt.Println("8. Membuat tabel perkalian")
	fmt.Println("Pilih Program yang akan dijalankan berdasarkan nomor yang tersedia:")
}

func yesNo() bool {
	fmt.Println("Apakah kamu ingin mencoba yang lain?")
	prompt := promptui.Select{
		Label: "Select[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}


func main() {
//mencariFaktorBilangan()
//	mencariBilanganPrima()
//	cariPalindrome()
//	cariPangkat()
//	buatAsterik()
//cetakTabelPerkalian()

	listProgram()
	programBasic()
	yesNo()
	switch yesNo(){
	case true :
		listProgram()
		programBasic()
	default:
		fmt.Println("Sampai jumpa lagi")

	}

}
