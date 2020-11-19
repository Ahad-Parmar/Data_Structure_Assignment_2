package main

import "fmt"

type mymap map[int]*myentry

type myentry struct {
	m mymap
	b bool
}

func (mm mymap) get(idx ...int) *myentry {
	if len(index) == 0 {
		return nil
	}
	entry, ok := mm[index[0]]
	if !ok {
		return nil
	} else if len(index) == 1 {
		return entry
	}
	for i := 1; i < len(index); i++ {
		if entry == nil || entry.m == nil {
			return nil
		}
		entry = entry.m[index[i]]
	}
	return entry
}

func (mm mymap) setbool(v bool, index ...int) {
	if len(index) == 0 {
		return
	}
	if mm[index[0]] == nil {
		mm[index[0]] = &myentry{m: make(mymap), b: false}
	} else if mm[index[0]].m == nil {
		mm[index[0]].m = make(mymap)
	}
	if len(index) == 1 {
		mm[index[0]].b = v
		return
	}
	entry := mm[index[0]]
	for i := 1; i < len(index); i++ {
		if entry.m == nil {
			entry.m = make(mymap)
			entry.m[index[i]] = &myentry{m: make(mymap), b: false}
		} else if entry.m[index[i]] == nil {
			entry.m[index[i]] = &myentry{m: make(mymap), b: false}
		}
		entry = entry.m[index[i]]
	}
	entry.b = v
}

func (m mymap) getbool(index ...int) bool {
	if val := m.get(index...); val != nil {
		return val.b
	}
	return false
}

func (m mymap) getmap(index ...int) mymap {
	if val := m.get(index...); val != nil {
		return val.m
	}
	return nil
}

func main() {
	m := make(mymap)
	m.setbool(true, 0, 2, 4)
	fmt.Println("Should be false:", m.getbool(0, 2))
	fmt.Println("Should be true:", m.getbool(0, 2, 4))
	m.setbool(true, 0, 2, 4, 6)
	fmt.Println("Should be true:", m.getbool(0, 2, 4, 6))
	m.setbool(false, 0, 2, 4)
	fmt.Println("Should be false:", m.getbool(0, 2, 4))
	fmt.Println("Should contain key 4 only:", m.getmap(0, 2))
	m.setbool(true, 0, 2, 4)
	m.setbool(true, 0, 2, 8)
	fmt.Println("Should contain keys 2, 4, and 8:", m.getmap(0, 2))
	q := m.getmap(0, 2)
	fmt.Println("should be true:", q.getbool(8))
}