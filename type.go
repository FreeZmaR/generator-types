package core

type Types interface {
	int | float64 | string | bool
}

type Type[T Types] struct {
	isProvided bool
	value      T
	tag        string
}

func (t Type[T]) IsProvided() bool {
	return t.isProvided
}

func (t Type[T]) Value() T {
	return t.value
}

func (t Type[T]) Tag() string {
	return t.tag
}

func (t Type[T]) Validate(rules ...Rule[T]) error {
	for i := range rules {
		if err := rules[i].Check(t.value); err != nil {
			return err
		}
	}

	return nil
}
