package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"my_code/UploadAndDownload/meta"
	"my_code/UploadAndDownload/util"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("failed to get data %s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newfile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("failed to create file %s\n", err.Error())
			return
		}
		defer newfile.Close()

		//内存中的文件拷贝到文件
		fileMeta.FileSize, err = io.Copy(newfile, file)
		if err != nil {
			fmt.Printf("failed to save data into file %s", err.Error())
			return
		}

		newfile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newfile)
		meta.UpdataFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func SucUploadHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload successfully!")
}

//获取文件原信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Description", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}
