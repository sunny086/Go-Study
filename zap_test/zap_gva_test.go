package zap_test_test

import (
	"GoTest/zap_test/config"
	"GoTest/zap_test/core"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestZap(t *testing.T) {
	logger := Zap()
	zap.ReplaceGlobals(logger)
	logger.Info("info", zap.String("key", "value"))
	err := errors.New("test error")
	logger.Error("error", zap.Error(err))
}

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := PathExists(config.ZapConfigDefault.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.ZapConfigDefault.Director)
		_ = os.Mkdir(config.ZapConfigDefault.Director, os.ModePerm)
	}

	cores := core.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if config.ZapConfigDefault.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
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
