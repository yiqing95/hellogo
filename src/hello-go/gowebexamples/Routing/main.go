package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// 创建多路路由器
	r := mux.NewRouter()
	// 像http标准包那样注册 请求处理器  不过注册的对象变为路由器啦！
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// 解析请求中的变量 得到一个映射map
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You 've request the book:%s on page %s\n", title, page)
	})

	// ------------------------------------------------------ +|
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
	// ------------------------------------------------------- +|

	// 设置http服务器的路由器
	http.ListenAndServe(":80", r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi create"))
}
func ReadBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi read"))
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi update"))
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi delete"))
}
