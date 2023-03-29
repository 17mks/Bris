package service

import (
	"context"
	"errors"
	"followup/api/models"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"gogs.buffalo-robot.com/gogs/auth"
	"gogs.buffalo-robot.com/gogs/module/net/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

type BasisService interface {
	JwtParseToken(token string) (*http.JWTclaims, error) // 解析TOKEN
	ParseContextHeader(ctx context.Context) (*models.HeaderParams, error)
}

type SBasisService struct {
}

func NewSBasisService() BasisService {
	return &SBasisService{}
}

// MAuthStore 为了解析TOKEN实现basis中的接口
type MAuthStore struct {
}

func (authStore *MAuthStore) Exists(ctx context.Context, token string) bool {
	// 直接返回true
	return true
}

const (
	HeaderToken = "UserToken"
)

// ParseContextHeader 解析Context中传递的元数据(需区分HTTP、GRPC两种情况)
func (service *SBasisService) ParseContextHeader(ctx context.Context) (*models.HeaderParams, error) {
	headerMap := make(map[string]string)
	// HTTP
	if serverContext, ok := transport.FromServerContext(ctx); ok {
		val := serverContext.RequestHeader().Get(HeaderToken)
		if "" != val {
			headerMap[HeaderToken] = val
		}
	}
	// GRPC
	if serverContext, ok := metadata.FromServerContext(ctx); ok {
		val := serverContext.Get(HeaderToken)
		if "" != val {
			headerMap[HeaderToken] = val
		}
	}

	headerParams := models.HeaderParams{
		Token:     headerMap[HeaderToken],
		TokenInfo: nil,
	}
	// 解析TOKEN
	if "" != headerParams.Token {
		tokenInfo, err := service.JwtParseToken(headerParams.Token)
		if err != nil {
			return nil, err
		}
		headerParams.TokenInfo = &models.TokenInfo{
			UserId:   strconv.FormatUint(uint64(tokenInfo.ID), 10),
			UserName: tokenInfo.UserName,
			Email:    tokenInfo.Email,
			Mobile:   tokenInfo.Phone,
		}
	}

	return &headerParams, nil
}

// JwtParseToken 调用JWT方式解析TOKEN
func (service *SBasisService) JwtParseToken(token string) (*http.JWTclaims, error) {
	authRepo := auth.NewAuthRepo(&MAuthStore{})
	checkToken, err := authRepo.CheckToken(context.Background(), token)
	if err != nil {
		return nil, err
	}
	if nil == checkToken {
		return nil, errors.New("无效TOKEN")
	}
	return checkToken, nil
}

// JwtParseToken 调用JWT方式解析TOKEN
func JwtParseToken(token string) (*http.JWTclaims, error) {
	authRepo := auth.NewAuthRepo(&MAuthStore{})
	checkToken, err := authRepo.CheckToken(context.Background(), token)
	if err != nil {
		return nil, err
	}
	log.Println("id = ", checkToken.ID)
	return checkToken, nil
}

// GetGrpcCliConn 获取GRPC连接
// grpcTarget：服务地址 e.g. 192.168.100.20:8701
func GetGrpcCliConn(grpcTarget string) (*grpc.ClientConn, error) {
	var err error
	grpcCliConn, err := grpc.Dial(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

	}

	return grpcCliConn, nil

}
