package openapi

type PlanRelate struct {

	// 创建时间
	CreateTime string `json:"createTime"`

	// 频次间隔(小时)
	FrequencyInterval int32 `json:"frequencyInterval,omitempty"`

	// 频次偏移(小时)
	FrequencyOffset int32 `json:"frequencyOffset,omitempty"`

	// 方案信息编码
	Id string `json:"id"`

	// 关联资源编码
	ResourceId string `json:"resourceId"`

	// 关联资源类型(‘FORM 表单’,'WI 工作项','BZ 病种','GNZA 功能障碍')
	ResourceType string `json:"resourceType"`

	// 排序号码
	SortNum int32 `json:"sortNum,omitempty"`

	TbPlanId string `json:"tbPlanId"`

	// 次数
	Times int32 `json:"times,omitempty"`

	// 标题(根据关联资源类型决定)
	Title string `json:"title"`

	// 更新时间
	UpdateTime string `json:"updateTime"`
}
