package openapi

type PlanDetailQueryByIdResponse struct {
	Plan Plan `json:"plan"`

	// 方案病种关联
	DiseaseRelates []PlanRelate `json:"diseaseRelates,omitempty"`

	// 方案功能障碍关联
	DysfunctionRelates []PlanRelate `json:"dysfunctionRelates,omitempty"`

	// 适用年龄段(多个年龄段用','分隔)
	ApplyAges []string `json:"applyAges,omitempty"`

	// 方案关联的所有外部资源
	PlanRelates []PlanRelate `json:"planRelates,omitempty"`

	// 方案关联资源详情(Map)
	Resources map[string]interface{} `json:"resources,omitempty"`
}
