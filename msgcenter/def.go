package msgcenter

const (
	EventMsgFormWarning = "msg_form_warning" // 表单预警消息
)
const (
	MsgTypeBackgroundMessage = "BM" // 消息类型 - 后台消息(background message)
)

type MsgStructure struct {
	Type      string      `json:"type" comment:"消息类型('')"`
	Payload   interface{} `json:"payload" comment:"数据对象"`
	Time      int64       `json:"time" comment:"时间"`
	RequestId string      `json:"requestId" comment:"请求编码"`
}
