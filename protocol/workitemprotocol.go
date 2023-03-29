package protocol

const (
	WorkItemTypeSymptomAssessment = "SA"        // 工作项类型 - 症状评估
	WorkItemTypeTreat             = "TREAT"     // 工作项类型 - 治疗任务
	WorkItemTypeFollowUp          = "FOLLOW_UP" // 工作项类型 - 随访
	WorkItemTypeTask              = "TASK"      // 工作项类型 - 任务

	WorkItemStatusModel    = "MODEL"    // 工作项状态 - 模型
	WorkItemStatusDraft    = "DRAFT"    // 工作项状态 - 草稿,对用户不可见
	WorkItemStatusNew      = "NEW"      // 工作项状态 - 新建/待处理
	WorkItemStatusActive   = "ACTIVE"   // 工作项状态 - 进行中
	WorkItemStatusResolved = "RESOLVED" // 工作项状态 - 已解决For Bug
	WorkItemStatusClosed   = "CLOSED"   // 工作项状态 - 已完成/已关闭
	WorkItemStatusRemoved  = "REMOVED"  // 工作项状态 - 已移除

	WorkItemAssignedTypeUser    = "USER"    // 工作项指派类型 - 指派给用户
	WorkItemAssignedTypeMember  = "MEMBER"  // 工作项指派类型 - 指派给成员
	WorkItemAssignedTypePatient = "PATIENT" // 工作项指派类型 - 指派给患者
	WorkItemAssignedTypeNone    = "NONE"    // 工作项指派类型 - 未指派

	WorkItemPrincipalTypeUser   = "USER"   // 工作项负责人类型 - 指派给用户
	WorkItemPrincipalTypeMember = "MEMBER" // 工作项指派负责人类型 - 指派给成员
	WorkItemPrincipalTypeNone   = "NONE"   // 工作项指派负责人类型 - 未指派

	WorkItemBelongTypeProject = "PROJECT" // 工作项归属 - 项目
	WorkItemBelongTypeNone    = "NONE"    // 工作项归属 - 无

)
