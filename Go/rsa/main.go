package main

import (
	"fmt"
	"github.com/wenzhenxi/gorsa"
	"log"
	"reflect"
	"sort"
	"strings"
	v1 "study/rsa/v1"
)

const pub_key = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQClIYYvZvbiivqlp7/uZkZ1jrds
siz8PbpRmKij0Otgn55fJx/H4r/f7VOaZOjSUMlTZH11fMAZl3QNVFrIXk9lAg+4
aqqI6HZXBwmKvn6T832H5Mb2drq9052E7org67jlpDm3+PAliXHBrWlXLE5xuliw
ZREU1tS8yrRvo8RW6QIDAQAB
-----END PUBLIC KEY-----
`
const pri_key = `-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKUhhi9m9uKK+qWn
v+5mRnWOt2yyLPw9ulGYqKPQ62Cfnl8nH8fiv9/tU5pk6NJQyVNkfXV8wBmXdA1U
WsheT2UCD7hqqojodlcHCYq+fpPzfYfkxvZ2ur3TnYTuiuDruOWkObf48CWJccGt
aVcsTnG6WLBlERTW1LzKtG+jxFbpAgMBAAECgYBUTN4E+wqUsqCsywunuBzGTPqF
Chz/FMA2gbmuD6zqIfKm3wFReQe9WJoymstk+wscgJOv0+T566C6YVMLMWJT+5vK
YssVPtASPPMSGwcB4Y7jask8Dp7i9UDMEAE4xVUt+Sb9fqm/R15d872fTVMwzOws
C6VXPCVn72qU1eAQRQJBANBqa2YmvbvyvHL4osO/6BjH6s0h8uTnIPDIP//fFVTJ
RQAEvxTf3EiY2Y+iFupbu4tQFmysUWvfs+z44QhLTVsCQQDK1S1R9+/q7/U6FKmm
EETQOn3iSO+GBUZoDqN+EVlkhxDddtcK9HUqP0R8WwLO02hvUKRBZQorJT06PJeI
g0wLAkEAo+M7Rx2pz3TGaaZI36M10N7MLbjHdualSZI+eWekL4MBxkz1MWYDo/bG
BiOQ34N5C8jTiWtLr4c+xlbDJjGIxwJAZjwnGL643hlO6JSLohlnJfGli84pdMrp
3v8p5xxFi4cMuCPzZiErgTzcfzW8Z5VleV4TPKb6Bh9CCj2KP5O2WwJBAIe7Sdc4
JiitRbL9yeBPr3jN1rNOAWYUVJ12K6GDE/4lQo93v2igvNR/gbB1exsAhCbxKn5x
gBB01975CQeLnvc=
-----END PRIVATE KEY-----
`

func structToMap(s interface{}) map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	fieldMap := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Name
		value := v.FieldByName(key)
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}
		tagfs := strings.Split(tag, ",")
		jsonName := tagfs[0]
		if jsonName == "" {
			jsonName = key
		}
		if !isBlank(value) {
			fieldMap[jsonName] = value.Interface()
		}
	}
	return fieldMap
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func VerifyFlinPlatSign(data interface{}, sign string) error {
	fieldMap := structToMap(data)
	keys := make([]string, 0, len(fieldMap))
	for key := range fieldMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var builder strings.Builder
	for _, key := range keys {
		builder.WriteString(fmt.Sprintf("%v", fieldMap[key]))
	}
	err := gorsa.VerifySignMd5WithRsa(builder.String(), sign, pub_key)
	if err != nil {
		return err
	}
	return nil
}

func SignFlinPlat(data interface{}) (string, error) {
	fieldMap := structToMap(data)
	keys := make([]string, 0, len(fieldMap))
	for key := range fieldMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var builder strings.Builder
	for _, key := range keys {
		builder.WriteString(fmt.Sprintf("%v", fieldMap[key]))
	}
	sign, err := gorsa.SignMd5WithRsa(builder.String(), pri_key)
	if err != nil {
		return "", err
	}
	return sign, nil
}

func main() {
	//req := v1.AddMemberWebReq{
	//	Pwd:    "123456",
	//	Level:  1,
	//	UserId: 77777,
	//}
	req := v1.GivePropWebReq{
		Id:     100,
		Count:  5000000000,
		UserId: 77777,
	}
	sign, err := SignFlinPlat(req)
	if err != nil {
		panic(err)
	}
	log.Printf("sign:%s", sign)
	err = VerifyFlinPlatSign(req, sign)
	if err != nil {
		log.Printf("VerifySignSha1WithRsa err:%s", err)
	}
	//sign, err := gorsa.SignMd5WithRsa(`{"pwd": "123456", "level": 1, "user_id": 77777}`, pri_key)
	//if err != nil {
	//	panic(err)
	//}
	//log.Printf("sign:%s", sign)
	//err = gorsa.VerifySignMd5WithRsa(`{"pwd": "123456", "level": 1, "user_id": 77777}`, sign, pub_key)
	//if err != nil {
	//	log.Printf("VerifySignSha1WithRsa err:%s", err)
	//}

}
