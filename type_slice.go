package types

import "github.com/goccy/go-json"

type Slice[T any] struct {
	isProvided bool
	value      []T
	tag        string
	prepareFN  SlicePrepareFN[T]
}

type SlicePrepareFN[T any] func(val *Slice[T])

func NewSlice[T any](isProvided bool, val []T, tag string, prepareFN SlicePrepareFN[T]) Slice[T] {
	return Slice[T]{
		isProvided: isProvided,
		value:      val,
		tag:        tag,
		prepareFN:  prepareFN,
	}
}

func (s *Slice[T]) IsProvided() bool {
	return s.isProvided
}

func (s *Slice[T]) Value() []T {
	return s.value
}

func (s *Slice[T]) Tag() string {
	return s.tag
}

func (s *Slice[T]) SetTag(tag string) {
	s.tag = tag
}

func (s *Slice[T]) Len() int {
	return len(s.value)
}

func (s *Slice[T]) UnmarshalJSON(b []byte) error {
	s.isProvided = true

	if err := json.Unmarshal(b, &s.value); err != nil {
		return err
	}

	if s.prepareFN != nil {
		s.prepareFN(s)
	}

	return nil
}
