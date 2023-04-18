package usecase

import (
	"finalWork/internal/entity"
	"finalWork/src"
	"strings"
)

type UseCase struct {
	repo Infrastructure
}

func New(i Infrastructure) *UseCase {
	return &UseCase{
		repo: i,
	}
}

func (uc *UseCase) GetSMSData(data []byte) []*entity.SMSData {
	providers := []string{"Topolo", "Rond", "Kildy"}
	dataSlice := getDataSlice(data, providers)
	dataStruct := uc.repo.GetSMSData(dataSlice)
	return dataStruct
}

func getDataSlice(data []byte, providers []string) (dataSlice []string) {
	dataString := string(data)
	dataSlice = strings.Split(dataString, "\n")

	dataSlice = checkDataFields(dataSlice, 4, providers, 3)
	return
}

func checkDataFields(data []string, fieldsAmount uint, providers []string, indexForProvider int) (correctData []string) {
	for _, elem := range data {
		elemSlice := strings.Split(elem, ";")
		if len(elemSlice) == int(fieldsAmount) {
			if hasCountryAlpha2(elemSlice[0]) {
				for _, rightProvider := range providers {
					if rightProvider == elemSlice[indexForProvider] {
						correctData = append(correctData, elem)
					}
				}

			}

		}

	}
	return
}

func hasCountryAlpha2(code string) (result bool) {
	country := src.Countries[code]
	if country != "" {
		result = true
	}
	return
}
