package main

import (
	"fmt"
	"my_code/UploadAndDownload/handler"

	//"my_code/UploadAndDownload/meta"
	//"my_code/UploadAndDownload/util"
	"net/http"
)

func main() {
	//模块中要导出的函数，必须首字母大写 函数之前写的uploadHandler报错cannot refer to unexported name
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.SucUploadHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("failed to start server,err :%s", err.Error())
	}

}
