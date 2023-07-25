package data

import (
	"sort"
	"strings"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func PrepareSMSData(path string) [][]SMSData {
	sortByProvider, sortByCountry := serveSMSData(path)
	if len(sortByCountry) == 0 {
		return nil
	}
	dataStoreByProvider := make([]SMSData, 0, len(sortByProvider))
	for _, elem := range sortByProvider {
		sms := SMSData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
		}
		dataStoreByProvider = append(dataStoreByProvider, sms)
	}

	dataStoreByCountry := make([]SMSData, 0, len(sortByCountry))
	for _, elem := range sortByCountry {
		sms := SMSData{
			elem.Country,
			elem.Bandwidth,
			elem.ResponseTime,
			elem.Provider,
		}
		dataStoreByCountry = append(dataStoreByCountry, sms)
	}
	dataStoreOrdered := make([][]SMSData, 0, 2)
	dataStoreOrdered = append(dataStoreOrdered, dataStoreByProvider, dataStoreByCountry)

	return dataStoreOrdered

}

func serveSMSData(path string) ([]*SMSData, []*SMSData) {
	data := getSMSData(path)
	for _, elem := range data {
		countryName := getCountryName(elem.Country)
		elem.Country = countryName
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Provider < data[j].Provider })
	dataSortByProvider := make([]*SMSData, len(data))
	copy(dataSortByProvider, data)

	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	dataSortByCountry := make([]*SMSData, len(data))
	copy(dataSortByCountry, data)

	return dataSortByProvider, dataSortByCountry

}

func getSMSData(path string) []*SMSData {
	dataStore := make([]*SMSData, 0)
	data, err := readFile(path)
	if err != nil {
		return nil
	}
	providers := []string{"Topolo", "Rond", "Kildy"}
	dataSlice := getDataStringSlice(data, "\n", 4, providers, 3)

	for _, elem := range dataSlice {
		elemSlice := strings.Split(elem, ";")
		str := SMSData{
			Country:      elemSlice[0],
			Bandwidth:    elemSlice[1],
			ResponseTime: elemSlice[2],
			Provider:     elemSlice[3],
		}
		dataStore = append(dataStore, &str)

	}

	return dataStore
}
