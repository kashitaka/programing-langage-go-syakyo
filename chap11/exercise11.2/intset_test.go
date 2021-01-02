package intset

import (
	"testing"
)

func TestHasandAdd(t *testing.T) {
	tests := []struct {
		keys []uint64
		add  int
	}{
		{[]uint64{1, 2, 3, 4}, 9},
	}
	for _, test := range tests {
		var m = map[int]bool{}
		var is IntSet
		for k := range test.keys {
			m[k] = true
			is.Add(k)
		}
		for k := range test.keys {
			_, ok := m[k]
			if ok != is.Has(k) {
				t.Errorf("%v Has(%d) dose not much to %v[%d]", is, k, m, k)
			}
		}

		is.Add(test.add)
		m[int(test.add)] = true
		_, ok := m[test.add]
		if ok != is.Has(test.add) {
			t.Errorf("Add value %d was failed to add whether %v or %v", test.add, m, is)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		keys1 []int
		keys2 []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{3, 4, 5},
		},
	}
	for _, test := range tests {
		var m = map[int]bool{}
		var is1 IntSet
		var is2 IntSet
		for _, v := range test.keys1 {
			m[v] = true
			is1.Add(v)
		}
		for _, v := range test.keys2 {
			m[v] = true
			is2.Add(v)
		}
		is1.UnionWith(&is2)
		for k := range m {
			if !is1.Has(k) {
				t.Errorf("factor of %v and %v not matched with %d", is1, m, k)
			}
		}
	}
}
