package aurora

import (
	"context"
	"net/http"
)

// CallbackService handles callback operations.
type CallbackService struct {
	client *Client
}

// Transaction sends a callback transaction request.
func (s *CallbackService) Transaction(ctx context.Context, req *CallbackTransactionRequest) (*CallbackTransactionResponse, error) {
	if req.PaymentID == "" {
		return nil, &ValidationError{Field: "payment_id", Message: "payment_id is required"}
	}
	if req.Status == "" {
		return nil, &ValidationError{Field: "status", Message: "status is required"}
	}

	httpReq, err := s.client.newRequest(ctx, http.MethodPost, "/callback/transaction", req)
	if err != nil {
		return nil, err
	}

	var resp CallbackTransactionResponse
	if err := s.client.do(httpReq, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
