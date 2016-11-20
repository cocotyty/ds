//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Type=string,int"
package ds

import "github.com/cheekybits/genny/generic"

type Type generic.Type

type TypeSet map[Type]struct{}

func NewTypeSet(list []Type) (s TypeSet) {
	s = TypeSet{}
	for _, v := range list {
		s[v] = struct{}{}
	}
	return
}
func (s TypeSet) Contains(t Type) (ok bool) {
	_, ok = s[t]
	return
}
func (s TypeSet) Add(t Type) {
	s[t] = struct{}{}
}
func (s TypeSet) Del(t Type) {
	delete(s, t)
}

func (s TypeSet) Union(s2 TypeSet) (result TypeSet) {
	result = TypeSet{}
	for v := range s {
		result[v] = struct{}{}
	}
	for v := range s2 {
		result[v] = struct{}{}
	}
	return
}

func (s TypeSet) Intersection(s2 TypeSet) (result TypeSet) {
	result = TypeSet{}
	for v := range s {
		if s2.Contains(v) {
			result[v] = struct{}{}
		}
	}
	return
}

func (s TypeSet) Include(s2 TypeSet) bool {

	for v := range s2 {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}

func (s TypeSet) ComplementarySet(s2 TypeSet) (result TypeSet) {
	result = TypeSet{}
	for v := range s {
		if !s2.Contains(v) {
			result[v] = struct{}{}
		}
	}
	return
}
func (s TypeSet) ToSlice() []Type {
	list := []Type{}
	for v := range s {
		list = append(list, v)
	}
	return list
}
