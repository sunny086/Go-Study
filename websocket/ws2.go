package main

//
//import (
//	"net/http"
//	"time"
//
//	"github.com/gorilla/websocket"
//	"github.com/student/1330/impl"
//)
//
//var (
//	upgrade = websocket.Upgrader{
//		//允许跨域
//		CheckOrigin: func(r *http.Request) bool {
//			return true
//		},
//	}
//)
//
//func wsHandler(w http.ResponseWriter, r *http.Request) {
//	var (
//		//websocket 长连接
//		wsConn *websocket.Conn
//		err    error
//		conn   *impl.Connection
//		data   []byte
//	)
//	//header中添加Upgrade:websocket
//	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
//		return
//	}
//
//	if conn, err = impl.InitConnection(wsConn); err != nil {
//		goto ERR
//	}
//
//	go func() {
//		var (
//			err error
//		)
//		for {
//			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
//				return
//			}
//			time.Sleep(time.Second * 1)
//		}
//	}()
//
//	for {
//		if data, err = conn.ReadMessage(); err != nil {
//			goto ERR
//		}
//		if err = conn.WriteMessage(data); err != nil {
//			goto ERR
//		}
//	}
//
//ERR:
//	conn.Close()
//}
//
//func main() {
//	//http标准库
//	http.HandleFunc("/ws", wsHandler)
//	http.ListenAndServe("0.0.0.0:5555", nil)
//}
