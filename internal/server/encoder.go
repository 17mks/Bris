package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	nethttp "net/http"
	"strings"
	"time"
)

func GenNewRequestId() string {
	userId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return strings.ToLower(userId)
}

func errorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	var response struct {
		Message   string `json:"message"`
		Status    string `json:"status"`
		DateTime  int64  `json:"datetime"`
		RequestID string `json:"requestID"`
	}
	response.Message = err.Error()
	response.DateTime = time.Now().UnixMilli()
	response.Status = "fail"
	response.RequestID = r.Header.Get("requestID")

	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(200)
	w.Write(body)
}

func responseEncoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
	type Response struct {
		Status    string      `json:"status"`
		RequestID string      `json:"request_id"`
		DateTime  int64       `json:"dateTime"`
		Data      interface{} `json:"data"`
	}

	//Code 与Message 直接写固定值
	res := &Response{
		Status:   "success",
		Data:     data,
		DateTime: time.Now().UnixMilli(),
		//RequestID: r.Header.Get("requestID"),
		RequestID: GenNewRequestId(),
	}

	//序列化
	msRes, err := json.Marshal((res))
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(msRes)
	return nil
}
