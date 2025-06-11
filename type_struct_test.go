package types

import (
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestStructUnmarshalSuite struct {
	suite.Suite
}

func TestStruct_UnmarshalJSON(t *testing.T) {
	suite.Run(t, new(TestStructUnmarshalSuite))
}

func (s *TestStructUnmarshalSuite) TestBaseUnmarshal() {
	type User struct {
		Name Type[string] `yaml:"name"`
		Age  Type[int]    `yaml:"age"`
	}

	type Data struct {
		User Struct[User] `json:"user"`
	}

	inputData := []byte(`{"user": {"name": "test", "age": 20}}`)

	expectData := Data{User: NewStruct[User](WithStructIsProvided[User](true))}
	expectData.User.value.Name = NewType(WithIsProvided[string](true), WithValue("test"))
	expectData.User.value.Age = NewType(WithIsProvided[int](true), WithValue(20))

	var testData Data

	err := json.Unmarshal(inputData, &testData)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectData, testData)
}

func (s *TestStructUnmarshalSuite) TestUnmarshalWithPrepareFN() {
	type User struct {
		Name Type[string] `yaml:"name"`
		Age  Type[int]    `yaml:"age"`
	}

	type Data struct {
		User Struct[User] `yaml:"user"`
	}

	inputData := []byte(`{"user": {"name": "test", "age": 20}}`)

	expectData := Data{User: NewStruct[User](WithStructIsProvided[User](true), WithStructTag[User]("user"))}
	expectData.User.value.Name = NewType(
		WithIsProvided[string](true),
		WithValue("test"),
		WithTag[string]("name"),
	)
	expectData.User.value.Age = NewType(
		WithIsProvided[int](true),
		WithValue(20),
		WithTag[int]("age"),
	)

	prepareFN := func() User {
		return User{
			Name: NewType(WithValue("test"), WithTag[string]("name")),
			Age:  NewType(WithValue(20), WithTag[int]("age")),
		}
	}

	testData := Data{
		User: NewStruct[User](
			WithStructTag[User]("user"),
			WithStructPrepareFN[User](prepareFN),
		),
	}

	err := json.Unmarshal(inputData, &testData)
	assert.NoError(s.T(), err)

	testData.User.prepareFN = nil
	assert.Equal(s.T(), expectData, testData)
}
