package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f(w http.ResponseWriter, r *http.Request) {
	str := `<h1 style="color:green">hello word</h1>`
	// 前后护法<h1> </h1>放大字体效果
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/Slyvia/Go", f)
	//无语子 之前网页一直加载不出来 原来是因为 把127写成了172
	http.ListenAndServe("127.0.0.1:1234", nil)

	resp, err := http.Get("https://127.0.0.1:1234/Slyvia/Go")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(body))

}
