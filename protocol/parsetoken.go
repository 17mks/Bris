package protocol

//func ParseToken(ctx context.Context) (*utils.JwtCustClaims, error) {
//	serverContext, ok := transport.FromServerContext(ctx)
//	if !ok {
//		return nil, fmt.Errorf("解析Context获取TOKEN失败")
//	}
//	rv := serverContext.RequestHeader().Get("Usertoken")
//	tokenInfo, err := utils.ParseToken(rv)
//	if err != nil {
//		return nil, err
//	}
//	return tokenInfo, nil
//}
