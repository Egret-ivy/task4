package main

import (
	"bufio"
	"net"
	"fmt"
	"os"
	"strings"
)

func main() {
	//1.与ser端建立连接
	conn,err := net.Dial("tcp","127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed,err :",err)
		return
	}

	//2.发送数据
	//var msg string
	//if len(os.Args) < 2 {
	//	msg = "hello Slyvia!"
	//}else {
	//	//Args[0]返回的是程序名字
	//	msg = os.Args[1] //返回在命令行输入的第一个参数
	//}

		reader := bufio.NewReader(os.Stdin) //从标准输入创造读对象
	for {
		fmt.Println("坦白从宽:")
		//fmt.Scanln(&msg) 遇到空格会停止读入
		msg,_ := reader.ReadString('\n') //读到换行
		msg = strings.TrimSpace(msg) //删除前后空格
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
