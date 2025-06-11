package types

type StructOption[T any] func(s *Struct[T])

func WithStructIsProvided[T any](val bool) StructOption[T] {
	return func(s *Struct[T]) {
		s.isProvided = val
	}
}

func WithStructTag[T any](tag string) StructOption[T] {
	return func(s *Struct[T]) {
		s.tag = tag
	}
}

func WithStructPrepareFN[T any](fn PrepareFN[T]) StructOption[T] {
	return func(s *Struct[T]) {
		s.prepareFN = fn
	}
}

func WithStructRules[T any](rules ...StructRule[T]) StructOption[T] {
	return func(s *Struct[T]) {
		s.rules = rules
	}
}
