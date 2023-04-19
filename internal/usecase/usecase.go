package usecase

import (
	"encoding/json"
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
	dataSlice := getDataStringSlice(data, "\n", 4, providers, 3)
	dataStruct := uc.repo.GetSMSData(dataSlice)
	return dataStruct
}

func (uc *UseCase) GetMMSData(data []byte) (result []*entity.MMSData, err error) {
	providers := []string{"Topolo", "Rond", "Kildy"}
	var str []*entity.MMSData
	if errJson := json.Unmarshal(data, &str); errJson != nil {

		return nil, nil
	}

	for _, elem := range str {
		if elem.HasCountryAlpha2Code() {
			if elem.CheckCorrectProviders(providers) {

				result = uc.repo.GetMMSData(elem)

			}
		}
	}

	return
}

func (uc *UseCase) GetVoiceCallData(data []byte) []*entity.VoiceCallData {
	providers := []string{"TransparentCalls", "E-Voice", "JustPhone"}
	dataSlice := getDataStringSlice(data, "\n", 8, providers, 3)
	dataStruct := uc.repo.GetVoiceCallData(dataSlice)
	return dataStruct
}

func (uc *UseCase) GetEmailData(data []byte) []*entity.EmailData {
	providers := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Proton Mail", "Yandex", "Mail.ru"}
	dataSlice := getDataStringSlice(data, "\n", 3, providers, 1)
	dataStruct := uc.repo.GetEmailData(dataSlice)
	return dataStruct
}

func getDataStringSlice(data []byte, sep string, fieldsAmount uint, providers []string, indexForProvider int) (dataSlice []string) {
	dataString := string(data)
	dataSlice = strings.Split(dataString, sep)

	dataSlice = checkDataStringFields(dataSlice, fieldsAmount, providers, indexForProvider)
	return
}

func checkDataStringFields(data []string, fieldsAmount uint, providers []string, indexForProvider int) (correctData []string) {
	for _, elem := range data {
		elemSlice := strings.Split(elem, ";")
		if len(elemSlice) == int(fieldsAmount) {
			if hasStringCountryAlpha2(elemSlice[0]) {
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

func hasStringCountryAlpha2(code string) (result bool) {
	country := src.Countries[code]
	if country != "" {
		result = true
	}
	return
}
