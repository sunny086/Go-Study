package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:tiger@tcp(127.0.0.1:3306)/usb?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	//var menus []SysMenu
	//err = db.Model(&SysMenu{}).Scopes(Scope1, Scope2).Debug().Find(&menus).Error
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(menus)

	err = db.Model(&SysUser{}).Debug().Where("user_id = 2").Update("lock_time", time.Now().Format("2006-01-02 15:04:05")).Error
	if err != nil {
		fmt.Println(err)
	}
}

func Scope1(db *gorm.DB) *gorm.DB {
	return db.Where("module = 1")
}

func Scope2(db *gorm.DB) *gorm.DB {
	return db.Where("menu_id = 543")
}

type SysMenu struct {
	MenuId     int       `json:"menuId" gorm:"primaryKey;autoIncrement"`
	MenuName   string    `json:"menuName" gorm:"size:128;"`
	Title      string    `json:"title" gorm:"size:128;"`
	Icon       string    `json:"icon" gorm:"size:128;"`
	Path       string    `json:"path" gorm:"size:128;"`
	Paths      string    `json:"paths" gorm:"size:128;"`
	MenuType   string    `json:"menuType" gorm:"size:1;"`
	Action     string    `json:"action" gorm:"size:16;"`
	Permission string    `json:"permission" gorm:"size:255;"`
	ParentId   int       `json:"parentId" gorm:"size:11;"`
	NoCache    bool      `json:"noCache" gorm:"size:8;"`
	Breadcrumb string    `json:"breadcrumb" gorm:"size:255;"`
	Component  string    `json:"component" gorm:"size:255;"`
	Sort       int       `json:"sort" gorm:"size:4;"`
	Visible    string    `json:"visible" gorm:"size:1;"`
	IsFrame    string    `json:"isFrame" gorm:"size:1;DEFAULT:0;"`
	Apis       []int     `json:"apis" gorm:"-"`
	DataScope  string    `json:"dataScope" gorm:"-"`
	Params     string    `json:"params" gorm:"-"`
	RoleId     int       `gorm:"-"`
	Children   []SysMenu `json:"children,omitempty" gorm:"-"`
	IsSelect   bool      `json:"is_select" gorm:"-"`
}

func (SysMenu) TableName() string {
	return "sys_menu"
}

type SysUser struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName,omitempty" gorm:"size:128;comment:昵称"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId   int    `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt     string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar   string `json:"avatar,omitempty" gorm:"size:255;comment:头像"`
	Sex      string `json:"sex,omitempty" gorm:"size:255;comment:性别"`
	Email    string `json:"email,omitempty" gorm:"size:128;comment:邮箱"`
	DeptId   int    `json:"deptId,omitempty" gorm:"size:20;comment:部门"`
	PostId   int    `json:"postId,omitempty" gorm:"size:20;comment:岗位"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Status   string `json:"status" gorm:"size:4;comment:状态"`
	DeptIds  []int  `json:"deptIds,omitempty" gorm:"-"`
	PostIds  []int  `json:"postIds,omitempty" gorm:"-"`
	RoleIds  []int  `json:"roleIds,omitempty" gorm:"-"`
	Ip       string `json:"ip" gorm:"type:varchar(30);comment:可信主机ip"`
	StartIp  int64  `json:"startIp" gorm:"type:int(64);comment:起始Ip"`
	EndIp    int64  `json:"endIp" gorm:"type:int(64);comment:起始Ip"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
