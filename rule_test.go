package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestRequiredRuleSuite struct {
	suite.Suite
}

func (s *TestRequiredRuleSuite) TestInt() {
	t := NewType(true, 1, "test", B2I)
	err := RequiredRule[int]()(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 1, "test", B2I)
	err = RequiredRule[int]()(&t)
	assert.ErrorContains(s.T(), err, "test: not provided")
}

func (s *TestRequiredRuleSuite) TestFloat64() {
	t := NewType(true, 1.0, "test", B2Float64)
	err := RequiredRule[float64]()(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 1.0, "test", B2Float64)
	err = RequiredRule[float64]()(&t)
	assert.ErrorContains(s.T(), err, "test: not provided")
}

func (s *TestRequiredRuleSuite) TestString() {
	t := NewType(true, "some", "test", B2S)
	err := RequiredRule[string]()(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, "some", "test", B2S)
	err = RequiredRule[string]()(&t)
	assert.ErrorContains(s.T(), err, "test: not provided")
}

func (s *TestRequiredRuleSuite) TestBool() {
	t := NewType(true, true, "test", B2Bool)
	err := RequiredRule[bool]()(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, true, "test", B2Bool)
	err = RequiredRule[bool]()(&t)
	assert.ErrorContains(s.T(), err, "test: not provided")
}

func TestRequiredRule(t *testing.T) {
	suite.Run(t, new(TestRequiredRuleSuite))
}

type TestEqualRuleSuite struct {
	suite.Suite
}

func (s *TestEqualRuleSuite) TestInt() {
	t := NewType(true, 1, "test", B2I)
	err := EqualRule[int](1)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1, "test", B2I)
	err = EqualRule[int](1, 2, 3, 4, 5)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1, "test", B2I)
	err = EqualRule[int](0)(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [0]")

	t = NewType(true, 1, "test", B2I)
	err = EqualRule[int](0, 2, 3, 4, 5)(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [0 2 3 4 5]")
}

func (s *TestEqualRuleSuite) TestFloat64() {
	t := NewType(true, 1.0, "test", B2Float64)
	err := EqualRule[float64](1.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1.1, "test", B2Float64)
	err = EqualRule[float64](1.1, 2.0, 3.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 1.0, "test", B2Float64)
	err = EqualRule[float64](2)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1.0, "test", B2Float64)
	err = EqualRule[float64](0)(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [0]")

	t = NewType(true, 1.0, "test", B2Float64)
	err = EqualRule[float64](0, 2.0, 3.0)(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [0 2 3]")
}

func (s *TestEqualRuleSuite) TestString() {
	t := NewType(true, "some", "test", B2S)
	err := EqualRule[string]("some")(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, "some", "test", B2S)
	err = EqualRule[string]("some", "other")(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, "some", "test", B2S)
	err = EqualRule[string]("other")(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, "some", "test", B2S)
	err = EqualRule[string]("different")(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [different]")

	t = NewType(true, "some", "test", B2S)
	err = EqualRule[string]("diff", "erent")(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [diff erent]")
}

func (s *TestEqualRuleSuite) TestBool() {
	t := NewType(true, true, "test", B2Bool)
	err := EqualRule[bool](true)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, false, "test", B2Bool)
	err = EqualRule[bool](false, true)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, true, "test", B2Bool)
	err = EqualRule[bool](false)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, true, "test", B2Bool)
	err = EqualRule[bool](false)(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [false]")

	t = NewType(true, false, "test", B2Bool)
	err = EqualRule[bool](true)(&t)
	assert.ErrorContains(s.T(), err, "test: not equal [true]")
}

func TestEqualRule(t *testing.T) {
	suite.Run(t, &TestEqualRuleSuite{})
}

type TestNotEqualRuleSuite struct {
	suite.Suite
}

func (s *TestNotEqualRuleSuite) TestInt() {
	t := NewType(true, 1, "test", B2I)
	err := NotEqualRule[int](2)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1, "test", B2I)
	err = NotEqualRule[int](0, 2, 3)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1, "test", B2I)
	err = NotEqualRule[int](1)(&t)
	assert.ErrorContains(s.T(), err, "test: equal [1]")

	t = NewType(true, 1, "test", B2I)
	err = NotEqualRule[int](0, 1, 2)(&t)
	assert.ErrorContains(s.T(), err, "test: equal [0 1 2]")
}

func (s *TestNotEqualRuleSuite) TestFloat64() {
	t := NewType(true, 1.0, "test", B2Float64)
	err := NotEqualRule[float64](2.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1.0, "test", B2Float64)
	err = NotEqualRule[float64](0.5, 2.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 1.0, "test", B2Float64)
	err = NotEqualRule[float64](1.0)(&t)
	assert.ErrorContains(s.T(), err, "test: equal [1]")

	t = NewType(true, 1.0, "test", B2Float64)
	err = NotEqualRule[float64](0.0, 1.0, 2.0)(&t)
	assert.ErrorContains(s.T(), err, "test: equal [0 1 2]")
}

func (s *TestNotEqualRuleSuite) TestString() {
	t := NewType(true, "some", "test", B2S)
	err := NotEqualRule[string]("other")(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, "some", "test", B2S)
	err = NotEqualRule[string]("diff", "erent")(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, "some", "test", B2S)
	err = NotEqualRule[string]("some")(&t)
	assert.ErrorContains(s.T(), err, "test: equal [some]")

	t = NewType(true, "some", "test", B2S)
	err = NotEqualRule[string]("any", "some", "thing")(&t)
	assert.ErrorContains(s.T(), err, "test: equal [any some thing]")
}

func (s *TestNotEqualRuleSuite) TestBool() {
	t := NewType(true, true, "test", B2Bool)
	err := NotEqualRule[bool](false)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, false, "test", B2Bool)
	err = NotEqualRule[bool](true)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, true, "test", B2Bool)
	err = NotEqualRule[bool](true)(&t)
	assert.ErrorContains(s.T(), err, "test: equal [true]")

	t = NewType(true, false, "test", B2Bool)
	err = NotEqualRule[bool](false)(&t)
	assert.ErrorContains(s.T(), err, "test: equal [false]")
}

func TestNotEqualRule(t *testing.T) {
	suite.Run(t, &TestNotEqualRuleSuite{})
}

type TestGTERuleSuite struct {
	suite.Suite
}

func (s *TestGTERuleSuite) TestInt() {
	t := NewType(true, 5, "test", B2I)
	err := GTERule[int](5)(&t)
	assert.NoError(s.T(), err)

	err = GTERule[int](4)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 5, "test", B2I)
	err = GTERule[int](5)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 3, "test", B2I)
	err = GTERule[int](5)(&t)
	assert.ErrorContains(s.T(), err, "test: not greater than or equal 5")
}

func (s *TestGTERuleSuite) TestFloat64() {
	t := NewType(true, 5.0, "test", B2Float64)
	err := GTERule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	err = GTERule[float64](4.9)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 5.0, "test", B2Float64)
	err = GTERule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 3.0, "test", B2Float64)
	err = GTERule[float64](5.0)(&t)
	assert.ErrorContains(s.T(), err, "test: not greater than or equal 5")
}

func TestGTERule(t *testing.T) {
	suite.Run(t, &TestGTERuleSuite{})
}

type TestGTRuleSuite struct {
	suite.Suite
}

func (s *TestGTRuleSuite) TestInt() {
	t := NewType(true, 6, "test", B2I)
	err := GTRule[int](5)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 6, "test", B2I)
	err = GTRule[int](5)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 5, "test", B2I)
	err = GTRule[int](5)(&t)
	assert.ErrorContains(s.T(), err, "test: not greater than 5")

	t = NewType(true, 4, "test", B2I)
	err = GTRule[int](5)(&t)
	assert.ErrorContains(s.T(), err, "test: not greater than 5")
}

func (s *TestGTRuleSuite) TestFloat64() {
	t := NewType(true, 5.1, "test", B2Float64)
	err := GTRule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 5.1, "test", B2Float64)
	err = GTRule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 5.0, "test", B2Float64)
	err = GTRule[float64](5.0)(&t)
	assert.ErrorContains(s.T(), err, "test: not greater than 5")

	t = NewType(true, 4.9, "test", B2Float64)
	err = GTRule[float64](5.0)(&t)
	assert.ErrorContains(s.T(), err, "test: not greater than 5")
}

func TestGTRule(t *testing.T) {
	suite.Run(t, &TestGTRuleSuite{})
}

type TestLTERuleSuite struct {
	suite.Suite
}

func (s *TestLTERuleSuite) TestInt() {
	t := NewType(true, 5, "test", B2I)
	err := LTERule[int](5)(&t)
	assert.NoError(s.T(), err)

	err = LTERule[int](6)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 5, "test", B2I)
	err = LTERule[int](5)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 7, "test", B2I)
	err = LTERule[int](5)(&t)
	assert.ErrorContains(s.T(), err, "test: not less than or equal 5")
}

func (s *TestLTERuleSuite) TestFloat64() {
	t := NewType(true, 5.0, "test", B2Float64)
	err := LTERule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	err = LTERule[float64](5.1)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 5.0, "test", B2Float64)
	err = LTERule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 5.2, "test", B2Float64)
	err = LTERule[float64](5.0)(&t)
	assert.ErrorContains(s.T(), err, "test: not less than or equal 5")
}

func TestLTERule(t *testing.T) {
	suite.Run(t, &TestLTERuleSuite{})
}

type TestLTRuleSuite struct {
	suite.Suite
}

func (s *TestLTRuleSuite) TestInt() {
	t := NewType(true, 4, "test", B2I)
	err := LTRule[int](5)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 4, "test", B2I)
	err = LTRule[int](5)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 5, "test", B2I)
	err = LTRule[int](5)(&t)
	assert.ErrorContains(s.T(), err, "test: not less than 5")

	t = NewType(true, 6, "test", B2I)
	err = LTRule[int](5)(&t)
	assert.ErrorContains(s.T(), err, "test: not less than 5")
}

func (s *TestLTRuleSuite) TestFloat64() {
	t := NewType(true, 4.9, "test", B2Float64)
	err := LTRule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(false, 4.9, "test", B2Float64)
	err = LTRule[float64](5.0)(&t)
	assert.NoError(s.T(), err)

	t = NewType(true, 5.0, "test", B2Float64)
	err = LTRule[float64](5.0)(&t)
	assert.ErrorContains(s.T(), err, "test: not less than 5")

	t = NewType(true, 5.1, "test", B2Float64)
	err = LTRule[float64](5.0)(&t)
	assert.ErrorContains(s.T(), err, "test: not less than 5")
}

func TestLTRule(t *testing.T) {
	suite.Run(t, &TestLTRuleSuite{})
}
