package string

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestUsbLicense(t *testing.T) {
	LicenseSplit()
	filePathJoin()
}

func LicenseSplit() {
	var decrypt = "SN=License-Test#prod=USB;project=USB;pid=PID-USB-1U;vid=VID-USB-WANGTENGTECH;sku=;FUC_BASIC=2;FUC_UPGRADE_VIRUS_DB=2;FUC_UPGRADE_SOFTWARE=2;FUC_ENCRYPT=2;apply_date=1658678400;expire_date=1661270400;effective_time=2592000;type=0;uuid=f09c6cd4-35a5-472c-ac2a-726650094c9d;#signature=83a2bdeb126d18ab8a4de3b619339cee76d065447e3e67e0c6f9d21e9e29597a"
	split := strings.Split(decrypt, "#")
	contents := strings.Split(split[1], ";")
	splitExpireDate := strings.Split(split[1], ";expire_date=")
	splitEffectiveTime := strings.Split(split[1], ";effective_time=")
	splitBasicList := strings.Split(split[1], "FUC_BASIC=")
	splitUpgradeVirusDbList := strings.Split(split[1], "FUC_UPGRADE_VIRUS_DB=")
	splitUpgradeSoftwareList := strings.Split(split[1], "FUC_UPGRADE_SOFTWARE=")
	splitEncryptList := strings.Split(split[1], "FUC_ENCRYPT=")
	splitType := strings.Split(split[1], ";type=")
	splitUuid := strings.Split(split[1], ";uuid=")
	fmt.Println(split)
	fmt.Println(contents)
	fmt.Println(splitExpireDate)
	fmt.Println(splitEffectiveTime)
	fmt.Println(splitBasicList)
	fmt.Println(splitUpgradeVirusDbList)
	fmt.Println(splitUpgradeSoftwareList)
	fmt.Println(splitEncryptList)
	fmt.Println(splitType)
	fmt.Println(splitUuid)
}

func filePathJoin() {
	filePath := filepath.Join("data", "safe_auth", "safe_auth_deviceid.yml")
	fmt.Println(filePath)
	join := strings.Join([]string{"data", "safe_auth", "safe_auth_deviceid.yml"}, "/")
	fmt.Println(join)
	join2 := strings.Join([]string{"SN", "OM10DX04SH2104", "#project", "USB", "product" + "@V1.0(1.5.571)", "#pid", "PID-USB-1U", "#vid", "VID-USB-WANGTENGTECH", "#sku"}, "=")
	fmt.Println(join2)
}
