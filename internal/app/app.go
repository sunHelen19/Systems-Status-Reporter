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

	data, err := c.GetSMSData()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, elem := range data {
		fmt.Println(elem)
	}
}
