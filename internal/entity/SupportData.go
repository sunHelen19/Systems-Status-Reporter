package entity

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

/*
func (data *SupportData) CheckCorrectProviders(providers []string) (result bool) {

	for _, provider := range providers {
		if data.Provider == provider {
			result = true
			return
		}
	}
	return
}

func (data *SupportData) HasCountryAlpha2Code() (result bool) {
	country := src.Countries[data.Country]
	if country != "" {
		result = true
	}
	return
}
*/
