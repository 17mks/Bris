package protocol

import (
	"errors"
	"fmt"
	"followup/xcrypto"
	"github.com/goccy/go-json"
	"log"
	"regexp"
	"strings"
	"time"
)

type TokenInfo struct {
	UserId string    `json:"UserId"` // 用户ID
	ExTime time.Time `json:"exTime"` // 过期时间
}

func AddLikeCharToStr(str string) string {
	likeStr := "%"
	for _, val := range str {
		likeStr = fmt.Sprintf("%s%c%c", likeStr, val, '%')
	}
	return likeStr
}

// GenToken 生成TOKEN
func GenToken(userId string) (string, error) {
	if "" == userId {
		return "", errors.New("token generate failed, user id is empty")
	}
	tokenInfo := TokenInfo{
		UserId: userId,
		ExTime: time.Now().Add(24 * time.Hour), // 过期时间：当前时间+24小时
	}
	bytes, err := json.Marshal(tokenInfo)
	if err != nil {
		return "", err
	}

	key, err := xcrypto.RsaEncryptionByPubKey(string(bytes), xcrypto.RsaPublicKey)
	if err != nil {
		return "", err
	}
	return key, nil
}

// TimeToString 将时间转换为sql查询的字符串格式(防止框架出现增加8小时的问题)
func TimeToString(time *time.Time) *string {
	val := ""
	if nil == time || time.IsZero() {
		return &val
	}
	val = time.Format("2006-01-02 15:04:05")
	return &val
}

// ParseTime 将时间字符串转换为time.Time类型
func ParseTime(value string) (*time.Time, error) {
	if "" == value {
		return nil, nil
	}
	// 处理"1971-04-07 00:00:00.0"格式
	if strings.HasSuffix(value, ".0") {
		value = value[0 : len(value)-2]
	}

	layout := "2006-01-02 15:04:05"
	// 时间格式正则判断
	dateTimePatter := `^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\s+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$`
	datePatter := `^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$`

	if match, err := regexp.Match(dateTimePatter, []byte(value)); err != nil {
		return nil, err
	} else if match {
		layout = "2006-01-02 15:04:05"
	}
	if match, err := regexp.Match(datePatter, []byte(value)); err != nil {
		return nil, err
	} else if match {
		layout = "2006-01-02"
	}
	parse, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}

	return &parse, nil
}

// ParseTimeIgnoreErr 将时间字符串转换为time.Time类型,忽略错误
func ParseTimeIgnoreErr(value string) *time.Time {
	t := time.Time{}
	timeStruct, err := ParseTime(value)
	if err != nil {
		log.Println(err)
		return &t
	}

	return timeStruct
}
