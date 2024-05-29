package dw

import (
	"net/url"
	"testing"
)

func TestHeader(t *testing.T) {
	accountInfo := url.QueryEscape("{\"status\":\"success\",\"code\":\"00000\",\"message\":\"成功\",\"tips\":null,\"extra\":null,\"params\":null,\"data\":{\"id\":null,\"userId\":1264010679140,\"avatar\":\"https://static-legacy.dingtalk.com/media/lADPD2sQ519qcqfNAtDNAtA_720_720.jpg\",\"name\":\"许金山\",\"username\":null,\"nickName\":\"许金山\",\"gender\":3,\"genderDesc\":\"\",\"age\":null,\"birthdate\":null,\"mobile\":\"187****3606\",\"email\":\"xujinshan@deepway.ai\",\"profile\":null,\"address\":null,\"company\":\"北京京深深向有限公司\",\"city\":null,\"province\":null,\"area\":\"\",\"provinceCode\":\"\",\"cityCode\":\"\",\"areaCode\":\"\",\"streetAddress\":null,\"externalId\":null,\"postalCode\":null,\"cancellationApplying\":false,\"createTime\":1711992661000,\"stayInUnits\":[{\"poolCode\":\"123456iop\",\"orgCode\":\"orgba84d81af4d541e1821e57c8691cbea5\",\"unitCode\":\"unit96ba9c3f06a44971ac57037a521c2feb\",\"unitName\":\"非生产开发组\",\"unitDesc\":null,\"codePath\":\"orgba84d81af4d541e1821e57c8691cbea5-unit98abeb9d51424840aeb2a5b5c242e915-unitedbdf07e402a4544bc8f8ab0f52e4b17-unitc3b332a1133a4036b9404891134d0374-unit96ba9c3f06a44971ac57037a521c2feb\",\"namePath\":\"北京京深深向有限公司(勿动)-技术研发中心-企业数字化部-数字化开发组-非生产开发组\",\"children\":null}],\"roleList\":[{\"creator\":259654633045,\"creatorName\":null,\"appCode\":\"e9c83e62492f4381bff0c807598499cd\",\"appName\":null,\"id\":211,\"poolCode\":\"123456iop\",\"roleCode\":\"592b33e1d7d34b6984428310fbdc2016\",\"roleCodeName\":\"admin\",\"roleName\":\"系统管理员\",\"roleDesc\":\"\",\"roleType\":2,\"active\":1,\"createTime\":1692068805000,\"memberCount\":12}],\"tenant\":{\"poolCode\":\"123456iop\",\"extra\":\"\",\"poolName\":\"测试环境用户池（不要改）\",\"appCode\":null,\"appName\":null}},\"list\":null,\"map\":null,\"track\":\"b5b8c90bb4644d39b48ad343cfc97e5b.73.17163567389327475\",\"time\":\"566ms\"}")
	t.Log(accountInfo)
}
