package aurora

import "time"

// ProcessRequest represents a rule processing request.
type ProcessRequest struct {
	RuleID        string       `json:"rule_id,omitempty"`
	RulesetID     string       `json:"ruleset_id,omitempty"`
	Transaction   *Transaction `json:"transaction,omitempty"`
	TransactionID string       `json:"transaction_id,omitempty"`
}

// Transaction contains all the data fields that can be evaluated by rules.
type Transaction struct {
	// Core transaction fields
	Amount      float64 `json:"amount,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Category    string  `json:"category,omitempty"`
	Date        string  `json:"date,omitempty"`        // format: 2006-02-01 15:04:05
	Description string  `json:"description,omitempty"` // description from payment gateway

	// Payment gateway fields
	PaymentID     string `json:"payment_id,omitempty"`
	PaymentMethod string `json:"payment_method,omitempty"` // cash_pickup, bank_deposit, mobile_wallet
	ReferenceID   string `json:"reference_id,omitempty"`
	MerchantID    string `json:"merchant_id,omitempty"`
	MCC           string `json:"mcc,omitempty"`

	// Card fields
	BinNumber    string `json:"bin_number,omitempty"`
	CardHolder   string `json:"card_holder,omitempty"` // card holder name
	CardIssuer   string `json:"card_issuer,omitempty"` // bank name
	CardLastFour string `json:"card_last_four,omitempty"`
	CardNetwork  string `json:"card_network,omitempty"` // visa, mastercard, american express, etc.
	CardToken    string `json:"card_token,omitempty"`   // hashed card bin non-reversible
	CardType     string `json:"card_type,omitempty"`    // debit, credit

	// User / customer fields
	UserID    string `json:"user_id,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`

	// Location fields
	Country   string `json:"country,omitempty"` // derived from IP
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`

	// Device / connection fields
	BrowserAgent   string `json:"browser_agent,omitempty"`
	ConnectionType string `json:"connection_type,omitempty"` // wifi, mobile, ethernet, dsl, cable, etc.

	// Login history
	LastLoginIP   string `json:"last_login_ip,omitempty"`
	LastLoginTime string `json:"last_login_time,omitempty"`

	// Customer profile fields
	RegisteredIncome string `json:"registered_income,omitempty"`
	IncomeMultiplier string `json:"income_multiplier,omitempty"` // ratio of transaction amount to registered income

	// International money transfer fields
	TransferType     string `json:"transfer_type,omitempty"`
	TransferPurpose  string `json:"transfer_purpose,omitempty"` // family_support, education, business, medical, gift, other
	SourceOfFunds    string `json:"source_of_funds,omitempty"`  // salary, savings, business_income, investments, other
	ExchangeRate     string `json:"exchange_rate,omitempty"`
	TargetCurrency   string `json:"target_currency,omitempty"`
	Relationship     string `json:"relationship,omitempty"` // relationship to beneficiary
	IsFirstTransfer  string `json:"is_first_transfer,omitempty"`
	TotalAmount24h   string `json:"total_amount_24h,omitempty"`
	TransferCount24h string `json:"transfer_count_24h,omitempty"`

	// Sender fields
	SenderName             string `json:"sender_name,omitempty"`
	SenderSurname          string `json:"sender_surname,omitempty"`
	SenderCountry          string `json:"sender_country,omitempty"`
	SenderAddress          string `json:"sender_address,omitempty"`
	SenderIBAN             string `json:"sender_iban,omitempty"`
	SenderIdentityPassport string `json:"sender_identity_passport,omitempty"`
	SenderWalletID         string `json:"sender_wallet_id,omitempty"`

	// Receiver fields
	ReceiverName             string `json:"receiver_name,omitempty"`
	ReceiverSurname          string `json:"receiver_surname,omitempty"`
	ReceiverCountry          string `json:"receiver_country,omitempty"`
	ReceiverAddress          string `json:"receiver_address,omitempty"`
	ReceiverIBAN             string `json:"receiver_iban,omitempty"`
	ReceiverIdentityPassport string `json:"receiver_identity_passport,omitempty"`
	ReceiverWalletID         string `json:"receiver_wallet_id,omitempty"`
}

// ProcessResponse represents the result of rule processing.
type ProcessResponse struct {
	Allow          bool          `json:"allow"`
	AllowMessage   string        `json:"allow_message,omitempty"`
	Error          bool          `json:"error"`
	ErrorMessage   string        `json:"error_message,omitempty"`
	ExecutionTime  time.Duration `json:"execution_time"`
	InspectMessage string        `json:"inspect_message,omitempty"`
	NeedInspect    bool          `json:"need_inspect"`
	RejectMessage  string        `json:"reject_message,omitempty"`
	Rejected       bool          `json:"rejected"`
	Score          int           `json:"score"`
}

// ErrorResponse represents an API error.
type ErrorResponse struct {
	Error string `json:"error"`
}
