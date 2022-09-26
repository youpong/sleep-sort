package sleep_sort

import (
	//"sleep_sort"
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

func TestFoo(t *testing.T) {

	tests := []struct {
		want   []int
		sample []int
	}{
		{
			want:   []int{1},
			sample: []int{1},
		}, {
			want:   []int{1, 2},
			sample: []int{2, 1},
		}, {
			want:   []int{1, 2, 3, 4, 5},
			sample: []int{5, 4, 2, 3, 1},
		}, {
			want:   []int{1, 1, 2, 2},
			sample: []int{2, 1, 1, 2},
		},
	}

	for i, tt := range tests {
		if got := SleepSort(tt.sample); !sameArray(got, tt.want) {
			t.Errorf("%v) got = %v, want %v", i, got, tt.want)
		}
	}
}

//sleep_sort.SleepSort(nil)
