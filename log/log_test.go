package log

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

// Go语言提供的基本日志功能
func TestLog(t *testing.T) {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
	simpleHttpGet("baidu.com")
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("http://www.baidu.com")
}

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}

// Only difference between those in terms of its printing behaviour is
// - `log.Println` writes to `Stderr`
// - `fmt.Println` writes to `Stdout`
// Both are not buffered. So the fact that `StdOut` came before `StdError` is specific to your terminal or environment.
func TestLogPrintln(t *testing.T) {
	var a string = "initail"
	log.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
}
