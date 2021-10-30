
package main

import (
"fmt"
"strings"
"sync"
)

type freqMap struct {
	hasil map[string]int
	mutex     sync.Mutex
}

func (s *freqMap) letterFrequency(huruf string, jumlahHuruf chan int, group *sync.WaitGroup) {
	s.mutex.Lock()
	s.hasil[huruf]++
	jumlahHuruf <- s.hasil[huruf]
	fmt.Println(huruf, ":", <-jumlahHuruf)
	group.Done()
	s.mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mapHasil = freqMap{hasil: make(map[string]int)}
	lorem := "lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	strFix := strings.ReplaceAll(lorem, " ", "")
	size := make(chan int, len(strFix))
	wg.Add(len(strFix))

	for _, value := range strFix {
	go mapHasil.letterFrequency(string(value), size, &wg)
	}
	wg.Wait()
	delete(mapHasil.hasil, ",")
	fmt.Println("Hasil:",mapHasil.hasil)
}

