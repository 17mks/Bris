package utils

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	"log"
)

func ParseHeaher(ctx context.Context) {
	serverContext, ok := transport.FromServerContext(ctx)

	if !ok {
		fmt.Println("解析Context获取TOKEN失败")
	}

	rv := serverContext.RequestHeader().Get("Usertoken")
	log.Println(rv)
	Token, _ := ParseToken(rv)
	log.Println(Token)
}
