package hr

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTempDir(t *testing.T) {
	sprintf := fmt.Sprintf("%s%d_%d", "./tmp/", time.Now().Unix(), rand.Intn(999999))
	fmt.Println(sprintf)
}
