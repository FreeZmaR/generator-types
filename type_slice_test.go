package types

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSliceUnmarshalSuite struct {
	suite.Suite
}

func TestSlice_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(TestSliceUnmarshalSuite))
}

func (s *TestSliceUnmarshalSuite) TestUnmarshalSliceString() {
	var (
		inputData    = []byte(`["a", "b", "c"]`)
		assertResult = Slice[string]{
			isProvided: true,
			value:      []string{"a", "b", "c"},
		}
	)

	var result Slice[string]

	err := result.UnmarshalJSON(inputData)
	s.NoError(err)
	s.Equal(assertResult, result)
}

func (s *TestSliceUnmarshalSuite) TestUnmarshalSliceStruct() {
	type User struct {
		Name Type[string] `type:"string"`
		Age  Type[int]    `type:"age"`
	}

	var (
		inputData    = []byte(`[{"name": "a", "age": 1}, {"name": "b", "age": 2}]`)
		assertResult = Slice[User]{
			isProvided: true,
			value: []User{
				{
					Name: NewType(WithIsProvided[string](true), WithValue("a")),
					Age:  NewType(WithIsProvided[int](true), WithValue(1)),
				},
				{
					Name: NewType(WithIsProvided[string](true), WithValue("b")),
					Age:  NewType(WithIsProvided[int](true), WithValue(2)),
				},
			},
		}
	)

	var result Slice[User]

	err := result.UnmarshalJSON(inputData)
	s.NoError(err)
	s.Equal(assertResult, result)
}

func (s *TestSliceUnmarshalSuite) TestUnmarshalSliceStructWithPrepareFN() {
	type User struct {
		Name Type[string] `type:"string"`
		Age  Type[int]    `type:"age"`
	}

	var (
		prepareFN = func(u *Slice[User]) {
			u.SetTag("test")
			vals := u.Value()

			for i := range vals {
				vals[i].Name.SetTag("name")
				vals[i].Age.SetTag("age")
			}
		}

		inputData    = []byte(`[{"name": "a", "age": 1}, {"name": "b", "age": 2}]`)
		assertResult = Slice[User]{
			isProvided: true,
			tag:        "test",
			value: []User{
				{
					Name: NewType(
						WithIsProvided[string](true),
						WithValue("a"),
						WithTag[string]("name"),
					),
					Age: NewType(
						WithIsProvided[int](true),
						WithValue(1),
						WithTag[int]("age"),
					),
				},
				{
					Name: NewType(
						WithIsProvided[string](true),
						WithValue("b"),
						WithTag[string]("name"),
					),
					Age: NewType(
						WithIsProvided[int](true),
						WithValue(2),
						WithTag[int]("age"),
					),
				},
			},
		}
	)

	var result Slice[User]
	result.prepareFN = prepareFN

	err := result.UnmarshalJSON(inputData)
	s.NoError(err)

	// to pass test, with func objects are difference
	result.prepareFN = nil
	s.Equal(assertResult, result)
}
