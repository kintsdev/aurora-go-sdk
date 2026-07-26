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
