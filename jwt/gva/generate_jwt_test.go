package gva

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenerateJwt(t *testing.T) {
	j := &JWT{SigningKey: NewJWT().SigningKey} // 唯一签名
	claims := j.CreateClaims(BaseClaims{
		UUID:        uuid.NewV4(),
		ID:          1,
		NickName:    "smp",
		Username:    "smp",
		AuthorityId: 888,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
