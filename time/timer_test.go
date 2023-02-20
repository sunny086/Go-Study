package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second * 20)
	<-timer.C
	fmt.Println("timer expired")
}
