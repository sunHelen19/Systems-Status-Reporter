package internal

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func PrepareBillingData() BillingData {
	data := getBillingData()

	billingData := BillingData{}
	if data != nil {
		billingData = BillingData{
			CreateCustomer: data.CreateCustomer,
			Purchase:       data.Purchase,
			Payout:         data.Payout,
			Recurring:      data.Recurring,
			FraudControl:   data.FraudControl,
			CheckoutPage:   data.CheckoutPage,
		}

	}
	return billingData
}

func getBillingData() *BillingData {
	data, err := readFile("src/simulator/data/billing.data")
	if err != nil {
		return nil
	}
	if len(data) != 0 {
		dataSum := getSumBits(data)
		dataSlice := dataSum.String()

		str := BillingData{
			CreateCustomer: dataSlice[0],
			Purchase:       dataSlice[1],
			Payout:         dataSlice[2],
			Recurring:      dataSlice[3],
			FraudControl:   dataSlice[4],
			CheckoutPage:   dataSlice[5],
		}

		return &str
	}
	return nil

}
