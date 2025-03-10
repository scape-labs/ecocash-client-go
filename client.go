package ecocash

import (
	"fmt"
	"resty.dev/v3"
)

// Client represents an EcoCash API client
type Client struct {
	client *resty.Client
	config ClientConfig
}

// NewClient creates a new EcoCash client with the provided configuration
func NewClient(config ClientConfig) *Client {
	client := resty.New()
	client.SetBasicAuth(config.Username, config.Password)
	client.SetDisableWarn(config.DisableHttpWarnings)

	return &Client{
		client: client,
		config: config,
	}
}

// Charge initiates a payment request with simplified parameters
func (c *Client) Charge(request ChargeSubscriberRequest) (*TransactionResponse, error) {
	req := &chargeRequest{
		ClientCorrelator:           request.ReferenceCode,
		NotifyURL:                  c.config.NotifyURL,
		ReferenceCode:              request.ReferenceCode,
		TranType:                   transactionType,
		EndUserID:                  request.PhoneNumber,
		Remarks:                    request.Description,
		TransactionOperationStatus: chargedStatus,
		PaymentAmount:              createBasePaymentAmount(request.Amount, request.Currency, request.Description),
		MerchantCode:               c.config.MerchantCode,
		MerchantPin:                c.config.MerchantPin,
		MerchantNumber:             c.config.MerchantNumber,
		CurrencyCode:               request.Currency,
		CountryCode:                c.config.CountryCode,
		TerminalID:                 c.config.TerminalID,
		Location:                   c.config.Location,
		SuperMerchantName:          c.config.SuperMerchantName,
		MerchantName:               c.config.MerchantName,
	}

	var response TransactionResponse

	resp, err := c.client.R().
		SetBody(req).
		SetResult(&response).
		SetDebug(c.config.TraceRequests).
		Post(c.config.BaseURL + "/payment/v1/transactions/amount")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("charge request failed with status code: %d", resp.StatusCode())
	}

	return &response, nil
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
		SetDebug(c.config.TraceRequests).
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
