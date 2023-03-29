package chdb

import "time"

// TbWorkItem 工作项表
type TbWorkItem struct {
	ID                string     `gorm:"primaryKey;column:id;type:varchar(45);not null;comment:'工作项编码'"`                                                                                                                                            // 工作项编码
	Title             string     `gorm:"column:title;type:varchar(255);default:null;comment:'标题'"`                                                                                                                                                  // 标题
	Type              string     `gorm:"column:type;type:enum('TASK','SA','TREAT','FOLLOW_UP','WORK_FLOW');not null;comment:'工作项类型(FOLLOW_UP:随访,SA:症状评估,TREAT:治疗任务,WORK_FLOW:工作流)'"`                                                                // 工作项类型(FOLLOW_UP:随访,SA:症状评估,TREAT:治疗任务,WORK_FLOW:工作流)
	Status            string     `gorm:"column:status;type:enum('MODEL','DRAFT','NEW','ACTIVE','RESOLVED','CLOSED','REMOVED');not null;comment:'状态('MODEL 模板', 'DRAFT 草稿', 'NEW 新建-待处理', 'ACTIVE 正在进行', 'RESOLVED', 'CLOSED 已关闭', 'REMOVED 已移除')'"` // 状态('MODEL 模板', 'DRAFT 草稿', 'NEW 新建-待处理', 'ACTIVE 正在进行', 'RESOLVED', 'CLOSED 已关闭', 'REMOVED 已移除')
	PrincipalType     string     `gorm:"column:principal_type;type:enum('USER','MEMBER','NONE');default:null;default:NONE;comment:'负责人类型('USER', 'MEMBER', 'NONE')'"`                                                                               // 负责人类型('USER', 'MEMBER', 'NONE')
	PrincipalID       string     `gorm:"column:principal_id;type:varchar(45);default:null;comment:'负责人成员编码'"`                                                                                                                                       // 负责人成员编码
	PrincipalName     string     `gorm:"column:principal_name;type:varchar(128);default:null;comment:'负责人姓名'"`                                                                                                                                      // 负责人姓名
	Participant       string     `gorm:"column:participant;type:text;default:null;comment:'参与者成员编码(多个值用','分隔)'"`                                                                                                                                    // 参与者成员编码(多个值用','分隔)
	Cc                string     `gorm:"column:cc;type:text;default:null;comment:'抄送成员编码(多个值用','分隔)'"`                                                                                                                                              // 抄送成员编码(多个值用','分隔)
	Tag               string     `gorm:"column:tag;type:text;default:null;comment:'标签(多个值用','分隔)'"`                                                                                                                                                 // 标签(多个值用','分隔)
	Pid               string     `gorm:"column:pid;type:varchar(45);default:null;comment:'父工作项编码'"`                                                                                                                                                 // 父工作项编码
	AssignedType      string     `gorm:"column:assigned_type;type:enum('USER','MEMBER','PATIENT','NONE');not null;default:NONE;comment:'指派类型(用户,成员,患者,未指派)'"`                                                                                       // 指派类型(用户,成员,患者,未指派)
	AssignedTo        string     `gorm:"column:assigned_to;type:text;default:null;comment:'指派给谁(多个值用','分隔)'"`                                                                                                                                       // 指派给谁(多个值用','分隔)
	AssignedToName    string     `gorm:"column:assigned_to_name;type:varchar(128);default:null;comment:'被指派人姓名'"`                                                                                                                                   // 被指派人姓名
	PlanStartTime     *time.Time `gorm:"column:plan_start_time;type:datetime;default:null;comment:'计划开始时间'"`                                                                                                                                        // 计划开始时间
	PlanEndTime       *time.Time `gorm:"column:plan_end_time;type:datetime;default:null;comment:'计划结束时间'"`                                                                                                                                          // 计划结束时间
	ActualStartTime   *time.Time `gorm:"column:actual_start_time;type:datetime;default:null;comment:'实际开始时间'"`                                                                                                                                      // 实际开始时间
	ActualEndTime     *time.Time `gorm:"column:actual_end_time;type:datetime;default:null;comment:'实际结束时间'"`                                                                                                                                        // 实际结束时间
	CreatBy           string     `gorm:"column:creat_by;type:varchar(45);default:null;comment:'创建人成员编码'"`                                                                                                                                           // 创建人成员编码
	CreatByName       string     `gorm:"column:creat_by_name;type:varchar(128);default:null;comment:'创建人姓名'"`                                                                                                                                       // 创建人姓名
	UpdateBy          string     `gorm:"column:update_by;type:varchar(45);default:null;comment:'最近更新人成员编码'"`                                                                                                                                        // 最近更新人成员编码
	BelongType        string     `gorm:"column:belong_type;type:enum('ORG','PROJECT','TEAM','GROUP','NONE');not null;default:NONE;comment:'归属类型'"`                                                                                                  // 归属类型
	BelongTo          string     `gorm:"column:belong_to;type:varchar(45);default:null;comment:'规则资源编码(e.g. 如果归属类型是项目则填项目编码)'"`                                                                                                                     // 规则资源编码(e.g. 如果归属类型是项目则填项目编码)
	SortNum           uint32     `gorm:"column:sort_num;type:int(10) unsigned;default:null;comment:'排序号码'"`                                                                                                                                         // 排序号码
	Event             string     `gorm:"column:event;type:enum('RY','SS','CY','JZ','CS','RC','NONE');not null;default:NONE;comment:'触发事件('RY 入院','SS 手术', 'CY 出院', 'JZ 就诊', 'CS 出生','RC 入组','NONE')'"`                                              // 触发事件('RY 入院','SS 手术', 'CY 出院', 'JZ 就诊', 'CS 出生','RC 入组','NONE')
	FrequencyInterval int32      `gorm:"column:frequency_interval;type:int(11);default:null;comment:'频次间隔'"`                                                                                                                                        // 频次间隔
	FrequencyUnit     string     `gorm:"column:frequency_unit;type:enum('DAY','WEEK','MONTH','YEAR','NONE');default:null;comment:'频次单位(年,月,周,日,小时)'"`                                                                                               // 频次单位(年,月,周,日,小时)
	AppID             string     `gorm:"column:app_id;type:varchar(45);default:null;comment:'应用编码'"`                                                                                                                                                // 应用编码
	NotifyLeftOffset  int32      `gorm:"column:notify_left_offset;type:int(11);default:null;comment:'通知左偏移'"`                                                                                                                                       // 通知左偏移
	NotifyRightOffset int32      `gorm:"column:notify_right_offset;type:int(11);default:null;comment:'通知右偏移'"`                                                                                                                                      // 通知右偏移
	NotifyOffsetUnit  string     `gorm:"column:notify_offset_unit;type:enum('DAY','WEEK','MONTH','YEAR','NONE');default:null;default:NONE;comment:'通知偏移时间单位'"`                                                                                      // 通知偏移时间单位
	NotifyLeftDate    *time.Time `gorm:"column:notify_left_date;type:datetime;default:null;comment:'通知开始时间'"`                                                                                                                                       // 通知开始时间
	NotifyRightDate   *time.Time `gorm:"column:notify_right_date;type:datetime;default:null;comment:'通知截止时间'"`                                                                                                                                      // 通知截止时间
	NotifyNode        string     `gorm:"column:notify_node;type:varchar(255);default:null;comment:'推送时间节点,多个值用逗号','分隔'"`                                                                                                                            // 推送时间节点,多个值用逗号','分隔
	ExecArea          string     `gorm:"column:exec_area;type:enum('INSIDE','OUTSIDE','NONE');default:null;default:NONE;comment:'执行区域(院内,院外)'"`                                                                                                     // 执行区域(院内,院外)
	Description       string     `gorm:"column:description;type:text;default:null;comment:'描述(可存富文本内容或Markdown内容)'"`                                                                                                                                // 描述(可存富文本内容或Markdown内容)
	CreateTime        *time.Time `gorm:"column:create_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`                                                                                                                    // 创建时间
	UpdateTime        *time.Time `gorm:"column:update_time;type:datetime;default:null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`                                                                                                                    // 更新时间
	DeletedAt         *time.Time `gorm:"column:deleted_at;type:datetime;default:null;comment:'删除时间'"`                                                                                                                                               // 删除时间
}

// TableName get sql table name.获取数据库表名
func (m *TbWorkItem) TableName() string {
	return "tb_work_item"
}
