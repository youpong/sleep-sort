// 07-Sep-22 satlinuxtube(武内 覚)
// THE FIRST CODE 寝てたら終わるスリープソート(https://youtu.be/PSeIEBPnq-E)

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func SleepSort(a []int) (error, []int) {
	r := make([]int, len(a))
	var wg sync.WaitGroup

	var current int
	var lock sync.Mutex
	for _, v := range a {
		if v < 0 {
			return fmt.Errorf("Must be non-negative: %d", v), nil
		}
		wg.Add(1)
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Second)
			lock.Lock()
			r[current] = v
			current++
			lock.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()

	return nil, r
}

func sortAndPrint(a []int) {
	err, res := SleepSort(a)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(res)
}

func main() {
	// sortAndPrint([]int{-1})
	sortAndPrint([]int{1})
	sortAndPrint([]int{2, 1})
	sortAndPrint([]int{5, 4, 2, 3, 1})
	sortAndPrint([]int{2, 1, 1, 2})
}
