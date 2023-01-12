package zap_test_test

import (
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func TestZap1(t *testing.T) {
	InitLogger1()
	defer logger.Sync()
	simpleHttpGet1("www.baidu.com")
	simpleHttpGet1("http://www.baidu.com")
}

var logger *zap.Logger

func InitLogger1() {
	logger, _ = zap.NewProduction()
}
func simpleHttpGet1(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

// =====================================================================================================================
//func TestZap2(t *testing.T) {
//
//}
//
//func InitLogger() {
//	writeSyncer := getLogWriter()
//	encoder := getEncoder()
//	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
//	logger := zap.New(core)
//	sugarLogger = logger.Sugar()
//}
//func getEncoder() zapcore.Encoder {
//	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
//}
//func getLogWriter() zapcore.WriteSyncer {
//	file, _ := os.Create("./test.log")
//	return zapcore.AddSync(file)
//}
