package slice

import "testing"

// 测试去重
func TestDataDuplicate(t *testing.T) {
	strList := []string{"1", "1", "2", "2", "3", "3", "4", "4", "5", "5"}
	t.Log(DataDuplicate(strList))           //5
	t.Log(strList)                          //[1 2 3 4 5 3 4 4 5 5]
	t.Log(strList[:DataDuplicate(strList)]) //[1 2 3 4 5 3 4 5]
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
