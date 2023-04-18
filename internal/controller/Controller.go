package controller

import (
	"finalWork/internal/entity"
	"finalWork/internal/usecase"
	"fmt"
	"io"
	"os"
)

type Controller struct {
	uc usecase.Controller
}

func New(uc usecase.Controller) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) GetSMSData() ([]*entity.SMSData, error) {
	data, err := readFile("src/simulator/sms.data")
	if err != nil {
		return nil, err
	}
	dataSlice := c.uc.GetSMSData(data)
	return dataSlice, nil
}

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Ошибка открытия файла", err)
		return nil, err
	}
	defer file.Close()
	resultBytes, errRB := io.ReadAll(file)
	if errRB != nil {
		panic(err)
	}
	return resultBytes, nil
}
