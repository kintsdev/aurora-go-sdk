package aurora

import (
	"context"
	"net/http"
)

// ProcessService handles rule processing operations.
type ProcessService struct {
	client *Client
}

// Execute processes a transaction against a rule or ruleset and returns the result.
func (s *ProcessService) Execute(ctx context.Context, req *ProcessRequest) (*ProcessResponse, error) {
	if req.TransactionID == "" && req.Transaction == nil {
		return nil, &ValidationError{Field: "transaction", Message: "transaction is required"}
	}
	if !validTransactionTypes[req.Transaction.Type] {
		return nil, &ValidationError{Field: "transaction.type", Message: "invalid or missing transaction type"}
	}

	httpReq, err := s.client.newRequest(ctx, http.MethodPost, "/process", req)
	if err != nil {
		return nil, err
	}

	var resp ProcessResponse
	if err := s.client.do(httpReq, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
