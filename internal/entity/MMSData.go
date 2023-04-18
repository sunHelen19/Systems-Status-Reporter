package entity

import "finalWork/src"

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
