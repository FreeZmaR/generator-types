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

	expectData := Data{User: NewStruct[User](true, "", nil)}
	expectData.User.value.Name = NewType(true, "test", "", nil)
	expectData.User.value.Age = NewType(true, 20, "", nil)

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

	expectData := Data{User: NewStruct[User](true, "user", nil)}
	expectData.User.value.Name = NewType(true, "test", "name", nil)
	expectData.User.value.Age = NewType(true, 20, "age", nil)

	prepareFN := func() User {
		return User{
			Name: NewType(false, "test", "name", nil),
			Age:  NewType(false, 20, "age", nil),
		}
	}

	testData := Data{
		User: NewStruct[User](false, "user", prepareFN),
	}

	err := json.Unmarshal(inputData, &testData)
	assert.NoError(s.T(), err)

	testData.User.prepareFN = nil
	assert.Equal(s.T(), expectData, testData)
}
