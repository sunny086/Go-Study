package main

import "fmt"

func main() {

	//sysLog, err := syslog.Dial("sys", "localhost:1234",
	//	syslog.LOG_WARNING|syslog.LOG_DAEMON, "demotag")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Fprintf(sysLog, "This is a daemon warning with demotag.")
	//sysLog.Emerg("And this is a daemon emergency with demotag.")

	//dial, err := net.Dial("tcp", "127.0.0.1:15555")
	//fmt.Println(err)
	//defer dial.Close()
	//str := string("123")
	//mes := []byte(str)
	//dial.Write(mes)

	//var n1 map[string]int
	//n2 := make(map[string]int)
	//fmt.Println(len(n2))
	//n2["1"] = 1
	//n2["2"] = 1
	//
	//fmt.Println(len(n2))

	sprintf := fmt.Sprintf("%.2f", "1.1")
	fmt.Println(sprintf)

	var a float64 = 1.0000000000000001
	var b float64 = 1.000000000000000001
	var c float64 = 1.000000000000001
	var d float64 = 1.0000000000000000001
	fmt.Println(a == b) //true
	fmt.Println(a > b)  //false
	fmt.Println(c == d) //false
	fmt.Println(c > d)  //true
}
