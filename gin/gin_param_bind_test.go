package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestParamBind(t *testing.T) {
	r := gin.Default()
	r.POST("/login", login)
	//http://127.0.0.1:3000/login?username=xujs&password=123
	r.Run(":3000")
}

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func login(c *gin.Context) {
	var user User
	fmt.Println("username:", c.Query("username"))
	fmt.Println("password:", c.Query("password"))
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"password": user.Password,
	})
}
