package uuid

import (
	"github.com/google/uuid"
	"testing"
)

func TestUuid(t *testing.T) {
	//生成uuid
	uuid := uuid.New()
	t.Log(len(uuid.String()))
	t.Log(uuid.String())
}
