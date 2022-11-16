package refelct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type SecurityDeepPacketInspectionSearch struct {
	Name          string   `json:"name" form:"name"`                   // 名称
	SourceIp      string   `json:"sourceIp" form:"sourceIp"`           // 源IP
	DestinationIp string   `json:"destinationIp" form:"destinationIp"` // 目的IP
	AppName       string   `json:"appName" form:"appName"`             // 应用-协议名称
	Content       string   `json:"content" form:"content"`             // 内容
	Status        int      `json:"status" form:"status"`               // 状态：0禁用，1启用
	FieldArray    []string `json:"fieldArray" form:"fieldArray"`       // 字段数组
	SearchContent string   `json:"searchContent" form:"searchContent"` // 搜索内容
	Action        int      `json:"action" form:"action"`               // 动作：0允许，1告警，2阻断
}

// TestReflect_SetStructValueByFieldArray 根据结构体的属性字段判断结构体是否存在 并获取内部的tag标签
func TestReflect_SetStructValueByFieldArray(t *testing.T) {
	var req SecurityDeepPacketInspectionSearch
	req.FieldArray = []string{"SourceIp", "Name", "DestinationIp", "AppName"}
	req.SearchContent = "123"
	//遍历字段数组 通过反射 设置属性值
	for _, fieldKey := range req.FieldArray {
		//反射获取json tag
		field, exist := reflect.TypeOf(req).FieldByName(fieldKey)
		if exist {
			//字段存在 设置属性值
			reflect.ValueOf(&req).Elem().FieldByName(fieldKey).SetString(req.SearchContent)
			tag := field.Tag
			fmt.Println(tag)
			// 获取json属性
			jsonTagValue := field.Tag.Get("json")
			fmt.Println(jsonTagValue)
		}
	}
	//解析成字符串
	marshal, err := json.Marshal(req)
	if err == nil {
		fmt.Println(string(marshal))
	}
}

func TestReflect_GetStructAllField(t *testing.T) {
	var req SecurityDeepPacketInspectionSearch
	//获取结构体的所有字段
	typeOf := reflect.TypeOf(req)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fmt.Println("field name:" + field.Name)
	}
}
