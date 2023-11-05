package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-template-wire/configs"
	"go-template-wire/pkg/failure"
	httpclient "go-template-wire/pkg/http_client"
	"net/http"
)

type Service struct {
	Cfg        *configs.Config
	HTTPClient *httpclient.HTTPClient
}

func New(
	cfg *configs.Config, httpClient *httpclient.HTTPClient,
) *Service {
	return &Service{Cfg: cfg, HTTPClient: httpClient}
}

func (s *Service) GetUserInfo(ctx context.Context, userID string) (*User, error) {
	type request struct {
		ID string `json:"id"`
	}
	dto := request{ID: userID}
	reqBody, err := json.Marshal(&dto)
	if err != nil {
		return nil, failure.ErrWithTrace(fmt.Errorf("Failed to marshal body: %w", err))
	}

	req, err := httpclient.CreateRequest(ctx, &httpclient.CreateHTTPRequestDTO{
		Method: http.MethodPost,
		URL:    fmt.Sprintf("%s/ekyc/v1/internal/user/info", s.Cfg.Server.Host),
		Body:   bytes.NewBuffer(reqBody),
		AuthHeader: &httpclient.AuthHeader{
			Method: httpclient.AuthMethodBearerAPIKey,
			Token:  s.Cfg.APIKey.EKYCAPIKey,
		},
	})
	if err != nil {
		return nil, failure.ErrWithTrace(fmt.Errorf("Failed to create request: %w", err))
	}

	responseBody, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, failure.ErrWithTrace(fmt.Errorf("Failed to send request: %w", err))
	}
	res, err := httpclient.ParseResponseBody[User](responseBody)
	if err != nil {
		return nil, failure.ErrWithTrace(fmt.Errorf("User=%s | Failed to typecast response: %w", userID, err))
	}
	return &res.Data, nil
}
