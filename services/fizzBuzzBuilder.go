package services

import (
	"strconv"
)

type FizzBuzzBuilderInterface interface {
	GetFizzBuzzService() FizzBuzzService
}

type StandartFizzBuzzBuilder struct {
	aliasConditionList []aliasCondition
	rule               func(n int) string
}

func NewStandartFizzBuzzBuilder() *StandartFizzBuzzBuilder {
	return &StandartFizzBuzzBuilder{}
}

func (b *StandartFizzBuzzBuilder) setAliasConditionList() {
	b.aliasConditionList = []aliasCondition{
		{alias: "Fizz", condition: func(i int) bool { return i%3 == 0 }},
		{alias: "Buzz", condition: func(i int) bool { return i%5 == 0 }},
	}
}

func (b *StandartFizzBuzzBuilder) setRule() {
	b.rule = func(n int) string {

		output := ""

		for _, item := range b.aliasConditionList {
			if item.condition(n) {
				output = output + item.alias
			}
		}

		if output == "" {
			return strconv.Itoa(n)
		} else {
			return output
		}

	}
}

func (b *StandartFizzBuzzBuilder) GetFizzBuzzService() FizzBuzzService {

	b.setAliasConditionList()
	b.setRule()

	return FizzBuzzService{
		GetNumberOrAlias: b.rule,
	}
}
