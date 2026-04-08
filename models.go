package aurora

import "time"

// Transaction type constants.
const (
	TransactionTypeCard           = "card"
	TransactionTypeTransfer       = "transfer"
	TransactionTypeRemittance     = "remittance"
	TransactionTypeAccountOpening = "account_opening"
	TransactionTypeAccountLogin   = "account_login"
	TransactionTypeAccountChange  = "account_change"
	TransactionTypeRefund         = "refund"
	TransactionTypeWithdrawal     = "withdrawal"
	TransactionTypeDeposit        = "deposit"
	TransactionTypeP2P            = "p2p"
	TransactionTypeCrypto         = "crypto"
	TransactionTypeBNPL           = "bnpl"
	TransactionTypeLoan           = "loan"
	TransactionTypeInvoice        = "invoice"
)

// validTransactionTypes is the set of accepted transaction types.
var validTransactionTypes = map[string]bool{
	TransactionTypeCard:           true,
	TransactionTypeTransfer:       true,
	TransactionTypeRemittance:     true,
	TransactionTypeAccountOpening: true,
	TransactionTypeAccountLogin:   true,
	TransactionTypeAccountChange:  true,
	TransactionTypeRefund:         true,
	TransactionTypeWithdrawal:     true,
	TransactionTypeDeposit:        true,
	TransactionTypeP2P:            true,
	TransactionTypeCrypto:         true,
	TransactionTypeBNPL:           true,
	TransactionTypeLoan:           true,
	TransactionTypeInvoice:        true,
}

// ProcessRequest represents a rule processing request.
type ProcessRequest struct {
	RuleID        string       `json:"rule_id,omitempty"`
	RulesetID     string       `json:"ruleset_id,omitempty"`
	Transaction   *Transaction `json:"transaction,omitempty"`
	TransactionID string       `json:"transaction_id,omitempty"`
}

// Transaction contains all the data fields that can be evaluated by rules.
type Transaction struct {
	// Nested sub-objects
	AccountChange  *AccountChange  `json:"account_change,omitempty"`
	AccountLogin   *AccountLogin   `json:"account_login,omitempty"`
	AccountOpening *AccountOpening `json:"account_opening,omitempty"`
	BNPL           *BNPL           `json:"bnpl,omitempty"`
	Card           *Card           `json:"card,omitempty"`
	Common         *Common         `json:"common,omitempty"`
	Crypto         *Crypto         `json:"crypto,omitempty"`
	Deposit        *Deposit        `json:"deposit,omitempty"`
	Invoice        *Invoice        `json:"invoice,omitempty"`
	Loan           *Loan           `json:"loan,omitempty"`
	P2P            *P2P            `json:"p2p,omitempty"`
	Refund         *Refund         `json:"refund,omitempty"`
	Transfer       *Transfer       `json:"transfer,omitempty"`
	Withdrawal     *Withdrawal     `json:"withdrawal,omitempty"`

	// Top-level scalar fields
	AccountAgeDays          string `json:"account_age_days,omitempty"`
	DeclinedCount           string `json:"declined_count,omitempty"`
	IncomeMultiplier        string `json:"income_multiplier,omitempty"`
	IsFirstTransfer         string `json:"is_first_transfer,omitempty"`
	IsNewDevice             string `json:"is_new_device,omitempty"`
	IsNewIP                 string `json:"is_new_ip,omitempty"`
	IsUnusualLocation       string `json:"is_unusual_location,omitempty"`
	PasswordChangedRecently string `json:"password_changed_recently,omitempty"`
	ProfileCompletion       string `json:"profile_completion,omitempty"`
	RefundCount             string `json:"refund_count,omitempty"`
	RefundRatio             string `json:"refund_ratio,omitempty"`
	RegisteredIncome        string `json:"registered_income,omitempty"`
	TotalAmount24h          string `json:"total_amount_24h,omitempty"`
	TotalAmount7d           string `json:"total_amount_7d,omitempty"`
	TransactionHour         string `json:"transaction_hour,omitempty"`
	TransactionType         string `json:"transaction_type,omitempty"`
	TransferCount24h        string `json:"transfer_count_24h,omitempty"`
	Type                    string `json:"type"`
	UniqueRecipients        string `json:"unique_recipients,omitempty"`
}

// Common holds core transaction and user identification fields.
type Common struct {
	Amount            float64 `json:"amount,omitempty"`
	BrowserAgent      string  `json:"browser_agent,omitempty"`
	Category          string  `json:"category,omitempty"`
	ConnectionType    string  `json:"connection_type,omitempty"`
	Country           string  `json:"country,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	Date              string  `json:"date,omitempty"`
	Description       string  `json:"description,omitempty"`
	DeviceFingerprint string  `json:"device_fingerprint,omitempty"`
	DeviceID          string  `json:"device_id,omitempty"`
	Email             string  `json:"email,omitempty"`
	IPAddress         string  `json:"ip_address,omitempty"`
	LastLoginIP       string  `json:"last_login_ip,omitempty"`
	LastLoginTime     string  `json:"last_login_time,omitempty"`
	Latitude          string  `json:"latitude,omitempty"`
	Longitude         string  `json:"longitude,omitempty"`
	MerchantID        string  `json:"merchant_id,omitempty"`
	PaymentID         string  `json:"payment_id,omitempty"`
	Phone             string  `json:"phone,omitempty"`
	ReferenceID       string  `json:"reference_id,omitempty"`
	UserID            string  `json:"user_id,omitempty"`
}

// Card holds payment card details.
type Card struct {
	BinNumber string `json:"bin_number,omitempty"`
	Holder    string `json:"holder,omitempty"`
	Issuer    string `json:"issuer,omitempty"`
	LastFour  string `json:"last_four,omitempty"`
	MCC       string `json:"mcc,omitempty"`
	Network   string `json:"network,omitempty"`
	Token     string `json:"token,omitempty"`
	Type      string `json:"type,omitempty"`
}

// Transfer holds international money transfer fields.
type Transfer struct {
	ExchangeRate             string `json:"exchange_rate,omitempty"`
	PaymentMethod            string `json:"payment_method,omitempty"`
	ReceiverAddress          string `json:"receiver_address,omitempty"`
	ReceiverCountry          string `json:"receiver_country,omitempty"`
	ReceiverIBAN             string `json:"receiver_iban,omitempty"`
	ReceiverIdentityPassport string `json:"receiver_identity_passport,omitempty"`
	ReceiverName             string `json:"receiver_name,omitempty"`
	ReceiverSurname          string `json:"receiver_surname,omitempty"`
	ReceiverWalletID         string `json:"receiver_wallet_id,omitempty"`
	Relationship             string `json:"relationship,omitempty"`
	SenderAddress            string `json:"sender_address,omitempty"`
	SenderCountry            string `json:"sender_country,omitempty"`
	SenderIBAN               string `json:"sender_iban,omitempty"`
	SenderIdentityPassport   string `json:"sender_identity_passport,omitempty"`
	SenderName               string `json:"sender_name,omitempty"`
	SenderSurname            string `json:"sender_surname,omitempty"`
	SenderWalletID           string `json:"sender_wallet_id,omitempty"`
	SourceOfFunds            string `json:"source_of_funds,omitempty"`
	TargetCurrency           string `json:"target_currency,omitempty"`
	TransferPurpose          string `json:"transfer_purpose,omitempty"`
	TransferType             string `json:"transfer_type,omitempty"`
}

// AccountChange holds account modification event data.
type AccountChange struct {
	ChangeType          string `json:"change_type,omitempty"`
	PreviousValueHash   string `json:"previous_value_hash,omitempty"`
	TimeSinceLastChange string `json:"time_since_last_change,omitempty"`
	VerificationMethod  string `json:"verification_method,omitempty"`
}

// AccountLogin holds login event data.
type AccountLogin struct {
	FailedAttempts string `json:"failed_attempts,omitempty"`
	LoginMethod    string `json:"login_method,omitempty"`
	LoginStatus    string `json:"login_status,omitempty"`
	MFAMethod      string `json:"mfa_method,omitempty"`
	MFAUsed        string `json:"mfa_used,omitempty"`
	SessionID      string `json:"session_id,omitempty"`
}

// AccountOpening holds account registration data.
type AccountOpening struct {
	DocumentType               string `json:"document_type,omitempty"`
	IdentityVerificationStatus string `json:"identity_verification_status,omitempty"`
	RegistrationMethod         string `json:"registration_method,omitempty"`
}

// BNPL holds buy-now-pay-later transaction data.
type BNPL struct {
	BNPLProvider          string `json:"bnpl_provider,omitempty"`
	InstallmentCount      string `json:"installment_count,omitempty"`
	OutstandingBNPLAmount string `json:"outstanding_bnpl_amount,omitempty"`
}

// Crypto holds cryptocurrency transaction data.
type Crypto struct {
	CryptoAmount    string `json:"crypto_amount,omitempty"`
	CryptoCurrency  string `json:"crypto_currency,omitempty"`
	ExchangeName    string `json:"exchange_name,omitempty"`
	IsSmartContract string `json:"is_smart_contract,omitempty"`
	WalletAddress   string `json:"wallet_address,omitempty"`
}

// Deposit holds deposit transaction data.
type Deposit struct {
	CheckNumber     string `json:"check_number,omitempty"`
	DepositMethod   string `json:"deposit_method,omitempty"`
	DepositSource   string `json:"deposit_source,omitempty"`
	IsRemoteDeposit string `json:"is_remote_deposit,omitempty"`
}

// Invoice holds invoice and vendor payment data.
type Invoice struct {
	InvoiceDueDate    string `json:"invoice_due_date,omitempty"`
	InvoiceNumber     string `json:"invoice_number,omitempty"`
	IsRecurringVendor string `json:"is_recurring_vendor,omitempty"`
	PurchaseOrderID   string `json:"purchase_order_id,omitempty"`
	VendorID          string `json:"vendor_id,omitempty"`
	VendorName        string `json:"vendor_name,omitempty"`
}

// Loan holds loan application data.
type Loan struct {
	CreditScore       string `json:"credit_score,omitempty"`
	DebtToIncomeRatio string `json:"debt_to_income_ratio,omitempty"`
	EmploymentStatus  string `json:"employment_status,omitempty"`
	LoanAmount        string `json:"loan_amount,omitempty"`
	LoanTerm          string `json:"loan_term,omitempty"`
	LoanType          string `json:"loan_type,omitempty"`
}

// P2P holds peer-to-peer payment data.
type P2P struct {
	PaymentNote         string `json:"payment_note,omitempty"`
	Platform            string `json:"platform,omitempty"`
	RecipientAccountAge string `json:"recipient_account_age,omitempty"`
}

// Refund holds refund transaction data.
type Refund struct {
	DaysSincePurchase     string `json:"days_since_purchase,omitempty"`
	OriginalTransactionID string `json:"original_transaction_id,omitempty"`
	RefundMethod          string `json:"refund_method,omitempty"`
	RefundReason          string `json:"refund_reason,omitempty"`
}

// Withdrawal holds withdrawal transaction data.
type Withdrawal struct {
	AccountType          string `json:"account_type,omitempty"`
	ATMID                string `json:"atm_id,omitempty"`
	DailyWithdrawalCount string `json:"daily_withdrawal_count,omitempty"`
	WithdrawalMethod     string `json:"withdrawal_method,omitempty"`
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
