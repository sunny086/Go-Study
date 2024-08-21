package udp

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestUdpAptNta(t *testing.T) {
	serverAddr := "10.25.16.41:5514"
	message := "<14>\"2018-12-23 17:17:42|!secevent|!!192.3.221.70|!{\"dealStatus\": 0, \"firstTime\": 1559379361, \"msg_type\": \"有害程序事件\", \"scan_time\": \"\", \"ip\": \"10.0.0.101\", \"tampered_url\": \"\", \"hostRisk\": \"10.0.0.101(未归类终端)\", \"password_autofill\": \"\", \"tag\": [\"黑URL\"], \"detectEngine\": \"安全日志分析引擎\", \"threat_level\": 2, \"dst_mac_addr\": \"\", \"upstream_traffic\": \"\", \"scan_success\": \"\", \"src_port\": \"\", \"branchId\": \"0\", \"src_mac_addr\": \"\", \"directory_browsing\": \"\", \"domain_name\": \"\", \"rsp_email\": \"\", \"src_ip\": \"10.0.0.101\", \"groupName\": \"未归类终端\", \"event_evidence\": \"主机访问URL <a class='btnDetail' btncon='<div class=\\\"viewUrl\\\" url=\\\"https://wiki.sec.sangfor.com.cn/api/v1/wiki-info?type=url&info=eyJkZXZpY2VTZXJpYWxOdW1iZXIiOiJFQjZDOURFOUNGOTJDMTk2MjQxRDJGRjZGQUExOEJGNSIsImRldmljZVZlcnNpb24iOiIzLjAuMjQiLCJkZXZpY2VJZCI6IjYxQTJFRTQ3In0=&content=d3d3LmZ5bDguY29tL2d1b25laXhpbndlbi8xMzg4NDUuaHRtbA==&source=NTA\\\">威胁维基</div><div class=\\\"viewLog\\\" param={\\\"domain_or_url\\\":\\\"www.fyl8.com/guoneixinwen/138845.html\\\"}>日志详情</div>'>www.fyl8.com/guoneixinwen/138845.html</a>（10次），其中控制者IP（TOP3）：<a class='btnDetail' btncon='<div class=\\\"viewUrl\\\" url=\\\"https://wiki.sec.sangfor.com.cn/api/v1/wiki-info?type=ip&info=eyJkZXZpY2VTZXJpYWxOdW1iZXIiOiJFQjZDOURFOUNGOTJDMTk2MjQxRDJGRjZGQUExOEJGNSIsImRldmljZVZlcnNpb24iOiIzLjAuMjQiLCJkZXZpY2VJZCI6IjYxQTJFRTQ3In0=&content=MS4xLjEuMQ==&source=NTA\\\">威胁维基</div><div class=\\\"viewLog\\\" param={\\\"dst_ip\\\":\\\"1.1.1.1\\\"}>日志详情</div>'>1.1.1.1</a>（-，10次）；\", \"attachment\": \"\", \"eventType\": 0, \"recordDate\": 20190601, \"type\": 3, \"dst_port\": \"\", \"hostName\": \"\", \"phishing_link\": \"\", \"ruleId\": \"103002\", \"attack_time\": \"\", \"attack_count\": \"\", \"suspect_level\": 2, \"branch_name\": \"-\", \"msg_sub_type\": \"僵尸网络事件\", \"downstream_traffic\": \"\", \"data_collection_size\": \"\", \"dst_ip\": \"0.0.0.0\", \"log_time\": 1559379361, \"principle\": \"cncert是指国家计算机网络应急技术处理协调中心（简称“国家互联网应急中心”），是我国网络安全应急体系的核心协调机构。VirusTotal.com是一个免费的病毒，蠕虫，木马和各种恶意软件分析服务，可以针对可疑文件和网址进行快速检测，最初由Hispasec维护。这些机构向公众提供已知的C&C服务器IP/域名/URL，如果组织单位的内网主机存在访问这些C&C服务器的行为，则说明主机正在与黑客通信。\", \"dst_group\": \"\", \"groupId\": \"2147483647\", \"vulnerability_name\": \"\", \"stage\": 2, \"msg_count\": \"\", \"email_from\": \"\", \"server_sensative_directory\": \"\", \"session_token\": \"\", \"url\": \"\", \"scan_count\": \"\", \"event_content\": \"主机访问了cncert/virustotal等机构标记的C&C通信URL\", \"solution\": \"1、确认该主机是否为DNS服务器或域控服务器（DNS代理），如果是请将该主机IP添加到失陷主机白名单即可；\\n2、推荐使用深信服EDR工具进行病毒查杀：http://edr.sangfor.com.cn/tool/SfabAntiBot.zip\\n3、如以上推荐工具查杀不出来，可使用第三方杀毒工具（如火绒等）进行查杀；\", \"src_group\": \"\", \"dst_host\": \"\", \"eventKey\": \"coengine|103002\", \"email_to\": \"\", \"vulnerability_description\": \"\", \"data_collection_time\": \"\", \"src_host\": \"\"}"
	for i := 0; i < 10; i++ {
		SendMessage(message, serverAddr)
	}
}

func TestUdpAptNtaOnce(t *testing.T) {
	serverAddr := "10.25.10.88:5514"
	message := "<14>\"2018-12-23 17:17:42|!secevent|!!192.3.221.70|!{\"dealStatus\": 0, \"firstTime\": 1559379361, \"msg_type\": \"有害程序事件\", \"scan_time\": \"\", \"ip\": \"10.0.0.101\", \"tampered_url\": \"\", \"hostRisk\": \"10.0.0.101(未归类终端)\", \"password_autofill\": \"\", \"tag\": [\"黑URL\"], \"detectEngine\": \"安全日志分析引擎\", \"threat_level\": 2, \"dst_mac_addr\": \"\", \"upstream_traffic\": \"\", \"scan_success\": \"\", \"src_port\": \"\", \"branchId\": \"0\", \"src_mac_addr\": \"\", \"directory_browsing\": \"\", \"domain_name\": \"\", \"rsp_email\": \"\", \"src_ip\": \"10.0.0.101\", \"groupName\": \"未归类终端\", \"event_evidence\": \"主机访问URL <a class='btnDetail' btncon='<div class=\\\"viewUrl\\\" url=\\\"https://wiki.sec.sangfor.com.cn/api/v1/wiki-info?type=url&info=eyJkZXZpY2VTZXJpYWxOdW1iZXIiOiJFQjZDOURFOUNGOTJDMTk2MjQxRDJGRjZGQUExOEJGNSIsImRldmljZVZlcnNpb24iOiIzLjAuMjQiLCJkZXZpY2VJZCI6IjYxQTJFRTQ3In0=&content=d3d3LmZ5bDguY29tL2d1b25laXhpbndlbi8xMzg4NDUuaHRtbA==&source=NTA\\\">威胁维基</div><div class=\\\"viewLog\\\" param={\\\"domain_or_url\\\":\\\"www.fyl8.com/guoneixinwen/138845.html\\\"}>日志详情</div>'>www.fyl8.com/guoneixinwen/138845.html</a>（10次），其中控制者IP（TOP3）：<a class='btnDetail' btncon='<div class=\\\"viewUrl\\\" url=\\\"https://wiki.sec.sangfor.com.cn/api/v1/wiki-info?type=ip&info=eyJkZXZpY2VTZXJpYWxOdW1iZXIiOiJFQjZDOURFOUNGOTJDMTk2MjQxRDJGRjZGQUExOEJGNSIsImRldmljZVZlcnNpb24iOiIzLjAuMjQiLCJkZXZpY2VJZCI6IjYxQTJFRTQ3In0=&content=MS4xLjEuMQ==&source=NTA\\\">威胁维基</div><div class=\\\"viewLog\\\" param={\\\"dst_ip\\\":\\\"1.1.1.1\\\"}>日志详情</div>'>1.1.1.1</a>（-，10次）；\", \"attachment\": \"\", \"eventType\": 0, \"recordDate\": 20190601, \"type\": 3, \"dst_port\": \"\", \"hostName\": \"\", \"phishing_link\": \"\", \"ruleId\": \"103002\", \"attack_time\": \"\", \"attack_count\": \"\", \"suspect_level\": 2, \"branch_name\": \"-\", \"msg_sub_type\": \"僵尸网络事件\", \"downstream_traffic\": \"\", \"data_collection_size\": \"\", \"dst_ip\": \"0.0.0.0\", \"log_time\": 1559379361, \"principle\": \"cncert是指国家计算机网络应急技术处理协调中心（简称“国家互联网应急中心”），是我国网络安全应急体系的核心协调机构。VirusTotal.com是一个免费的病毒，蠕虫，木马和各种恶意软件分析服务，可以针对可疑文件和网址进行快速检测，最初由Hispasec维护。这些机构向公众提供已知的C&C服务器IP/域名/URL，如果组织单位的内网主机存在访问这些C&C服务器的行为，则说明主机正在与黑客通信。\", \"dst_group\": \"\", \"groupId\": \"2147483647\", \"vulnerability_name\": \"\", \"stage\": 2, \"msg_count\": \"\", \"email_from\": \"\", \"server_sensative_directory\": \"\", \"session_token\": \"\", \"url\": \"\", \"scan_count\": \"\", \"event_content\": \"主机访问了cncert/virustotal等机构标记的C&C通信URL\", \"solution\": \"1、确认该主机是否为DNS服务器或域控服务器（DNS代理），如果是请将该主机IP添加到失陷主机白名单即可；\\n2、推荐使用深信服EDR工具进行病毒查杀：http://edr.sangfor.com.cn/tool/SfabAntiBot.zip\\n3、如以上推荐工具查杀不出来，可使用第三方杀毒工具（如火绒等）进行查杀；\", \"src_group\": \"\", \"dst_host\": \"\", \"eventKey\": \"coengine|103002\", \"email_to\": \"\", \"vulnerability_description\": \"\", \"data_collection_time\": \"\", \"src_host\": \"\"}"
	SendMessage(message, serverAddr)

}

func TestUdpAptSip(t *testing.T) {
	serverAddr := "10.25.16.41:5514"
	//<14>"2018-12-23 17:17:42|!secevent|!192.3.221.70|!{\"eventDes\": \"${url}存在黑链\", \"src_ip\":\"0.0.0.0\", \"eventKey\": \"coengine|102001|www.autotest.com\", \"infosecuritysub\": \"信息篡改事件\"}
	message := "<14>\"2018-12-23 17:17:42|!secevent|!192.3.221.70|!{\"eventDes\": \"${url}存在黑链\", \"src_ip\":\"0.0.0.0\", \"eventKey\": \"coengine|102001|www.autotest.com\", \"infosecuritysub\": \"信息篡改事件\"}"
	for i := 0; i < 10; i++ {
		SendMessage(message, serverAddr)
	}
}

func TestAudit1(t *testing.T) {
	serverAddr := "10.25.16.130:1514"
	message := "Oct 24 14:34:36 netvine/10.25.10.120 {\"logType\":15,\"time\":\"2023-12-13 10:05:41\",\"message\":{\"data_type\":\"EVENT_NETWORK_ASSENT\",\"payload\":{\"dev_ipv4\":\"192.168.122.1\",\"dev_mac\":\"52:54:00:a9:a6:7e\",\"dns_info\":[{\"class_index\":1,\"data\":\"LOCALHOST._device-info._tcp.local\",\"domain\":\"_device-info._tcp.local\",\"ttl\":10,\"type\":\"PTR\",\"type_index\":12},{\"class_index\":1,\"data\":\"model=MacSamba\",\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"TXT\",\"type_index\":16},{\"class_index\":1,\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"SRV\",\"type_index\":33}],\"protocol_type\":\"mdns\"}}}"
	for i := 0; i < 100; i++ {
		SendMessage(message, serverAddr)
	}
}
func TestAudit2(t *testing.T) {
	serverAddr := "10.25.16.130:514"
	message := "Oct 24 14:34:36 netvine/10.25.10.120 {\"logType\":15,\"time\":\"2023-12-13 10:05:41\",\"message\":{\"data_type\":\"EVENT_NETWORK_ASSENT\",\"payload\":{\"dev_ipv4\":\"192.168.122.1\",\"dev_mac\":\"52:54:00:a9:a6:7e\",\"dns_info\":[{\"class_index\":1,\"data\":\"LOCALHOST._device-info._tcp.local\",\"domain\":\"_device-info._tcp.local\",\"ttl\":10,\"type\":\"PTR\",\"type_index\":12},{\"class_index\":1,\"data\":\"model=MacSamba\",\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"TXT\",\"type_index\":16},{\"class_index\":1,\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"SRV\",\"type_index\":33}],\"protocol_type\":\"mdns\"}}}"
	for i := 0; i < 100; i++ {
		SendMessage(message, serverAddr)
	}
}

func TestAudit3(t *testing.T) {
	serverAddr := "10.25.16.130:1514"
	message := "{\"logType\":15,\"time\":\"2023-12-13 10:05:41\",\"message\":{\"data_type\":\"EVENT_NETWORK_ASSENT\",\"payload\":{\"dev_ipv4\":\"192.168.122.1\",\"dev_mac\":\"52:54:00:a9:a6:7e\",\"dns_info\":[{\"class_index\":1,\"data\":\"LOCALHOST._device-info._tcp.local\",\"domain\":\"_device-info._tcp.local\",\"ttl\":10,\"type\":\"PTR\",\"type_index\":12},{\"class_index\":1,\"data\":\"model=MacSamba\",\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"TXT\",\"type_index\":16},{\"class_index\":1,\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"SRV\",\"type_index\":33}],\"protocol_type\":\"mdns\"}}}"
	for i := 0; i < 100; i++ {
		SendMessage(message, serverAddr)
	}
}
func TestAudit4(t *testing.T) {
	serverAddr := "10.25.16.130:514"
	message := "{\"logType\":15,\"time\":\"2023-12-13 10:05:41\",\"message\":{\"data_type\":\"EVENT_NETWORK_ASSENT\",\"payload\":{\"dev_ipv4\":\"192.168.122.1\",\"dev_mac\":\"52:54:00:a9:a6:7e\",\"dns_info\":[{\"class_index\":1,\"data\":\"LOCALHOST._device-info._tcp.local\",\"domain\":\"_device-info._tcp.local\",\"ttl\":10,\"type\":\"PTR\",\"type_index\":12},{\"class_index\":1,\"data\":\"model=MacSamba\",\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"TXT\",\"type_index\":16},{\"class_index\":1,\"domain\":\"LOCALHOST._device-info._tcp.local\",\"ttl\":10,\"type\":\"SRV\",\"type_index\":33}],\"protocol_type\":\"mdns\"}}}"
	for i := 0; i < 100; i++ {
		SendMessage(message, serverAddr)
	}
}

func SendMessage(message, serverAddr string) {
	time.Sleep(1 * time.Second)
	//// UDP 服务器地址
	//serverAddr := "10.25.10.88:5514"
	// 创建 UDP 连接
	conn, err := net.Dial("udp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 发送消息
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully!")
}

func TestStr(t *testing.T) {
	str := "log_time, branch_name, reliability, priority, src_ip, dst_ip, event_evidence, event_content, solution, eventKey, branchId, groupId, principle, tag, type, groupName, recordDate, detectEngine, ruleId, firstTime, infosecurity, infosecuritysub, stage, hostRisk, dealStatus, hostName, eventType, domain_name, url, src_mac_addr, src_port, src_host, src_group, dst_mac_addr, dst_port, dst_host, dst_group, msg_count, upstream_traffic, downstream_traffic, attack_count, attack_time, vulnerability_name, vulnerability_description, scan_time, scan_count, scan_success, rsp_email, email_from, email_to, attachment, phishing_link, tampered_url, directory_browsing, password_autofill, server_sensitive_directory, session_token, data_collection_size, data_collection_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	//VALUES分割
	split := strings.Split(str, "VALUES")
	i := strings.Split(split[0], ", ")
	j := strings.Split(split[1], ", ")
	fmt.Println(i)
	fmt.Println(j)
}
