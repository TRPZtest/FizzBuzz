package main

import (
	services "FizzBuzz/services"
	"fmt"
)

func main() {

	builder := services.NewStandartFizzBuzzBuilder()

	fizzBuzzService := builder.GetFizzBuzzService()

	fmt.Println(fizzBuzzService.GetNumberOrAlias(75))

	for i := 0; i <= 100; i++ {
		fmt.Println(fizzBuzzService.GetNumberOrAlias(i))
	}
}
