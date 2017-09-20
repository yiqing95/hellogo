package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		// 	IdleTimeout:  time.Minute, // 新版本添加的
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	timer := time.NewTimer(time.Second * 3)

	select {
	case <-timer.C:
		log.Println("Done")
	case <-r.Context().Done(): // context包中的东东 可以取消处理哦！ curl 请求时 ctrl+c 关闭客户端后 服务端就被取消了！
		log.Println("cancelled")
		return // 终止流程
	}

	w.WriteHeader(200)
}
