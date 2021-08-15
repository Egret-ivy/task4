package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f(w http.ResponseWriter, r *http.Request) {

	// 前后护法<h1> </h1>放大字体效果
	//b, err := ioutil.ReadFile("081121243185.html")
	//b, err := ioutil.ReadFile("hello.html")
	b, err := ioutil.ReadFile("bilibili.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/Slyvia/Go", f)
	//无语子 之前网页一直加载不出来 原来是因为 把127写成了172
	http.ListenAndServe("127.0.0.1:4321", nil)
}
