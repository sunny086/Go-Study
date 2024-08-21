package slice

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

//https://www.jb51.net/article/244560.htm
//还可以使用map进行去重

// ---------------------------------------------------------------------------------------------------------------------
// 测试去重
func TestDataDuplicate(t *testing.T) {
	source := []string{
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1/24",
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.2",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1/24",
		"10.25.10.1/24",
		"10.25.10.1/24",
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1/24",
		"10.25.10.1",
		"10.25.10.2",
		"10.25.10.2",
		"10.25.10.2",
		"10.25.10.3",
		"10.25.10.1/24",
		"10.25.10.1/24",
		"10.25.10.1/24",
	}
	sort.Strings(source)
	index := RemoveDuplicate1(source)
	t.Log(index)
	t.Log(source[:index])
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
	t.Log(RemoveDuplicate1(strList)) //第一次调用
	t.Log(strList)
	t.Log(strList[:RemoveDuplicate1(strList)]) //第二次调用，所以结果有问题
}

// 返回去重后的数据长度 最后一个元素的索引下标
func RemoveDuplicate1(strList []string) int {
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

func TestDuplicate(t *testing.T) {
	test := []string{"10.25.10.1", "10.25.10.11", "10.25.10.12", "10.25.10.12", "10.25.10.1", "10.25.10.11"}
	duplicate := RemoveDuplicate2(test)
	fmt.Println(duplicate)
}

// []string 去重 但是有一个bug 必须初始化一下这个string数组{""}
func RemoveDuplicate2(list []string) []string {
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
