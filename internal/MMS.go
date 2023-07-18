package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"netWorkService/src"
	"sort"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func (data *MMSData) CheckCorrectProviders(providers []string) (result bool) {

	for _, provider := range providers {
		if data.Provider == provider {
			result = true
			return
		}
	}
	return
}

func (data *MMSData) HasCountryAlpha2Code() (result bool) {
	country := src.Countries[data.Country]
	if country != "" {
		result = true
	}
	return
}

func PrepareMMSData() [][]MMSData {
	sortByProvider, sortByCountry := serveMMSData()
	if len(sortByCountry) == 0 {
		return nil
	}
	dataStoreByProvider := make([]MMSData, 0, len(sortByProvider))
	for _, elem := range sortByProvider {
		mms := MMSData{
			elem.Country,
			elem.Provider,
			elem.Bandwidth,
			elem.ResponseTime,
		}
		dataStoreByProvider = append(dataStoreByProvider, mms)
	}

	dataStoreByCountry := make([]MMSData, 0, len(sortByCountry))
	for _, elem := range sortByCountry {
		mms := MMSData{
			elem.Country,
			elem.Provider,
			elem.Bandwidth,
			elem.ResponseTime,
		}
		dataStoreByCountry = append(dataStoreByCountry, mms)
	}
	dataStoreOrdered := make([][]MMSData, 0, 2)
	dataStoreOrdered = append(dataStoreOrdered, dataStoreByProvider, dataStoreByCountry)

	return dataStoreOrdered

}

func serveMMSData() ([]*MMSData, []*MMSData) {
	data := getMMSData()
	for _, elem := range data {
		countryName := getCountryName(elem.Country)
		elem.Country = countryName
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Provider < data[j].Provider })
	dataSortByProvider := make([]*MMSData, len(data))
	copy(dataSortByProvider, data)

	sort.Slice(data, func(i, j int) bool { return data[i].Country < data[j].Country })
	dataSortByCountry := make([]*MMSData, len(data))
	copy(dataSortByCountry, data)

	return dataSortByProvider, dataSortByCountry

}

func getMMSData() []*MMSData {
	dataStore := make([]*MMSData, 0)
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, errReadBody := io.ReadAll(resp.Body)
		if errReadBody != nil {
			return nil
		}
		providers := []string{"Topolo", "Rond", "Kildy"}
		var str []*MMSData
		if errJson := json.Unmarshal(body, &str); errJson != nil {
			return nil
		}

		for _, elem := range str {
			if elem.HasCountryAlpha2Code() {
				if elem.CheckCorrectProviders(providers) {
					dataStore = append(dataStore, elem)

				}
			}
		}

		return dataStore

	}
	return nil
}
