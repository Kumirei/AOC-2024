package util

import "fmt"

type Set[T comparable] struct {
	List map[T]struct{} //empty structs occupy 0 memory
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.List[v]
	return ok
}

func (s *Set[T]) Add(v T) {
	s.List[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.List, v)
}

func (s *Set[T]) Clear() {
	s.List = make(map[T]struct{})
}

func (s *Set[T]) Size() int {
	return len(s.List)
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.List = make(map[T]struct{})
	return s
}

//optional functionalities

// AddMulti Add multiple values in the set
func (s *Set[T]) AddMulti(list ...T) {
	for _, v := range list {
		s.Add(v)
	}
}

type FilterFunc[T comparable] func(v T) bool

// Filter returns a subset, that contains only the values that satisfies the given predicate P
func (s *Set[T]) Filter(P FilterFunc[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.List {
		if P(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

func (s *Set[T]) ToList() (list []T) {
	for v := range s.List {
		list = append(list, v)
	}
	return list
}

func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.List {
		res.Add(v)
	}

	for v := range s2.List {
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.List {
		if s2.Has(v) == false {
			continue
		}
		res.Add(v)
	}
	return res
}

// Difference returns the subset from s, that doesn't exists in s2 (param)
func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	res := NewSet[T]()
	for v := range s.List {
		if s2.Has(v) {
			continue
		}
		res.Add(v)
	}
	return res
}

func (s *Set[T]) Combine(s2 *Set[T]) {
	for x := range s2.List {
		s.Add(x)
	}
}

func (s Set[T]) String() string {
	return fmt.Sprint(s.ToList())
}

func (s *Set[T]) Clone() *Set[T] {
	clone := NewSet[T]()
	clone.Combine(s)
	return clone
}
