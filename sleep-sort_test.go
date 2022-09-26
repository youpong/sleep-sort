package sleep_sort

import (
	"sync"
	"testing"
)

func testOrder(a []int) bool {
	if len(a) == 0 {
		return true
	}

	prev := a[0]
	for _, v := range a {
		if prev > v {
			return false
		}
	}
	return true
}

func sameArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNormal(t *testing.T) {
	type testSet struct {
		name   string
		want   []int
		sample []int
	}

	tests := []testSet{
		{
			name:   "one element",
			want:   []int{1},
			sample: []int{1},
		}, {
			name:   "two elements",
			want:   []int{1, 2},
			sample: []int{2, 1},
		}, {
			name:   "five elements",
			want:   []int{1, 2, 3, 4, 5},
			sample: []int{5, 4, 2, 3, 1},
		}, {
			name:   "same value",
			want:   []int{1, 1, 2, 2},
			sample: []int{2, 1, 1, 2},
		},
	}

	var wg sync.WaitGroup
	for _, ts := range tests {
		wg.Add(1)
		go func(ts testSet) {
			if got := SleepSort(ts.sample); !sameArray(got, ts.want) {
				t.Errorf("%v) got = %v, want %v", ts.name, got, ts.want)
			}
			wg.Done()
		}(ts)
	}
	wg.Wait()
}

func TestScale(t *testing.T) {
	sample := make([]int, 256*1024)
	for i, _ := range sample {
		sample[i] = 2
	}
	sample[len(sample)-1] = 1

	if r := SleepSort(sample); !testOrder(r) {
		for i, _ := range r {
			if r[i] == 1 {
				t.Errorf("r[%v] = 1", i)
			}
		}
	}
}
