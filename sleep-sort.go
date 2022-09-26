package sleep_sort

import (
	"fmt"
	"sync"
	"time"
)

func SleepSort(a []int) []int {
	res := make([]int, len(a))
	var wg sync.WaitGroup

	tail := 0
	var lock sync.Mutex
	for _, v := range a {
		wg.Add(1)
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Millisecond)
			lock.Lock()
			res[tail] = v
			tail++
			lock.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()

	return res
}

func sortAndPrint(a []int) {
	res := SleepSort(a)
	fmt.Println(res)
}

func main() {
	sortAndPrint([]int{1})
	sortAndPrint([]int{2, 1})
	sortAndPrint([]int{5, 4, 2, 3, 1})
	sortAndPrint([]int{2, 1, 1, 2})
}
