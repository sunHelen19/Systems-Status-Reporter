package usecase

import (
	"finalWork/internal/entity"
	"finalWork/src"
	"sort"
)

type UseCase struct {
	repo Infrastructure
}

func New(i Infrastructure) *UseCase {
	return &UseCase{
		repo: i,
	}
}

func (uc *UseCase) GetSMSData() ([]*entity.SMSData, []*entity.SMSData) {
	smsData := uc.repo.GetSMSData()
	for _, elem := range smsData {
		countryName := getCountryName(elem.Country)
		elem.Country = countryName
	}
	sort.Slice(smsData, func(i, j int) bool { return smsData[i].Provider < smsData[j].Provider })
	smsDataSortByProvider := make([]*entity.SMSData, len(smsData))
	copy(smsDataSortByProvider, smsData)

	sort.Slice(smsData, func(i, j int) bool { return smsData[i].Country < smsData[j].Country })
	smsDataSortByCountry := make([]*entity.SMSData, len(smsData))
	copy(smsDataSortByCountry, smsData)

	return smsDataSortByProvider, smsDataSortByCountry

}

func getCountryName(code string) string {
	countryName := src.Countries[code]
	return countryName
}
