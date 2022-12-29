package zap

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestZap(t *testing.T) {

}

func Zap1() (logger *zap.Logger) {
	if ok, _ := PathExists("./log"); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", "./log")
		_ = os.Mkdir("./log", os.ModePerm)
	}
	//cores := internal.Zap.GetZapCores()
	//logger = zap.New(zapcore.NewTee(cores...))
	//
	//if global.NETVINE_CONFIG.Zap.ShowLine {
	//	logger = logger.WithOptions(zap.AddCaller())
	//}
	return logger
}

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
