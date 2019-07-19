package main

/**
* http：go语言标准库内建提供的net/http包
*
* 提供的http.ListenAndServe()方法，可以在指定的地址进行监听
* 开启一个http服务，服务该方法的原型如下：
* 	func ListenAndServe(addr string,handler Handler)error
* 	用于在滴定的TCP网络地址addr进行监听，然后调用服务端处理程序来处理传入的连接请求。
*
* 参数： addr--监听地址
* 		handler--标识服务端处理程序，通常为空，意味着服务端调用http.DefaultServeMux
*       进行处理。而服务端编写的业务处理程序http.Handle()或http.HandleFunc()默认注入
*		http.DefaultServeMux中
* 处理https请求：
* func ListenAndServeTLS(addr string,keyFile string,handler Handler)error
*
**/

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	http.ListenAndServe("127.0.0.1:8080",nil)
	//resp, err := http.Get("http://www.baidu.com")
	//resp, err := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	//if err != nil {
	//	panic("the http nil")
	//}
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//
	//defer resp.Body.Close()

}
