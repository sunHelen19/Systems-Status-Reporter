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
	data := uc.repo.GetSMSData()
	for _, elem := range data {
		countryName := getCountryName(elem.Country)
		elem.Country = countryName
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Provider < data[j].Provider })
	dataSortByProvider := make([]*entity.SMSData, len(data))
	copy(dataSortByProvider, data)

	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	dataSortByCountry := make([]*entity.SMSData, len(data))
	copy(dataSortByCountry, data)

	return dataSortByProvider, dataSortByCountry

}

func (uc *UseCase) GetMMSData() ([]*entity.MMSData, []*entity.MMSData) {
	data := uc.repo.GetMMSData()
	for _, elem := range data {
		countryName := getCountryName(elem.Country)
		elem.Country = countryName
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Provider < data[j].Provider })
	dataSortByProvider := make([]*entity.MMSData, len(data))
	copy(dataSortByProvider, data)

	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	dataSortByCountry := make([]*entity.MMSData, len(data))
	copy(dataSortByCountry, data)

	return dataSortByProvider, dataSortByCountry

}

func (uc *UseCase) GetVoiceCallData() []*entity.VoiceCallData {
	data := uc.repo.GetVoiceCallData()
	return data

}

func (uc *UseCase) GetEmailData() map[string][][]*entity.EmailData {
	data := uc.repo.GetEmailData()

	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	result := make(map[string][][]*entity.EmailData)
	providers := make([]*entity.EmailData, 0, 0)
	providers = append(providers, data[0])
	fastProviders := make([]*entity.EmailData, 0, 3)
	slowProviders := make([]*entity.EmailData, 0, 3)
	for i := 1; i < len(data); i++ {

		if data[i].Country == data[i-1].Country {
			providers = append(providers, data[i])

		} else {
			sort.Slice(providers, func(i, j int) bool { return providers[i].DeliveryTime < providers[j].DeliveryTime })
			length := len(providers)

			if length >= 3 {
				fastProviders = providers[length-3:]
				slowProviders = providers[0:3]
			} else {
				fastProviders = providers
				slowProviders = providers
			}
			/*length := len(providers)
			for i := 0; length > 0 && i < 3; i++ {
				fastProviders = append(fastProviders, providers[length-1])
				length--
			}

			length = len(providers)
			for i := 0; length > 0 && i < 3; i++ {
				slowProviders = append(slowProviders, providers[i])
				length--
			}

			*/
			providerstoResult := make([][]*entity.EmailData, 0, 2)
			providerstoResult = append(providerstoResult, fastProviders, slowProviders)
			result[data[i-1].Country] = providerstoResult
			providers = make([]*entity.EmailData, 0, 0)

			providers = append(providers, data[i])

		}

	}

	return result
}

func getCountryName(code string) string {
	countryName := src.Countries[code]
	return countryName
}
