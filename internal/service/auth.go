package service

import (
	"context"
	v1 "followup/api"
)

func (s *AuthTokenService) CreateAuthToken(ctx context.Context, req *v1.CreateAuthTokenRequest) (*v1.CreateAuthTokenResponse, error) {
	au, err := s.Auth.CreateToken(ctx, req.UserName, req.UserId, req.AppId)
	if err != nil {
		return nil, err
	}
	return &v1.CreateAuthTokenResponse{
		Data: &v1.CreateAuthTokenResponse_Data{
			Token: au.Token,
		},
	}, nil
}
