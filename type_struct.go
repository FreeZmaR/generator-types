package main

import "github.com/goccy/go-json"

type Struct[T any] struct {
	isProvided bool
	value      T
	tag        string
	prepareFN  PrepareFN[T]
}

type PrepareFN[T any] func() T

func NewStruct[T any](isProvided bool, tag string, prepareFN PrepareFN[T]) Struct[T] {
	return Struct[T]{
		isProvided: isProvided,
		tag:        tag,
		prepareFN:  prepareFN,
	}
}

func (s *Struct[T]) IsProvided() bool {
	return s.isProvided
}

func (s *Struct[T]) Value() T {
	return s.value
}

func (s *Struct[T]) Tag() string {
	return s.tag
}

func (s *Struct[T]) SetTag(tag string) {
	s.tag = tag
}

func (s *Struct[T]) UnmarshalJSON(b []byte) error {
	s.isProvided = true

	if s.prepareFN == nil {
		return json.Unmarshal(b, &s.value)
	}

	dummy := s.prepareFN()

	if err := json.Unmarshal(b, &dummy); err != nil {
		return err
	}

	s.value = dummy

	return nil
}
