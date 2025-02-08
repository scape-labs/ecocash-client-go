package ecocash

import (
	"fmt"
	"resty.dev/v3"
)

// ClientConfig holds the configuration for the EcoCash client
type ClientConfig struct {
	BaseURL           string
	Username          string
	Password          string
	MerchantCode      string
	MerchantPin       string
	MerchantNumber    string
	MerchantName      string
	SuperMerchantName string
	TerminalID        string
	Location          string
	CountryCode       string // Usually "ZW"
	NotifyURL         string
}

// Client represents an EcoCash API client
type Client struct {
	client *resty.Client
	config ClientConfig
}

// SimplePayment represents the simplified payment request
type SimplePayment struct {
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

// NewClient creates a new EcoCash client with the provided configuration
func NewClient(config ClientConfig) *Client {
	client := resty.New()
	client.SetBasicAuth(config.Username, config.Password)

	return &Client{
		client: client,
		config: config,
	}
}

// createBasePaymentAmount creates a base payment amount structure
func createBasePaymentAmount(amount float64, currency, description string) paymentAmount {
	return paymentAmount{
		ChargingInformation: struct {
			Amount      float64 `json:"amount"`
			Currency    string  `json:"currency"`
			Description string  `json:"description"`
		}{
			Amount:      amount,
			Currency:    currency,
			Description: description,
		},
		ChargeMetaData: struct {
			Channel              string `json:"channel"`
			PurchaseCategoryCode string `json:"purchaseCategoryCode"`
			OnBehalfOf           string `json:"onBeHalfOf"`
		}{
			Channel:              "WEB",
			PurchaseCategoryCode: "Online Payment",
			OnBehalfOf:           "Online Payment",
		},
	}
}

// Charge initiates a payment request with simplified parameters
func (c *Client) Charge(payment SimplePayment) (*TransactionResponse, error) {
	req := &chargeRequest{
		ClientCorrelator:           payment.ReferenceCode,
		NotifyURL:                  c.config.NotifyURL,
		ReferenceCode:              payment.ReferenceCode,
		TranType:                   "MER",
		EndUserID:                  payment.PhoneNumber,
		Remarks:                    payment.Description,
		TransactionOperationStatus: "Charged",
		PaymentAmount:              createBasePaymentAmount(payment.Amount, payment.Currency, payment.Description),
		MerchantCode:               c.config.MerchantCode,
		MerchantPin:                c.config.MerchantPin,
		MerchantNumber:             c.config.MerchantNumber,
		CurrencyCode:               payment.Currency,
		CountryCode:                c.config.CountryCode,
		TerminalID:                 c.config.TerminalID,
		Location:                   c.config.Location,
		SuperMerchantName:          c.config.SuperMerchantName,
		MerchantName:               c.config.MerchantName,
	}

	resp, err := c.client.R().
		SetBody(req).
		SetResult(&TransactionResponse{}).
		Post(c.config.BaseURL + "/payment/v1/transactions/amount")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("charge request failed with status code: %d", resp.StatusCode())
	}

	return resp.Result().(*TransactionResponse), nil
}

// Refund initiates a refund request with simplified parameters
func (c *Client) Refund(refund SimpleRefund) (*TransactionResponse, error) {
	req := &chargeRequest{
		ClientCorrelator:  refund.ReferenceCode,
		ReferenceCode:     refund.ReferenceCode,
		EndUserID:         refund.PhoneNumber,
		TranType:          "REF",
		Remarks:           refund.Description,
		PaymentAmount:     createBasePaymentAmount(refund.Amount, refund.Currency, refund.Description),
		MerchantCode:      c.config.MerchantCode,
		MerchantPin:       c.config.MerchantPin,
		MerchantNumber:    c.config.MerchantNumber,
		CurrencyCode:      refund.Currency,
		CountryCode:       c.config.CountryCode,
		TerminalID:        c.config.TerminalID,
		Location:          c.config.Location,
		SuperMerchantName: c.config.SuperMerchantName,
		MerchantName:      c.config.MerchantName,
	}

	resp, err := c.client.R().
		SetBody(req).
		SetResult(&TransactionResponse{}).
		Post(c.config.BaseURL + "/payment/v1/transactions/refund")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("refund request failed with status code: %d", resp.StatusCode())
	}

	return resp.Result().(*TransactionResponse), nil
}

// QueryTransaction queries the status of a transaction
func (c *Client) QueryTransaction(phoneNumber, referenceCode string) (*TransactionResponse, error) {
	resp, err := c.client.R().
		SetResult(&TransactionResponse{}).
		Get(fmt.Sprintf("%s/payment/v1/%s/transactions/amount/%s",
			c.config.BaseURL, phoneNumber, referenceCode))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("query transaction failed with status code: %d", resp.StatusCode())
	}

	return resp.Result().(*TransactionResponse), nil
}
