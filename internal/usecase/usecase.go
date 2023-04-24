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
			providerstoResult := make([][]*entity.EmailData, 0, 2)
			providerstoResult = append(providerstoResult, fastProviders, slowProviders)
			result[data[i-1].Country] = providerstoResult
			providers = make([]*entity.EmailData, 0, 0)

			providers = append(providers, data[i])

		}

	}

	return result
}

func (uc *UseCase) GetBillingData() *entity.BillingData {
	data := uc.repo.GetBillingData()
	return data

}

func (uc *UseCase) GetSupportData() []int {
	result := make([]int, 0, 2)
	data := uc.repo.GetSupportData()
	sumTickets := 0
	workload := 1

	for _, topic := range data {
		sumTickets += topic.ActiveTickets
	}

	if sumTickets >= 9 && sumTickets <= 16 {
		workload = 2
	} else if sumTickets > 16 {
		workload = 3
	}
	result = append(result, workload)

	var waitTime int
	oneTicketTime := 60 / 18
	waitTime = oneTicketTime * sumTickets
	result = append(result, waitTime)

	return result

}

func (uc *UseCase) GetIncidentData() []*entity.IncidentData {
	data := uc.repo.GetIncidentData()

	sort.Slice(data, func(i, j int) bool {
		return data[i].Status == "active"
	})

	return data

}

func getCountryName(code string) string {
	countryName := src.Countries[code]
	return countryName
}
