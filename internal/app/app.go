package app

import (
	"finalWork/internal/controller"
	"finalWork/internal/infrastructure"
	"finalWork/internal/usecase"
	"fmt"
)

func Run() {
	repository := infrastructure.CreateStore()
	useCase := usecase.New(repository)
	c := controller.New(useCase)

	data := c.GetBillingData()

	for _, elem := range data {
		fmt.Println(elem)
	}

	//fmt.Println(string(data))
}
