package protocol

import (
	"github.com/google/uuid"
	"strings"
)

const (
	BodyTypeUser       = "USER"       // 主体类型 - 用户
	BodyTypeOrg        = "ORG"        // 主体类型 - 组织
	BodyTypeDepartment = "DEPARTMENT" // 主体类型 - 部门
	BodyTypeProject    = "PROJECT"    // 主体类型 - 项目
	BodyTypeTeam       = "TEAM"       // 主体类型 - 团队
	BodyTypeMember     = "MEMBER"     // 主体类型 - 成员
	BodyTypeDoctor     = "DOCTOR"     // 主体类型 - 医生
	BodyTypePatient    = "PATIENT"    // 主体类型 - 患者
	BodyTypeRole       = "ROLE"       // 主体类型 - 角色
	BodyTypePlan       = "PLAN"       // 主体类型 - 方案
	BodyTypeCohort     = "COHORT"     // 主体类型 - 队列
	BodyTypeNone       = "NONE"       // 主体类型 - 空
)

const (
	ResourceTypeOrg         = "ORG"         // 资源类型 - 组织
	ResourceTypeProject     = "PROJECT"     // 资源类型 - 项目
	ResourceTypeTeam        = "TEAM"        // 资源类型 - 团队
	ResourceTypeRole        = "ROLE"        // 资源类型 - 角色
	ResourceTypePatient     = "PATIENT"     // 资源类型 - 患者
	ResourceTypePermission  = "PERMISSION"  // 资源类型 - 权限
	ResourceTypePlan        = "PLAN"        // 资源类型 - 方案
	ResourceTypeCertificate = "CERTIFICATE" // 资源类型 - 证件
	ResourceTypeMenu        = "MENU"        // 资源类型 - 菜单
	ResourceTypeMember      = "MEMBER"      // 主体类型 - 成员
	ResourceTypeCohort      = "COHORT"      // 主体类型 - 队列
	ResourceTypeForm        = "FORM"        // 主体类型 - 表单
)

// 'FORM','WI','BZ','GNZA'
const (
	PlanRelateResTypeForm       = "FORM" // 方案关联资源类型 - 表单
	PlanRelateResTypeWorkItem   = "WI"   // 方案关联资源类型 - 工作项
	PlanRelateResTypeDisease    = "BZ"   // 方案关联资源类型 - 病种
	PlanRelateResTypDysfunction = "GNZA" // 方案关联资源类型 - 功能障碍
)

const (
	PlanTypeSA       = "SA"       // 方案类型 - 症状评估方案
	PlanTypeFollowUp = "FOLLOWUP" // 方案类型 - 随访方案
)
const (
	PlanStatusDraft    = "DRAFT"    // 方案状态 - 草稿
	PlanStatusEnabled  = "ENABLED"  // 方案状态 - 启用
	PlanStatusDisabled = "DISABLED" // 方案状态 - 禁用
)
const (
	PlanBelongTypeNone    = "NONE"    // 方案归属 - 无
	PlanBelongTypeProject = "PROJECT" // 方案归属 - 项目
)
const (
	MemberTypeORG        = "ORG"        // 组织成员
	MemberTypeDepartment = "DEPARTMENT" // 部门成员
	MemberTypeProject    = "PROJECT"    // 组织成员
	MemberTypeCohort     = "COHORT"     // 队列成员
	MemberTypeTeam       = "TEAM"       // 团队成员
)

var (
	StatusSuccess = StatusCode{
		Code:    SuccessCode,
		Message: "成功",
	}
)

// StatusCode 状态码
type StatusCode struct {
	// 状态码
	Code int32 `json:"code"`
	// 状态描述
	Message string `json:"message"`
}

// 状态Code
const (
	SuccessCode               = iota // 成功
	BusinessErrCode                  // 业务异常
	DatabaseErrCode                  // 数据库异常
	CacheErrCode                     // 缓存异常
	BadRequestErrCode                // 请求数据异常
	AccNotExistsErrCode              // 账号未注册异常
	AccRepeatErrCode                 // 账号已存在异常
	ResourceNotExistsErrCode         // 资源不存在异常
	TokenExpiredErrCode              // TOKEN过期异常
	AccOrPasswordErrCode             // 账号或密码错误异常
	PasswordNotSetErrCode            // 账号未设置密码异常
	PermissionDeniedErrCode          // 权限受限
	CardHadBindByOtherErrCode        // 证件已被其他人绑定
	AccountLockedErrCode             // 账号已锁定
	ResourceHasExistsErrCode         // 资源已存在
	ReJoinCohortErrCode              // 重复入组
)

// NewDefaultResponse 创建默认的响应数据
func NewDefaultResponse() Response {
	return Response{
		StatusCode: StatusSuccess,
		RequestId:  GenNewRequestId(),
	}
}

// Response API响应数据结构
type Response struct {
	// 业务状态码
	StatusCode
	// 详细错误信息，用于开发调试问题定位
	Err string `json:"err,omitempty"`
	// 请求编码,用于查询本次请求相关信息
	RequestId string `json:"requestId"`
	// API响应数据
	Data interface{} `json:"data,omitempty"`
}

func GenNewRequestId() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToLower(userId)
}
