package main

import (
	"flag"
	"log"
	"net/http"
)

//var (
//	StartCmd = &cobra.Command{
//		Use:     "websocket-server",
//		Short:   "start the websocket-server",
//		Example: "go-admin websocket-server",
//		Run: func(cmd *cobra.Command, args []string) {
//			Run()
//		},
//	}
//)
var addr = flag.String("addr", ":8088", "http service address")

//func Run() {
//	flag.Parse()
//	go h.run()
//	http.HandleFunc("/", serveHome)
//	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
//		serveWs(w, r)
//	})
//	err := http.ListenAndServe(*addr, nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
