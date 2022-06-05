package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "通过详细配置服务器的信息处理请求", r.URL.Path)

	/*
		Fprintln采用默认格式将其参数格式化并写入w。
		总是会在相邻参数的输出之间添加空格并在输出结束后添加换行符。
		返回写入的字节数和遇到的任何错误。
	*/

	i := m.mul(10, w)
	fmt.Fprintln(w, i)
	//问题：如下方式使用Fprintln在浏览器访问localhost:8090会下载一个文件
	//i1 := string(i)
	//fmt.Fprintln(w, i1)
	fmt.Fprintln(w, "hoahwfnalfhia", string(i))

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
