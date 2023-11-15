package model_wsp

type BrowserPaymentRequest struct {
	ReturnURL   string `json:"returnUrl,omitempty"`
	CallbackURL string `json:"callbackUrl,omitempty"`
}

type Account struct {
	ID string `json:"id,omitempty"`
}

type Customer struct {
	Account Account `json:"account,omitempty"`
	Email   string  `json:"email,omitempty"`
	Name    string  `json:"name,omitempty"`
	Phone   string  `json:"phone,omitempty"`
}

type SourceOfFunds struct {
	Types []string `json:"types,omitempty"`
}

type Device struct {
	IPAddress        string `json:"ipAddress,omitempty"`
	Browser          string `json:"browser,omitempty"`
	MobilePhoneModel string `json:"mobilePhoneModel,omitempty"`
}

type Response struct {
	GatewayCode string `json:"gatewayCode,omitempty"`
}

type BrowserPaymentResponse struct {
	RedirectURL string `json:"redirectUrl,omitempty"`
}

type Error struct {
	Cause          string `json:"cause,omitempty"`
	Explanation    string `json:"explanation,omitempty"`
	Field          string `json:"field,omitempty"`
	SupportCode    string `json:"supportCode,omitempty"`
	ValidationType string `json:"validationType,omitempty"`
}

type CreateRegistrationTokenRequest struct {
	MerchantID       string                `uri:"merchant_id,omitempty"`
	MerchantTokenRef string                `uri:"merch_token_ref,omitempty"`
	APIOperation     string                `json:"apiOperation,omitempty"`
	CorrelationID    string                `json:"correlationId,omitempty"`
	BrowserPayment   BrowserPaymentRequest `json:"browserPayment,omitempty"`
	Customer         Customer              `json:"customer,omitempty"`
	SourceOfFunds    SourceOfFunds         `json:"sourceOfFunds,omitempty"`
	Device           Device                `json:"device,omitempty"`
}

type CreateRegistrationTokenResponse struct {
	CorrelationID  string                 `json:"correlationId,omitempty"`
	Merchant       string                 `json:"merchant,omitempty"`
	Response       Response               `json:"response,omitempty"`
	BrowserPayment BrowserPaymentResponse `json:"browserPayment,omitempty"`
	Result         string                 `json:"result,omitempty"`
	Error          Error                  `json:"error,omitempty"`
}

type GetRegisterTokenRequest struct {
	MerchantId     string `uri:"merchant_id,omitempty"`
	RegistrationId string `uri:"registration_id,omitempty"`
}

type CustomerGetToken struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
}

type DirectDebitSacombank struct {
	AccountType       string `json:"accountType,omitempty"`
	BankAccountHolder string `json:"bankAccountHolder,omitempty"`
	BankAccountNumber string `json:"bankAccountNumber,omitempty"`
}

type Provided struct {
	DirectDebitSacombank DirectDebitSacombank `json:"directDebitSacombank,omitempty"`
}

type SourceOfFundsGetToken struct {
	Type     string   `json:"type,omitempty"`
	Provided Provided `json:"provided,omitempty"`
}

type GetBankData struct {
	Token             string `json:"token"`
	BankTransactionId string `json:"bank_transaction_id"`
}

type GetRegisterTokenResponse struct {
	GetBankData
	BankId        string                `json:"bank_id"`
	Status        string                `json:"status"`
	Customer      CustomerGetToken      `json:"customer"`
	SourceOfFunds SourceOfFundsGetToken `json:"sourceOfFunds"`
}
