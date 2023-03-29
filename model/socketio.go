package model

import socketIO "github.com/googollee/go-socket.io"

// SocketIOConnInfo SocketIO连接信息
type SocketIOConnInfo struct {
	EIO       string // 目前服务端使用的是版本3，Cli需要与服务端同步；
	Transport string
	CliType   string
	Token     string    // RSA加密TOKEN
	TokenInfo TokenInfo // Token解密后数据
	Conn      socketIO.Conn
}

// MsgPushInfo 消息推送信息
type MsgPushInfo struct {
	Type      string   `json:"type" comment:"消息类型"` //
	PatientId string   `json:"patientId" comment:"患者编码"`
	WorkItems []string `json:"workItems" comment:"工作项编码"`
}
