package slice

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// 测试去重
func TestDataDuplicate(t *testing.T) {
	strList := []string{"1", "1", "2", "2", "3", "3", "4", "4", "5", "5"}
	t.Log(DataDuplicate(strList))           //5
	t.Log(strList)                          //[1 2 3 4 5 3 4 4 5 5]
	t.Log(strList[:DataDuplicate(strList)]) //[1 2 3 4 5 3 4 5]
}

// 测试去重 因为使用的是同一个地址引用 调用多次会导致数据变化
func TestDataDuplicate2(t *testing.T) {
	strList := []string{
		"10.25.10.1",
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.3",
		"10.25.10.1/24",
		"10.25.10.1/24",
	}
	t.Log(DataDuplicate(strList)) //第一次调用
	t.Log(strList)
	t.Log(strList[:DataDuplicate(strList)]) //第二次调用，所以结果有问题
}

// 返回去重后的数据长度 最后一个元素的索引下标
func DataDuplicate(strList []string) int {
	//如果是空切片，那就返回0
	if len(strList) == 0 {
		return 0
	}
	//用两个标记来比较相邻位置的值
	//当一样的话，那就不管继续
	//当不一样的时候，就把right指向的值赋值给left下一位
	left, right := 0, 1
	for ; right < len(strList); right++ {
		if strList[left] == strList[right] {
			continue
		}
		left++
		strList[left] = strList[right]
	}
	return left + 1
}

// ---------------------------------------------------------------------------------------------------------------------

// []string 去重 但是有一个bug 必须初始化一下这个string数组{""}
func RemoveDuplicate(list []string) []string {
	// 这个排序很关键
	sort.Strings(list)
	i := 0
	var newlist = []string{""}
	for j := 0; j < len(list); j++ {
		if strings.Compare(newlist[i], list[j]) == -1 {
			newlist = append(newlist, list[j])
			i++
		}
	}
	return newlist
}

func TestDuplicate(t *testing.T) {
	test := []string{"10.25.10.1", "10.25.10.11", "10.25.10.12", "10.25.10.12", "10.25.10.1", "10.25.10.11"}
	duplicate := RemoveDuplicate(test)
	fmt.Println(duplicate)
}
