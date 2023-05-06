package main

import (
	"net/http"
	"runtime"
	"strconv"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/go", func(w http.ResponseWriter, r *http.Request) {
		num := strconv.FormatInt(int64(runtime.NumGoroutine()), 10)
		w.Write([]byte(num))
	})
	http.ListenAndServe("127.0.0.1:6061", mux)
	http.ListenAndServe("127.0.0.1:6060", nil)

}
