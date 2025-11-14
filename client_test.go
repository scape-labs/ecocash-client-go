package ecocash

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	config := ClientConfig{
		BaseURL:             "https://api.ecocash.co.zw",
		Username:            "test-user",
		Password:            "test-pass",
		MerchantCode:        "TEST123",
		MerchantPin:         "1234",
		MerchantNumber:      "263123456789",
		MerchantName:        "Test Business",
		SuperMerchantName:   "Test Super Merchant",
		TerminalID:          "TERM001",
		Location:            "Harare",
		CountryCode:         "ZW",
		NotifyURL:           "https://example.com/webhook",
		TraceRequests:       false,
		DisableHttpWarnings: true,
	}

	client := NewClient(config)

	if client == nil {
		t.Fatal("NewClient returned nil")
	}

	if client.config.Username != config.Username {
		t.Errorf("Expected username %s, got %s", config.Username, client.config.Username)
	}

	if client.config.MerchantCode != config.MerchantCode {
		t.Errorf("Expected merchant code %s, got %s", config.MerchantCode, client.config.MerchantCode)
	}
}

func TestCreateBasePaymentAmount(t *testing.T) {
	amount := 25.50
	currency := "USD"
	description := "Test payment"

	payment := createBasePaymentAmount(amount, currency, description)

	if payment.ChargingInformation.Amount != amount {
		t.Errorf("Expected amount %f, got %f", amount, payment.ChargingInformation.Amount)
	}

	if payment.ChargingInformation.Currency != currency {
		t.Errorf("Expected currency %s, got %s", currency, payment.ChargingInformation.Currency)
	}

	if payment.ChargingInformation.Description != description {
		t.Errorf("Expected description %s, got %s", description, payment.ChargingInformation.Description)
	}

	if payment.ChargeMetaData.Channel != "WEB" {
		t.Errorf("Expected channel WEB, got %s", payment.ChargeMetaData.Channel)
	}

	if payment.ChargeMetaData.PurchaseCategoryCode != "Online Payment" {
		t.Errorf("Expected purchase category code 'Online Payment', got %s", payment.ChargeMetaData.PurchaseCategoryCode)
	}

	if payment.ChargeMetaData.OnBehalfOf != "Online Payment" {
		t.Errorf("Expected on behalf of 'Online Payment', got %s", payment.ChargeMetaData.OnBehalfOf)
	}
}