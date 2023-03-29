package openapi

type Plan struct {

	// 应用编码
	AppId string `json:"appId,omitempty"`

	// 适用年龄段(多个年龄段用','分隔)
	ApplyAges string `json:"applyAges,omitempty"`

	// 适用病种
	ApplyDisease string `json:"applyDisease,omitempty"`

	// 适用功能障碍(多个功能障碍用','分隔)
	ApplyDysfunction string `json:"applyDysfunction,omitempty"`

	// 规则资源编码(e.g. 如果归属类型时是项目则填项目编码)
	BelongTo string `json:"belongTo,omitempty"`

	// 归属类型(组织、项目、团队、组等)
	BelongType string `json:"belongType"`

	// 创建时间
	CreateTime string `json:"createTime,omitempty"`

	// 创建人编码
	CreatorId string `json:"creatorId,omitempty"`

	// 创建人编码
	CreatorName string `json:"creatorName,omitempty"`

	// 事件(随访开始触发事件)
	Event string `json:"event,omitempty"`

	// 方案编码
	Id string `json:"id"`

	// 方案名称
	Name string `json:"name"`

	// 推送时间节点,多个值用逗号','分隔
	NotifyNode string `json:"notifyNode,omitempty"`

	// 方案状态('DRAFT 草稿', 'ENABLED 启用', 'DISABLED 禁用')
	Status string `json:"status"`

	// 方案类型(SA 症状评估, FOLLOWUP 随访)
	Type string `json:"type"`

	// 更新时间
	UpdateTime string `json:"updateTime,omitempty"`
}
