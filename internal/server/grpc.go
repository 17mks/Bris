package server

import (
	v1 "followup/api"
	"followup/internal/conf"
	"followup/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, plan *service.PlanService,
	authtoken *service.AuthTokenService, disease *service.DiseaseService,
	article *service.ArticleService,
	jwtc *conf.Data, logger log.Logger,
	disFunction *service.DisFunctionService,
	followup *service.FollowupService,
	form *service.FormService,
	formRow *service.FormRowService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			metadata.Server(),
			recovery.Recovery(),
			//selector.Server(auth.JWTAuth(jwtc.Secret)).Match(NewWhiteListMatcher()).Build(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterPlanServer(srv, plan)
	v1.RegisterArticleServer(srv, article)
	v1.RegisterAuthTokenServer(srv, authtoken)
	v1.RegisterDiseaseServer(srv, disease)
	v1.RegisterDisFunctionServer(srv, disFunction)
	v1.RegisterFollowupServer(srv, followup)
	v1.RegisterFormServer(srv, form)
	v1.RegisterFormRowServiceServer(srv, formRow)
	return srv
}
