package log

import (
	"fmt"
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
) // 防止标准库的log和我们的文件夹重名发生冲突

// 自定义一个fileLog类型, 将日志写入文件
type fileLog string

// ok
var log *stlog.Logger // Logger对象
// 向文件中写入日志
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600) // 分别表示如果没有文件创建, 向文件后写, 只写
	if err != nil {
		return 0, err // 表示出错
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("err !")
		}
	}(f) // 关闭文件
	return f.Write(data)
}

// 接受一个目标文件路径作为参数，用于初始化日志记录器。
func Run(destination string) {
	log = stlog.New(fileLog(destination), "go: ", stlog.LstdFlags)
}

// 定义了一个RegisterHandlers函数，用于注册HTTP请求处理器。
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

// 写入日志
func write(message string) {
	log.Printf("%v\n", message)
}
