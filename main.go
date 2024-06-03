package main

import "net/http"

// func main() {
// 	// 设置/网页路径的handleFunc，
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello world"))
// 	})
// 	// 设置web服务器
// 	http.ListenAndServe("localhost:8080", nil) //nil是默认defaultServeMux
// }

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World")) //将字符串转换成字节切片（byte slice）
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!")) //将字符串转换成字节切片（byte slice）
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!")) //将字符串转换成字节切片（byte slice）
}
func main() {
	hh := helloHandler{}
	a := aboutHandler{}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, //default
	}
	// http.ListenAndServe("localhost:8080",nil)

	http.Handle("/hello", &hh)
	http.Handle("/about", &a)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home !"))
	})
	// http.HandleFunc("/welcome", Welcome) //传入Welcome函数，不要有Welcome（）--会直接执行
	http.Handle("welcome", http.HandlerFunc(Welcome)) //HandlerFunc将Welcome类型转换了！
	server.ListenAndServe()
	// http.DefaultServeMux
}
