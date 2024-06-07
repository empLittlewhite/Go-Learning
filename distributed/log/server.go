// 日志服务式一个Web服务，可以接受post请求，
// 将POST请求中的内容写入到log中

package log

// 给log库取名stlog
import (
	stlog "log"
	"net/http"
	"os"
)

// log.Logger是一个用于记录日志的工具。它是标准库中log包提供的一个类型，
// 用于向应用程序的控制台、文件或其他输出目标输出日志信息
var log *stlog.Logger

// 将日志写入文件系统,就是string的别名
type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	//文件不存在，需要先创建os.O_create,设置的模式是write-only,
	// 数据是直接往后加的,保证不同系统之间的连贯性-0600
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	// 函数执行完毕：最后关闭文件，比return还要更后面呢
	defer f.Close()
	// 将数据写入到file中
	return f.Write(data)
}

// 将log指向某个文件地址
func Run(destination string) {
	// stlog.New()会创建一个新的logger
	// LstdFlags：日期+时间
	log = stlog.New(fileLog(destination), "go", stlog.LstdFlags)
}

func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request)){
		//判断类型
		switch r.Method{
		case http.MethodPost:
			// 读取body中的内容-->出错或者msg位0-->错误请求，立刻终止
			msg,err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) = 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
			// 默认情况下禁止了对根路径 / 的处理，
			// 可以增加网站的安全性，防止未经授权的访问。
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func write(message string ) {
	// 将msg写入到日志中
	log.Printf("%v\n",message)
}