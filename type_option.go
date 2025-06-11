package types

type TypeOption[T any] func(t *Type[T])

func WithIsProvided[T any](val bool) TypeOption[T] {
	return func(t *Type[T]) { t.isProvided = val }
}

func WithValue[T any](val T) TypeOption[T] {
	return func(t *Type[T]) { t.value = val }
}

func WithTag[T any](tag string) TypeOption[T] {
	return func(t *Type[T]) { t.tag = tag }
}

func WithCastFn[T any](fn TypeCastFn[T]) TypeOption[T] {
	return func(t *Type[T]) { t.castFN = fn }
}

func WithRules[T any](rules ...TypeRule[T]) TypeOption[T] {
	return func(t *Type[T]) {
		t.rules = rules
	}
}
