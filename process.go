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
	httpReq, err := s.client.newRequest(ctx, http.MethodPost, "/process/", req)
	if err != nil {
		return nil, err
	}

	var resp ProcessResponse
	if err := s.client.do(httpReq, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
