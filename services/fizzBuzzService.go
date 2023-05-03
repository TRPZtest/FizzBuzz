package services

type FizzBuzzServiceInterface interface {
	GetNumberOrAlias(i int) string
}

type FizzBuzzService struct {
	aliasConditionList []aliasCondition
	rule               func(n int, aliasConditionList []aliasCondition) string
}

type aliasCondition struct {
	alias     string
	condition func(int) bool
}

func (s *FizzBuzzService) GetNumberOrAlias(n int) string {

	return s.rule(n, s.aliasConditionList)
}
