// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package ds

type StringSet map[string]struct{}

func NewStringSet(list []string) (s StringSet) {
	s = StringSet{}
	for _, v := range list {
		s[v] = struct{}{}
	}
	return
}
func (s StringSet) Contains(t string) (ok bool) {
	_, ok = s[t]
	return
}
func (s StringSet) Add(t string) {
	s[t] = struct{}{}
}
func (s StringSet) Del(t string) {
	delete(s, t)
}

func (s StringSet) Union(s2 StringSet) (result StringSet) {
	result = StringSet{}
	for v := range s {
		result[v] = struct{}{}
	}
	for v := range s2 {
		result[v] = struct{}{}
	}
	return
}

func (s StringSet) Intersection(s2 StringSet) (result StringSet) {
	result = StringSet{}
	for v := range s {
		if s2.Contains(v) {
			result[v] = struct{}{}
		}
	}
	return
}

func (s StringSet) Include(s2 StringSet) bool {

	for v := range s2 {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

func (s StringSet) ComplementarySet(s2 StringSet) (result StringSet) {
	result = StringSet{}
	for v := range s {
		if !s2.Contains(v) {
			result[v] = struct{}{}
		}
	}
	return
}
func (s StringSet) ToSlice() []string {
	list := []string{}
	for v := range s {
		list = append(list, v)
	}
	return list
}

type IntSet map[int]struct{}

func NewIntSet(list []int) (s IntSet) {
	s = IntSet{}
	for _, v := range list {
		s[v] = struct{}{}
	}
	return
}
func (s IntSet) Contains(t int) (ok bool) {
	_, ok = s[t]
	return
}
func (s IntSet) Add(t int) {
	s[t] = struct{}{}
}
func (s IntSet) Del(t int) {
	delete(s, t)
}

func (s IntSet) Union(s2 IntSet) (result IntSet) {
	result = IntSet{}
	for v := range s {
		result[v] = struct{}{}
	}
	for v := range s2 {
		result[v] = struct{}{}
	}
	return
}

func (s IntSet) Intersection(s2 IntSet) (result IntSet) {
	result = IntSet{}
	for v := range s {
		if s2.Contains(v) {
			result[v] = struct{}{}
		}
	}
	return
}

func (s IntSet) Include(s2 IntSet) bool {

	for v := range s2 {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

func (s IntSet) ComplementarySet(s2 IntSet) (result IntSet) {
	result = IntSet{}
	for v := range s {
		if !s2.Contains(v) {
			result[v] = struct{}{}
		}
	}
	return
}
func (s IntSet) ToSlice() []int {
	list := []int{}
	for v := range s {
		list = append(list, v)
	}
	return list
}
