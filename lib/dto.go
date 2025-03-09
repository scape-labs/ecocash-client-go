package ecocash

type chargeRequest struct {
	ClientCorrelator           string        `json:"clientCorrelator"`
	NotifyURL                  string        `json:"notifyUrl"`
	ReferenceCode              string        `json:"referenceCode"`
	TranType                   string        `json:"tranType"`
	EndUserID                  string        `json:"endUserId"`
	Remarks                    string        `json:"remarks"`
	TransactionOperationStatus string        `json:"transactionOperationStatus"`
	PaymentAmount              paymentAmount `json:"paymentAmount"`
	MerchantCode               string        `json:"merchantCode"`
	MerchantPin                string        `json:"merchantPin"`
	MerchantNumber             string        `json:"merchantNumber"`
	CurrencyCode               string        `json:"currencyCode"`
	CountryCode                string        `json:"countryCode"`
	TerminalID                 string        `json:"terminalID"`
	Location                   string        `json:"location"`
	SuperMerchantName          string        `json:"superMerchantName"`
	MerchantName               string        `json:"merchantName"`
}


// ChargeSubscriberRequest represents the simplified payment request
type ChargeSubscriberRequest struct {
	ReferenceCode string  // Your unique reference for this transaction
	PhoneNumber   string  // Customer's phone number
	Amount        float64 // Amount to charge
	Currency      string  // "USD" or "ZWG"
	Description   string  // Optional description
}

// SimpleRefund represents the simplified refund request
type SimpleRefund struct {
	ReferenceCode            string  // Your unique reference for this refund
	PhoneNumber              string  // Customer's phone number
	Amount                   float64 // Amount to refund
	Currency                 string  // "USD" or "ZWG"
	OriginalEcocashReference string  // Original transaction reference
	Description              string  // Optional description
}

// Internal types for API communication
type paymentAmount struct {
	ChargingInformation struct {
		Amount      float64 `json:"amount"`
		Currency    string  `json:"currency"`
		Description string  `json:"description"`
	} `json:"charginginformation"`
	ChargeMetaData struct {
		Channel              string `json:"channel"`
		PurchaseCategoryCode string `json:"purchaseCategoryCode"`
		OnBehalfOf           string `json:"onBeHalfOf"`
	} `json:"chargeMetaData"`
}

// TransactionResponse represents the API response for all operations
type TransactionResponse struct {
	ClientCorrelator           string        `json:"clientCorrelator"`
	EndTime                    int64         `json:"endTime"`
	StartTime                  int64         `json:"startTime"`
	ServerReferenceCode        string        `json:"serverReferenceCode"`
	TransactionOperationStatus string        `json:"transactionOperationStatus"`
	ResponseCode               string        `json:"responseCode"`
	EcocashReference           string        `json:"ecocashReference"`
	PaymentAmount              paymentAmount `json:"paymentAmount"`
}
