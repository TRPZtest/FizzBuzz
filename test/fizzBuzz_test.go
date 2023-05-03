package test

import (
	services "FizzBuzz/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplicityOfThree(t *testing.T) {
	builder := services.NewStandartFizzBuzzBuilder()

	fizzBuzzService := builder.GetFizzBuzzService()

	numStr := fizzBuzzService.GetNumberOrAlias(48)

	assert.EqualValues(t, "Fizz", numStr)
}

func TestMultiplicityOfFive(t *testing.T) {
	builder := services.NewStandartFizzBuzzBuilder()

	fizzBuzzService := builder.GetFizzBuzzService()

	numStr := fizzBuzzService.GetNumberOrAlias(65)

	assert.EqualValues(t, "Buzz", numStr)
}

func TestMultiplicityOfThreeAndFive(t *testing.T) {
	builder := services.NewStandartFizzBuzzBuilder()

	fizzBuzzService := builder.GetFizzBuzzService()

	numStr := fizzBuzzService.GetNumberOrAlias(45)

	assert.EqualValues(t, "FizzBuzz", numStr)
}

func TestNotMultiplicityOfThreeAndFive(t *testing.T) {
	builder := services.NewStandartFizzBuzzBuilder()

	fizzBuzzService := builder.GetFizzBuzzService()

	numStr := fizzBuzzService.GetNumberOrAlias(34)

	assert.EqualValues(t, "34", numStr)
}
