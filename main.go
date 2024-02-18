package main

import (
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"sre/tools/config"
	"sre/tools/logic"
)

func Loginit(logfilepath string) *os.File {
	file, err := os.OpenFile(logfilepath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("初始化日志文件失败", err)
		return nil
	}
	//defer file.Close()  // 这里不能调用defer 关闭，关闭后，后面的日志都不能记录了
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(file)
	log.Println("Init logs success")
	return file
}

func main() {
	// 初始化配置文件
	var inifilename string
	flag.StringVar(&inifilename, "inifilename", "", "配置文件绝对路径")
	flag.Parse()
	if inifilename == "" {
		log.Println("配置文件参数缺失： -h 查看帮助")
		return
	}
	err := config.Init(inifilename)
	if err != nil {
		log.Println("初始化配置文件失败", err)
		return
	}
	_ = Loginit(config.C.Run_log_file_path)
	// 新建一个定时任务对象
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	// 添加定时任务,
	_, err = crontab.AddFunc(config.C.Scheduled_tasks, logic.Update_wan_ip)
	if err != nil {
		log.Println("添加定时任务失败：", err)
		return
	}
	// 启动定时器
	crontab.Start()
	select {} //阻塞主线程停止
}
