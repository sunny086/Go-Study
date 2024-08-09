package bpm

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestSign(t *testing.T) {

	paramStr := "{\"jcxx-bz\":\"\",\"jcxx-ckfzr\":\"张凯\",\"jcxx-clVINm\":\"LJ1ER2DU5R3450612\",\"jcxx-clnbbh\":\"R3450612\",\"jcxx-clsyssq\":\"安徽省-合肥市-庐阳区\",\"jcxx-clsyxxdz\":\"天威保变（合肥）变压器有限公司庐阳区洪河路58号\",\"jcxx-clxt\":\"整车\",\"jcxx-cph\":\"\",\"jcxx-csys\":\"银河灰\",\"jcxx-cxbm\":\"E1808-01\",\"jcxx-dqck\":\"天津研发整车库\",\"jcxx-dqclfzr\":\"丁致远\",\"jcxx-fj\":[],\"jcxx-ghck\":\"\",\"jcxx-ghrq\":0,\"jcxx-sqbh\":\"BR2024070003\",\"jcxx-sqlx\":\"借用\",\"jcxx-sqr\":\"张凯\",\"jcxx-sqsj\":1719821331983,\"jcxx-yjghrq\":1720454400000,\"jcxx-yt\":\"1、多端远程锁车\\n2、验证OTA升级线上Bug\\n3、App远程控制空调-针对低配车\"}"
	param := make(map[string]interface{})
	err := json.Unmarshal([]byte(paramStr), &param)
	if err != nil {
		t.Fatal(err)
	}
	sign := generateSign(param)
	fmt.Println(sign)

}

func generateSign(param map[string]interface{}) string {
	// 排序
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := ""
	for i, key := range keys {
		if i != 0 {
			str = fmt.Sprintf("%s&", str)
		}
		v, ok := param[key]
		if !ok {
			return ""
		}
		rv := reflect.ValueOf(v)
		rt := rv.Type()
		switch rt.Kind() {
		case reflect.Slice, reflect.Array:
			if rv.Len() == 0 {
				continue
			}
		case reflect.String:
			if v == "" {
				continue
			}
		}
		jsonStr, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		str = fmt.Sprintf("%s%s=%s", str, key, strings.Trim(string(jsonStr), "\""))
	}

	str = fmt.Sprintf("%s&key=%s", str, "D593EA2CA746B9251AC736B002E79484")

	fmt.Println(str)

	// 需要加密的字符串
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
