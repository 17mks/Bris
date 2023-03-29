package model

import (
	v1 "followup/api"
	"followup/gencode/chdb"
)

type FormInfo struct {
	chdb.TbForm
	TbFormCSS       chdb.TbFormCSS   `gorm:"foreignKey:tb_form_id"`
	FormColumnInfos []FormColumnInfo `gorm:"foreignKey:tb_form_id"`
}

type FormColumnInfo struct {
	chdb.TbFormColumn
	TbColumnOpts       []chdb.TbColumnOpt       `gorm:"foreignKey:tb_form_column_id"`
	TbColumnThresholds []chdb.TbColumnThreshold `gorm:"foreignKey:tb_form_column_id"`
}

type FormImportApiReq struct {
	HeaderParams
	Data v1.FormImportRequest
}
type FormExportApiReq struct {
	HeaderParams
	Data v1.FormExportRequest
}

//type TokenInfo struct {
//	UserId string    `json:"UserId"` // 用户ID
//	UserName string    `json:"UserName"` // 用户姓名
//	AppId    string    `json:"AppId"`    // AppId
//	ExTime time.Time `json:"exTime"` // 过期时间
//}
//
//// HeaderParams API会用到的所有请求头信息
//type HeaderParams struct {
//	UserToken string    `json:"UserToken,omitempty" form:"UserToken"` // Cli发送的加密TOKEN
//	CliType   string    `json:"CliType,omitempty" form:"CliType"`     //
//	TokenInfo TokenInfo `json:"-"`                                    // Token解密后的数据
//}

type FormDetailPreloadInfo struct {
	chdb.TbForm                       // 表单基本信息
	TbFormColumns []chdb.TbFormColumn `gorm:"foreignKey:tb_form_id"` // 表单列信息
	TbFormCSS     chdb.TbFormCSS      `gorm:"foreignKey:tb_form_id"` // 表单样式
}
