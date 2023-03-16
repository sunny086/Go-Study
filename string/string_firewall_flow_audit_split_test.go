package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestFlowAuditSplit(t *testing.T) {
	message := "25;2023-03-13 14:26:12;59986;22;10.25.16.6;a4:1a:3a:91:a4:1a;273;0----\n"
	//message := "0;2022-12-28 10:15:20;14336;12288;172.21.1.254;c8:5b:76:ef:ad:ee;67;0-------"
	messageContent := strings.Split(message, "----")
	split := strings.Split(messageContent[0], ";")
	fmt.Println(split[0])
}
