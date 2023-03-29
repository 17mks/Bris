package server

import (
	"context"
	"fmt"
	v1 "followup/api"
	"followup/internal/biz"
	"followup/internal/conf"
	"followup/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	netHttp "net/http"
)

func NewWhiteListMatcher(whiteList map[string]bool) selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, configData *conf.Data,
	plan *service.PlanService, authtoken *service.AuthTokenService,
	disease *service.DiseaseService,
	dataRepo biz.DataRepo,
	article *service.ArticleService,
	disFunction *service.DisFunctionService,
	form *service.FormService,
	followup *service.FollowupService,
	filesService *service.FileService,
	formRow *service.FormRowService,
	logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		//自定义错误返回
		http.ErrorEncoder(errorEncoder),
		//自定义返回数据结构
		http.ResponseEncoder(responseEncoder),
		http.Middleware(
			metadata.Server(),
			recovery.Recovery(),
			//selector.Server(auth.JWTAuth(jwtc.Secret)).Match(NewSkipRoutersMatcher()).Build(),
			logging.Server(logger),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "U-Token", "Accept-Encoding", "Host", "Connection", "Referer", "User-Agent"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// 表单模板静态资源
	formModelPath := configData.FileService.UnZipPath
	engine := gin.Default()
	//engine.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware)) // 【TODO 此处不加customMiddleware也能正常执行，稍后研究】
	engine.StaticFS("/fs/form/model", netHttp.Dir(formModelPath))
	srv.HandlePrefix("/fs", engine)

	v1.RegisterPlanHTTPServer(srv, plan)
	v1.RegisterAuthTokenHTTPServer(srv, authtoken)
	v1.RegisterDiseaseHTTPServer(srv, disease)
	v1.RegisterArticleHTTPServer(srv, article)
	v1.RegisterDisFunctionHTTPServer(srv, disFunction)
	v1.RegisterFormHTTPServer(srv, form)
	v1.RegisterFollowupHTTPServer(srv, followup)
	v1.RegisterFormRowServiceHTTPServer(srv, formRow)
	return srv
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}
