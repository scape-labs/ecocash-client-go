package ecocash

// ClientConfig holds the configuration for the EcoCash client
type ClientConfig struct {
	BaseURL             string
	Username            string
	Password            string
	MerchantCode        string
	MerchantPin         string
	MerchantNumber      string
	MerchantName        string
	SuperMerchantName   string
	TerminalID          string
	Location            string
	CountryCode         string // Usually "ZW"
	NotifyURL           string
	TraceRequests       bool
	DisableHttpWarnings bool
}