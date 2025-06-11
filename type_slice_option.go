package types

type SliceOption[T any] func(t *Slice[T])

func WithSliceIsProvided[T any](val bool) SliceOption[T] {
	return func(t *Slice[T]) { t.isProvided = val }
}

func WithSliceValue[T any](val []T) SliceOption[T] {
	return func(t *Slice[T]) { t.value = val }
}

func WithSliceTag[T any](tag string) SliceOption[T] {
	return func(t *Slice[T]) { t.tag = tag }
}

func WithSlicePrepareFN[T any](fn SlicePrepareFN[T]) SliceOption[T] {
	return func(t *Slice[T]) { t.prepareFN = fn }
}

func WithSliceRules[T any](rules ...SliceRule[T]) SliceOption[T] {
	return func(t *Slice[T]) { t.rules = rules }
}
