package main

import (
	"net/http"
	"strconv"

	"github.com/NewsNIng/goctx/context"
)

func main() {
	//启动一个Web服务
	http.Handle("/", http.HandlerFunc(myHander))
	http.ListenAndServe(":1234", nil)
}

//定义一个Hander
func myHander(rw http.ResponseWriter, r *http.Request) {
	//模拟为Request附加值，这里附加了2个
	context.Set(r, "user", "张三")
	context.Set(r, "age", 18)
	//这个模拟一个方法或者函数的调用，大部分情况下可能不在一个包里
	doHander(rw, r)
}
func doHander(rw http.ResponseWriter, r *http.Request) {
	//我们从这个Request里取出对应的值。
	user := context.Get(r, "user").(string)
	age := context.Get(r, "age").(int)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}
