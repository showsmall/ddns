package logic

import (
	"encoding/json"
	"fmt"
	"log"
	"sre/tools/config"
	"sre/tools/request"
)

func Update_wan_ip() {
	// 获取当前公网出口ip
	status, req := request.Req(config.C.Get_wan_ip_url)
	if status != true {
		log.Println("获取公网ip失败，详情查询日志")
		return
	}
	// 解析获取到的公网ip信息到struct中
	wan_info := config.ReqData{}
	err := json.Unmarshal(req, &wan_info)
	if err != nil {
		log.Println("公网iP获取的数据转struct失败:", err, string(req))
		return
	}

	// 判断是否拿到了公网
	if wan_info.IP == "" {
		log.Println("公网iP获取为空", string(req))
		return
	}

	// 更新域名的wan ip
	url := fmt.Sprintf("%s%s", config.C.Update_ipv4_url, wan_info.IP)
	s, _ := request.Req(url)
	if s != true {
		log.Println("更新公网ip失败，详情查询日志")
		return
	}
	log.Println("更新公网ip成功")
	return
}
