package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

func main() {

	syncMapTest()
	fmt.Println("==========================")
	SendMsg()
	fmt.Println("==========================")
	struct2json()
}

func struct2json() {
	var user User
	user.Name = "张三"
	user.PhoneNumber = "13888888888"
	userJson, err := json.Marshal(user)
	if err == nil {
		fmt.Println(string(userJson))
	} else {
		fmt.Println(err)
		return
	}
}

func syncMapTest() {
	var sm sync.Map
	var user User
	user.Name = "张三"
	user.PhoneNumber = "13888888888"
	sm.Store("user", user)
	load, _ := sm.Load("user")
	u := load.(User)

	fmt.Println(u.Name)

	fmt.Println(sm.Load("user"))

	fmt.Println("==========================")
	m3 := map[string]interface{}{
		"role_key": "admin",
	}
	m2 := map[string]interface{}{
		"name":         "whw2",
		"phone_number": "13333333333",
		"hobbies":      []string{"football", "basketball"},
		"role_info":    m3,
	}

	u2 := User{}
	// 序列化
	arr, err := json.Marshal(m2)
	if err != nil {
		panic(err)
	}
	// 反序列化
	err2 := json.Unmarshal(arr, &u2)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("name>>> ", u2.Name)                // whw2
	fmt.Println("phone_number>>> ", u2.PhoneNumber) // 13333333333
	fmt.Println("hobbies>>> ", u2.Hobbies)          // [football basketball]
	fmt.Println("role_info>>> ", u2.Role)           // map[role_key:admin]
}

type User struct {
	Name        string   `json:"name"`
	PhoneNumber string   `json:"phone_number"`
	Hobbies     []string `json:"hobbies"`
	Role        Role     `json:"role_info"`
}

type Role struct {
	RoleKey string `json:"role_key"`
}

func SendMsg() {
	m2 := map[string]interface{}{
		"previous_version": "1.0.1",
		"current_version":  "1.0.1",
	}
	m3 := map[string]interface{}{
		"previous_version": "1.0.1",
		"current_version":  "1.0.1",
	}
	m1 := map[string]interface{}{
		"msg_type":      "upgrade_virus_db",
		"status":        2,
		"file_id":       1,
		"clamav_engine": m2,
		"avira_engine":  m3,
	}
	u2 := VirusDBVersion{}
	// 序列化
	arr, err := json.Marshal(m1)
	if err != nil {
		panic(err)
	}
	// 反序列化
	err2 := json.Unmarshal(arr, &u2)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(m1)
	fmt.Println("clamav_engine>>> ", u2.ClamAvEngine)
	fmt.Println("clamav_engine version>>> ", u2.ClamAvEngine.CurrentVersion)
	fmt.Println("avira_engine>>> ", u2.AviraEngine)
	fmt.Println("")
}

//UpdateVirusDB 发送给引擎的消息
//type UpdateVirusDB struct {
//	MsgType                    string `json:"msg_type"`                      //消息类型
//	FileId                     int64  `json:"file_id"`                       //文件id
//	ClamavEngineCurrentVersion string `json:"clamav_engine_current_version"` //当前clamav引擎版本
//	AviraEngineCurrentVersion  string `json:"avira_engine_current_version"`  //小红伞引擎当前版本
//	EffectiveDirectory         string `json:"effective_directory"`           //当前生效目录
//}

//VirusDBVersion 更新病毒库版本号
type VirusDBVersion struct {
	MsgType string `json:"msg_type"` //消息类型
	Status  int8   `json:"status"`   //引擎重启状态,1 启动中 2 启动成功 3 启动失败
	FileId  int64  `json:"file_id"`  //文件id
	//ClamavEngineCurrentVersion string `json:"clamav_engine_current_version"` //当前clamav引擎版本
	//AviraEngineCurrentVersion  string `json:"avira_engine_current_version"`  //小红伞引擎当前版本
	ClamAvEngine ClamAvEngine `json:"clamav_engine"` //clamav引擎版本
	AviraEngine  AviraEngine  `json:"avira_engine"`  //小红伞引擎版本

}

type ClamAvEngine struct {
	PreviousVersion string `json:"previous_version"` //上一个版本
	CurrentVersion  string `json:"current_version"`  //当前版本
}
type AviraEngine struct {
	PreviousVersion string `json:"previous_version"` //上一个版本
	CurrentVersion  string `json:"current_version"`  //当前版本
}
