package data

import (
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func PrepareEmailData(path string, providers []string) map[string][][]EmailData {
	data := serveEmailData(path, providers)
	if len(data) == 0 {
		return nil
	}
	dataStore := make(map[string][][]EmailData)

	for key, country := range data {
		providers := make([][]EmailData, 0, 2)
		fastProviders := make([]EmailData, 0, 3)

		for _, elem := range country[0] {
			provider := EmailData{
				Country:      elem.Country,
				Provider:     elem.Provider,
				DeliveryTime: elem.DeliveryTime,
			}
			fastProviders = append(fastProviders, provider)

		}

		slowProviders := make([]EmailData, 0, 3)
		for _, elem := range country[1] {
			provider := EmailData{
				Country:      elem.Country,
				Provider:     elem.Provider,
				DeliveryTime: elem.DeliveryTime,
			}
			slowProviders = append(slowProviders, provider)
		}
		sort.Slice(fastProviders, func(i, j int) bool { return fastProviders[i].DeliveryTime > fastProviders[j].DeliveryTime })
		providers = append(providers, fastProviders, slowProviders)
		dataStore[key] = providers
	}

	return dataStore

}

func serveEmailData(path string, providers []string) map[string][][]*EmailData {
	data := getEmailData(path, providers)
	if len(data) == 0 {
		return nil
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	result := make(map[string][][]*EmailData)
	providersEmail := make([]*EmailData, 0)

	providersEmail = append(providersEmail, data[0])
	var fastProviders []*EmailData
	var slowProviders []*EmailData

	for i := 1; i < len(data); i++ {

		if data[i].Country == data[i-1].Country {
			providersEmail = append(providersEmail, data[i])

		} else {
			sort.Slice(providersEmail, func(i, j int) bool { return providersEmail[i].DeliveryTime < providersEmail[j].DeliveryTime })
			length := len(providers)

			if length >= 3 {
				fastProviders = providersEmail[length-3:]
				slowProviders = providersEmail[0:3]
			} else {
				fastProviders = providersEmail
				slowProviders = providersEmail
			}

			providersToResult := make([][]*EmailData, 0, 2)
			providersToResult = append(providersToResult, fastProviders, slowProviders)
			result[data[i-1].Country] = providersToResult
			providersEmail = make([]*EmailData, 0)

			providersEmail = append(providersEmail, data[i])

		}

	}

	return result

}

func getEmailData(path string, providers []string) []*EmailData {

	data, err := readFile(path)
	if err != nil {
		return nil
	}

	//	providers := []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail", "Yandex", "Mail.ru"}
	dataSlice := getDataStringSlice(data, "\n", 3, providers, 1)
	dataStore := make([]*EmailData, 0)
	for _, elem := range dataSlice {

		elemSlice := strings.Split(elem, ";")

		deliveryTime, errDT := strconv.Atoi(elemSlice[2])
		if errDT != nil {
			continue
		}

		str := EmailData{
			Country:      elemSlice[0],
			Provider:     elemSlice[1],
			DeliveryTime: deliveryTime,
		}
		dataStore = append(dataStore, &str)

	}

	return dataStore
}
