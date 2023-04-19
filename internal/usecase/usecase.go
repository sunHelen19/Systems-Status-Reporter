package usecase

import (
	"encoding/json"
	"finalWork/internal/entity"
	"finalWork/src"
	"math"
	"strconv"
	"strings"
)

type keySet uint8

type UseCase struct {
	repo Infrastructure
}

const (
	CreateCustomer keySet = 1 << iota
	Purchase
	Payout
	Recurring
	FraudControl
	CheckoutPage
	maxKey
)

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

func (uc *UseCase) GetBillingData(data []byte) []*entity.BillingData {
	dataSum := getSumBits(data)
	dataSlice := dataSum.String()
	dataStruct := uc.repo.GetBillingData(dataSlice)
	return dataStruct
}

func (uc *UseCase) GetSupportData(data []byte) (result []*entity.SupportData, err error) {

	var str []*entity.SupportData
	if errJson := json.Unmarshal(data, &str); errJson != nil {

		return nil, nil
	}
	for _, elem := range str {

		result = uc.repo.GetSupportData(elem)

	}
	return
}

func (uc *UseCase) GetIncidentData(data []byte) (result []*entity.IncidentData, err error) {

	var str []*entity.IncidentData
	if errJson := json.Unmarshal(data, &str); errJson != nil {

		return nil, nil
	}
	for _, elem := range str {

		result = uc.repo.GetIncidentData(elem)

	}
	return
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

func (k keySet) String() (data []bool) {
	if k >= maxKey {
		panic("Broken keyset")
	}

	for key := CreateCustomer; key < maxKey; key <<= 1 {
		if k&key != 0 {
			data = append(data, true)
		} else {
			data = append(data, false)
		}
	}
	return
}

func getSumBits(data []byte) (sum keySet) {
	length := len(data)
	dataString := string(data)

	for index, elem := range dataString {
		elemStr := string(elem)
		elemInt, errElemInt := strconv.Atoi(elemStr)
		if errElemInt != nil {
			panic(errElemInt)
		}

		if elemInt == 1 {
			sum += keySet(math.Pow(2, float64(length-1-index)))

		}
	}
	return
}
