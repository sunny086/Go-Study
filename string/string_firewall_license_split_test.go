package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestFirewallLicense(t *testing.T) {
	var decrypt = "SN=WT0210205222071009#prod=firewall;project=firewall;pid=;vid=;sku=;BLACK_LIST=1;BEHAVIOR_LIST=1;modelAuthDay&BLACK_LIST=-1;modelAuthDay&BEHAVIOR_LIST=-1;apply_date=1672934400;expire_date=1704470400;effective_time=31536000;type=1;uuid=50eb539f-b202-4711-a936-3d1e60ea9f5f;#signature=5c4c36c0628e073e8be49d3b5ae1366b6a38035a66afb014cf1da3580f9e17ce"
	split := strings.Split(decrypt, "#")
	contents := strings.Split(split[1], ";")
	applyDate := strings.Split(split[1], ";apply_date=")[1][:13]
	fmt.Println(split)
	fmt.Println(contents)
	fmt.Println(applyDate)
	var str = "1672934400;expire_date=1704470400;effective_time=31536000;type=1;uuid=50eb539f-b202-4711-a936-3d1e60ea9f5f;"
	index := strings.Index(str, ";")
	fmt.Println(index)
	//s := strings.Split(split[1], ";apply_date=")[1]
	//fmt.Println(s)
}
