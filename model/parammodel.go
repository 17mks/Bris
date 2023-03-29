package model

import "time"

type TokenInfo struct {
	UserId   string    `json:"UserId"`   // 用户ID
	UserName string    `json:"UserName"` // 用户姓名
	AppId    string    `json:"AppId"`    // AppId
	ExTime   time.Time `json:"exTime"`   // 过期时间
}

// PathParams API会用到的所有路径参数
type PathParams struct {
	BodyType string `json:"bodyType" form:"bodyType"` //
	BodyId   string `json:"bodyId" form:"bodyId"`     //
	Field    string `json:"field" form:"field"`       //
	CohortId string `json:"cohortId" form:"cohortId"` //
}

// QueryParams API会用到的所有查询参数
type QueryParams struct {
	Id string `json:"id,omitempty" form:"id"` // 资源主键编码
	//PriId    int64  `json:"priId,omitempty" form:"priId"` // 资源主键编码(雪花)
	Page      int    `json:"page,omitempty" form:"page"`
	PageSize  int    `json:"perPage,omitempty" form:"perPage"`
	Account   string `json:"account,omitempty" form:"account"`     // 资源主键编码
	AuthCode  string `json:"authCode,omitempty" form:"authCode"`   // 短信验证码
	Token     string `json:"UserToken,omitempty" form:"UserToken"` // 用户Token
	BodyType  string `json:"bodyType,omitempty" form:"bodyType"`   // 主体类型、成员类型
	BodyId    string `json:"bodyId,omitempty" form:"bodyId"`       // 主体编码
	MemberId  string `json:"memberId,omitempty" form:"memberId"`   // 成员编码
	JsCode    string `json:"js_code,omitempty" form:"js_code"`     // js_code
	UserId    string `json:"userId,omitempty" form:"userId"`       // userId
	AppId     string `json:"appId,omitempty" form:"appId"`         // userId
	CohortId  string `json:"cohortId,omitempty" form:"cohortId"`   // 队列编码
	ProjectId string `json:"projectId,omitempty" form:"projectId"` // 项目编码
}

// HeaderParams API会用到的所有请求头信息
type HeaderParams struct {
	UserToken   string    `json:"UserToken,omitempty" form:"UserToken"`     // Cli发送的加密TOKEN
	UserId      string    `json:"UserId,omitempty" form:"UserId"`           // 用户ID
	AppId       string    `json:"AppId,omitempty" form:"AppId"`             // 应用编码
	AccessToken string    `json:"AccessToken,omitempty" form:"AccessToken"` // AccessToken(包含AppId、AppKey)
	CliType     string    `json:"CliType,omitempty" form:"CliType"`         //
	TokenInfo   TokenInfo `json:"-"`                                        // Token解密后的数据
	Tag         string    `json:"tag,omitempty" form:"tag"`                 // 校验位，如果是指定的TAG则不做TOKEN校验【APP_UPGRADE_DOWNLOAD】
}

func (HeaderParams *HeaderParams) GetUserId() string {
	return ""
}
