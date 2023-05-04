package services

type FizzBuzzServiceInterface interface {
	GetNumberOrAlias(i int) string
}

type FizzBuzzService struct {
	GetNumberOrAlias func(int) string
}

type aliasCondition struct {
	alias     string
	condition func(int) bool
}
