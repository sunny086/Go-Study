package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

var client *gorm.DB

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:tiger@tcp(127.0.0.1:3306)/firewall?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                             // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                            // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                            // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                           // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm.Open err:", err)
	}
	client = db
}

func TestGormCasbin(t *testing.T) {
	var err error
	var b bool
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// If it doesn't exist, the adapter will create it automatically.
	gormAdapter, _ := gormadapter.NewAdapterByDBUseTableName(client, "aaa", "casbin")
	enforcer, err := casbin.NewEnforcer("./rabc_model.conf", gormAdapter)
	if err != nil {
		fmt.Println("casbin.NewEnforcer err:", err)
	}

	// Load the policy from DB.
	err = enforcer.LoadPolicy()
	if err != nil {
		t.Log("e.LoadPolicy err:" + err.Error())
	}

	// Check the permission.
	b, err = enforcer.Enforce("Gorilla", "/api/add", "POST")
	if err != nil {
		t.Log("e.Enforce err:" + err.Error())
	}
	if !b {
		t.Log("enforce failed")
	} else {
		t.Log("enforce success")
	}

	// Modify the policy.
	enforcer.AddPolicy("Gorilla", "/api/query", "POST")
	enforcer.RemovePolicy("Gorilla", "/api/query", "POST")

	// Save the policy back to DB.
	err = enforcer.SavePolicy()
	if err != nil {
		t.Log("e.SavePolicy err:" + err.Error())
	}
}

func TestCustom(t *testing.T) {
	r := gin.Default()
	r.POST("/api/add", func(c *gin.Context) {
		path := c.Request.URL.Path
		// 获取请求方法
		method := c.Request.Method
		var count int64
		var msg string
		err := client.Table("aaa_casbin").Debug().Where("v0 = ? and v1 = ? and v2 = ?", "Gorilla", path, method).Count(&count).Error
		if err != nil {
			msg = "查询失败"
			t.Log("client.Table err:", err.Error())
			c.JSON(200, gin.H{
				"msg": msg,
			})
			return
		}
		if count > 0 {
			msg = "api auth success"
		} else {
			msg = "api auth failed"
		}
		c.JSON(200, gin.H{
			"msg": msg,
		})
	})
	r.Run(":9999")
}
