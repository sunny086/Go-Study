package string

import (
	"fmt"
	"strings"
	"testing"
)

// func SplitN(s, sep string, n int) []string { return genSplit(s, sep, 0, n)
// 按照sep分割n份，slices s into substrings separated by sep
func TestStringSplitN(t *testing.T) {
	s := "1a2a3a4a5a6a7a8a9"
	splitN := strings.SplitN(s, "a", 4)
	fmt.Println(splitN)      //[1 2 3 4a5a6a7a8a9]
	fmt.Println(len(splitN)) //4
}

// func SplitAfterN(s, sep string, n int) []string
// 子串和sep结合后的切片，slices s into substrings after each instance of sep
func TestStringSplitAfterN(t *testing.T) {
	s := "1a2a3a4a5a6a7a8a9"
	splitAfterN := strings.SplitAfterN(s, "a", 4)
	fmt.Println(splitAfterN)      //[1a 2a 3a 4a5a6a7a8a9]
	fmt.Println(len(splitAfterN)) //4
}

func TestSplitLogTag(t *testing.T) {
	s := "cm-1#1000#cmTag"
	splitN := strings.SplitN(s, "#", 3)
	fmt.Println(splitN) //[cm-1 1000 cmTag]
	split := strings.Split(s, "#")
	fmt.Println(split) //[cm-1 1000 cmTag]
}
func TestSplitDpiLog(t *testing.T) {
	str := "(0,'10.4.14.102','10.130.130.130',1,'UDP','fins',56,0,'Fins Response timeout, Please Check Operation Rationality!','N/A','{protocol:fins,comc:514}', '2022-08-30 17:05:58',0,'00D003B3A7FC00137297A2D408004500002A828B4000401112CA0A040E660A828282E562258000167DF0C100020000000000007A0C010205', '00:13:72:97:a2:d4','00:d0:03:b3:a7:fc',58722,9600,'24507')"
	str = strings.ReplaceAll(str, " ", "")
	split := strings.Split(str[1:len(str)-1], ",")
	fmt.Println(split)
	sourceIp := strings.Trim(split[1], "'")
	destinationIp := strings.Trim(split[2], "'")
	action := strings.Trim(split[3], "'")
	app := strings.Trim(split[5], "'")
	sourceMac := strings.Trim(split[16], "'")
	destinationMac := strings.Trim(split[17], "'")
	sourcePort := strings.Trim(split[18], "'")
	destinationPort := strings.Trim(split[19], "'")
	fmt.Println(sourceIp)
	fmt.Println(destinationIp)
	fmt.Println(action)
	fmt.Println(app)
	fmt.Println(sourceMac)
	fmt.Println(destinationMac)
	fmt.Println(sourcePort)
	fmt.Println(destinationPort)
}

func TestSplitTableName(t *testing.T) {
	tableName := "log_audit_flow_11"
	//拿到最后一个_的位置
	lastIndex := strings.LastIndex(tableName, "_")
	fmt.Println(lastIndex)
	//拿到最后一个_后面的字符串
	suffix := tableName[lastIndex+1:]
	fmt.Println(suffix)
	//拿到最后一个_前面的字符串
	prefix := tableName[:lastIndex]
	fmt.Println(prefix)

}
