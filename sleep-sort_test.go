package main

import "testing"

func TestTail2Head(t *testing.T) {
	t.Parallel()
	sample := make([]int, (128+16)*1024) // Intel Core2 Duo T9400@2.53GHz
	for i, _ := range sample {
		sample[i] = 2
	}
	sample[len(sample)-1] = 1

	err, r := SleepSort(sample)
	if err != nil {
		t.Errorf("%v", err)
	}
	if i, ok := checkOrder(r); !ok {
		t.Errorf("wrong value(%d) at %d of %d", r[i], i, len(r)-1)
	}
}

type testSet struct {
	name   string
	want   []int
	sample []int
}

func TestNegative(t *testing.T) {
	t.Parallel()
	ts := testSet{
		name:   "negative",
		want:   []int{},
		sample: []int{-1},
	}

	err, _ := SleepSort(ts.sample)
	if err == nil {
		t.Errorf("Sould return error")
	}
}

func TestNormal(t *testing.T) {
	t.Parallel()
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
			name:   "duplicated values",
			want:   []int{1, 1, 2, 2},
			sample: []int{2, 1, 1, 2},
		},
	}

	for _, ts := range tests {
		ts := ts
		t.Run(ts.name, func(t *testing.T) {
			t.Parallel()
			err, got := SleepSort(ts.sample)
			if err != nil {
				t.Errorf("%v", err)
			}
			if !sameArray(got, ts.want) {
				t.Errorf("%v) got = %v, want %v", ts.name, got, ts.want)
			}
		})
	}
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

func checkOrder(a []int) (int, bool) {
	if len(a) == 0 {
		return -1, true
	}

	prev := a[0]
	for i, v := range a {
		if prev > v {
			return i, false
		}
		prev = v
	}
	return -1, true
}
