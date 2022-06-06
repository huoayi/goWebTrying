package main

import (
	"fmt"
	"net/http"
)

//Handler 是一个接口，但是前一小节中的 helloName 函数并没有实现
//ServeHTTP 这个接口，为什么能添加呢？
//原来在 http 包里面还定义了一个类型 HandlerFunc,
//我们定义的函数 helloName 就是这个 HandlerFunc 调用之后的结果，这个类型默认就实现了 ServeHTTP 这个接口，
//即我们调用了 HandlerFunc (f), 强制类型转换 f 成为 HandlerFunc 类型，这样 f 就拥有了 ServeHTTP 方法。

func helloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HelloWorld!", r.URL.Path)
}
func main() {
	//HandleFunc注册一个处理器函数handler和对应的模式pattern。
	http.HandleFunc("/", helloName)
	http.ListenAndServe(":9090", nil)
}
