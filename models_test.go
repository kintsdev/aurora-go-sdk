package aurora

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestTransactionCustomerAgeYearsJSON(t *testing.T) {
	payload, err := json.Marshal(Transaction{Type: TransactionTypeTransfer, CustomerAgeYears: "29"})
	if err != nil {
		t.Fatalf("marshal transaction: %v", err)
	}

	if !strings.Contains(string(payload), `"customer_age_years":"29"`) {
		t.Fatalf("expected customer_age_years in payload, got %s", payload)
	}
}

func TestTransactionCustomerProfileRiskFieldsJSON(t *testing.T) {
	payload, err := json.Marshal(Transaction{
		Type:                    TransactionTypeTransfer,
		CustomerAccountType:     CustomerAccountTypeIndividual,
		FirstTransactionAgeDays: "30",
	})
	if err != nil {
		t.Fatalf("marshal transaction: %v", err)
	}

	for _, expected := range []string{
		`"customer_account_type":"individual"`,
		`"first_transaction_age_days":"30"`,
	} {
		if !strings.Contains(string(payload), expected) {
			t.Errorf("expected %s in payload, got %s", expected, payload)
		}
	}
}

func TestProcessRequestLifecycleRootTransactionIDJSON(t *testing.T) {
	payload, err := json.Marshal(ProcessRequest{
		RuleID:                     "1e7b3b7b-0b3b-4b7b-8b3b-0b3b7b0b3b7b",
		LifecycleRootTransactionID: "9c2f5a41-6b0e-4c8a-9f31-2d7b8e4a1c05",
		Transaction:                &Transaction{Type: TransactionTypeRemittance},
	})
	if err != nil {
		t.Fatalf("marshal process request: %v", err)
	}

	if !strings.Contains(string(payload), `"lifecycle_root_transaction_id":"9c2f5a41-6b0e-4c8a-9f31-2d7b8e4a1c05"`) {
		t.Fatalf("expected lifecycle_root_transaction_id in payload, got %s", payload)
	}
}

func TestProcessRequestLifecycleRootTransactionIDOmittedJSON(t *testing.T) {
	payload, err := json.Marshal(ProcessRequest{
		RuleID:      "1e7b3b7b-0b3b-4b7b-8b3b-0b3b7b0b3b7b",
		Transaction: &Transaction{Type: TransactionTypeRemittance},
	})
	if err != nil {
		t.Fatalf("marshal process request: %v", err)
	}

	if strings.Contains(string(payload), "lifecycle_root_transaction_id") {
		t.Fatalf("expected lifecycle_root_transaction_id to be omitted, got %s", payload)
	}
}

func TestTransferSubtypeConstants(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  string
		want string
	}{
		{"TransferTypeWalletToWallet", TransferTypeWalletToWallet, "wallet_to_wallet"},
		{"TransferTypeWalletToIBAN", TransferTypeWalletToIBAN, "wallet_to_iban"},
		{"TransferTypeIBANToWallet", TransferTypeIBANToWallet, "iban_to_wallet"},
		{"TransferTypeRemittance", TransferTypeRemittance, "remittance"},
	} {
		if tc.got != tc.want {
			t.Errorf("expected %s to be %q, got %q", tc.name, tc.want, tc.got)
		}
	}
}

func TestWithdrawalMethodConstants(t *testing.T) {
	if WithdrawalMethodMoneyTransferCollection != "money_transfer_collection" {
		t.Errorf("expected WithdrawalMethodMoneyTransferCollection to be %q, got %q",
			"money_transfer_collection", WithdrawalMethodMoneyTransferCollection)
	}
}

func TestTransactionClientIPLocationJSON(t *testing.T) {
	payload, err := json.Marshal(Transaction{
		Type:      TransactionTypeTransfer,
		IPCountry: "AF",
		IPCity:    "Adana",
		Common: &Common{
			LocationCountry:    "AF",
			LocationRegion:     "Adana",
			LocationRegionCode: "01",
			LocationCity:       "Adana",
			LocationSource:     "ip.tapsilat.dev",
			Latitude:           "37.0000",
			Longitude:          "35.3213",
		},
	})
	if err != nil {
		t.Fatalf("marshal transaction: %v", err)
	}

	for _, expected := range []string{
		`"ip_country":"AF"`,
		`"ip_city":"Adana"`,
		`"location_country":"AF"`,
		`"location_region":"Adana"`,
		`"location_region_code":"01"`,
		`"location_city":"Adana"`,
		`"location_source":"ip.tapsilat.dev"`,
		`"latitude":"37.0000"`,
		`"longitude":"35.3213"`,
	} {
		if !strings.Contains(string(payload), expected) {
			t.Errorf("expected %s in payload, got %s", expected, payload)
		}
	}
}
