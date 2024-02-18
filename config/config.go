package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	Update_ipv4_url   string `ini:"update_ipv4_url"`
	Get_wan_ip_url    string `ini:"get_wan_ip_url"`
	Run_log_file_path string `ini:"run_log_file_path"`
	Scheduled_tasks   string `ini:"scheduled_tasks"`
}

var C = new(Config)

type ReqData struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func Init(inifilenamepaths string) error {

	cfg, err := ini.Load(inifilenamepaths)
	if err != nil {
		fmt.Println("加载配置文件失败：", err)
		return err
	}

	err = cfg.MapTo(C)
	if err != nil {
		fmt.Println("映射到结构体失败：", err)
		return err
	}
	return nil
}
