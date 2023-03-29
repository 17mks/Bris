package msgcenter

import (
	"encoding/json"
	"errors"
	"followup/model"
	"followup/xcrypto"
	"log"
	"net/url"
	"sync"

	socketIO "github.com/googollee/go-socket.io"
)

type MsgService interface {
	initServer()
	Server() *socketIO.Server
	Emit(eventName string, userId string, v ...interface{})
	Emits(eventName string, userIds []string, v ...interface{})
}

type EProMsgService struct {
	server      *socketIO.Server // socket io server
	connections sync.Map         // cli connections
	idMaps      sync.Map         // conn id与userId对应关系
}

var (
	msgService     MsgService
	msgServiceOnce sync.Once
)

func GetEProMsgService() MsgService {
	msgServiceOnce.Do(func() {
		msgService = &EProMsgService{}
		msgService.initServer()
	})
	return msgService
}

func (service *EProMsgService) Server() *socketIO.Server {
	return service.server
}

func (service *EProMsgService) initServer() {
	server := socketIO.NewServer(nil)

	server.OnConnect("/", func(conn socketIO.Conn) error {
		// 解析查询参数
		ioQueryParams := model.SocketIOConnInfo{}
		if err := service.ParseQueryParams(conn.URL(), &ioQueryParams); err != nil {
			return err
		}
		log.Println(ioQueryParams)
		if "" != ioQueryParams.Token {
			ioQueryParams.Conn = conn
			service.PutNewConn(ioQueryParams)
		}

		return nil
	})

	server.OnEvent("/", "notice", func(conn socketIO.Conn, msg string) {
		log.Println("OnEvent:", msg)
		conn.Emit("connection", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(conn socketIO.Conn, msg string) string {
		conn.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/websocket", "test", func(conn socketIO.Conn, msg string) string {
		conn.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(conn socketIO.Conn) string {
		last := conn.Context().(string)
		conn.Emit("bye", last)
		_ = conn.Close()
		return last
	})

	server.OnError("/", func(conn socketIO.Conn, e error) {
		log.Println("OnError,Error = ", e.Error())
	})

	server.OnDisconnect("/", func(conn socketIO.Conn, msg string) {
		log.Println("OnDisconnect,ID = ", conn.ID(), ",Msg = ", msg)
		service.RemoveConn(conn)
	})
	service.server = server
}

// Emit 向某个用户发送消息
func (service *EProMsgService) Emit(eventName string, userId string, v ...interface{}) {
	value, ok := service.connections.Load(userId)
	if !ok {
		log.Printf("user '%s' not connected yet,can't emit msg", userId)
		return
	}
	//log.Println("type = ", reflect.TypeOf(value))
	if connInfo, ok := value.(model.SocketIOConnInfo); ok {
		connInfo.Conn.Emit(eventName, v)
	}
}

func (service *EProMsgService) Emits(eventName string, userIds []string, v ...interface{}) {
	for _, userId := range userIds {
		service.Emit(eventName, userId, v)
	}
}

func (service *EProMsgService) ParseQueryParams(connUrl url.URL, queryParams *model.SocketIOConnInfo) error {
	values, err := url.ParseQuery(connUrl.RawQuery)
	if err != nil {
		return err
	}
	// 解析EIO,需要是版本3
	if eio, ok := values["EIO"]; ok && len(eio) > 0 {
		if eio[0] != "3" {
			return errors.New("the EIO version does not match the server")
		}
		queryParams.EIO = eio[0]
	}
	// 解析CliType
	if cliType, ok := values["type"]; ok && len(cliType) > 0 {
		queryParams.CliType = cliType[0]
	}
	// 解析EIO,需要是版本3
	if transport, ok := values["transport"]; ok && len(transport) > 0 {
		queryParams.Transport = transport[0]
	}
	// 解析TOKEN
	if token, ok := values["token"]; ok && len(token) > 0 {
		// 解密TOKEN获取JSON对象
		tokenJsonStr, err := xcrypto.RsaDecryptionByPriKey(token[0], xcrypto.RsaPrivateKey)
		if err != nil {
			return err
		}
		tokenInfo := model.TokenInfo{}
		if err := json.Unmarshal([]byte(tokenJsonStr), &tokenInfo); err != nil {
			return err
		}
		// 用户ID和Token有效期校验
		if "" == tokenInfo.UserId {
			return errors.New("bad token generator, user id is empty")
		}
		queryParams.Token = token[0]
		queryParams.TokenInfo = tokenInfo
	}

	return nil
}

func (service *EProMsgService) PutNewConn(connInfo model.SocketIOConnInfo) {
	userId := connInfo.TokenInfo.UserId
	connId := connInfo.Conn.ID()
	log.Println("PutNewConn,userId = ", userId, ",connId = ", connId, ",RemoteAddr = ", connInfo.Conn.RemoteAddr())
	service.connections.Store(userId, connInfo)
	service.idMaps.Store(connId, userId)
}

func (service *EProMsgService) RemoveConn(conn socketIO.Conn) {
	userId, ok := service.idMaps.Load(conn.ID())
	if ok {
		service.idMaps.Delete(conn.ID())
		service.connections.Delete(userId)
	}

}
