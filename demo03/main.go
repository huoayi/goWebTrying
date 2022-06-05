package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过详细配置服务器的信息处理请求", r.URL.Path)

	i := m.mul(10, w)
	fmt.Fprintln(w, i)
	fmt.Fprintln(w, string(i))

}
func (m *MyHandler) mul(i int, w http.ResponseWriter) int {
	fmt.Fprintf(w, "这里正在进行+1")
	return i + 1
}
func main() {
	myHandler := MyHandler{}
	//创建server结构，并详细配置里面的字段
	server := http.Server{
		Addr:        ":8090",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}
