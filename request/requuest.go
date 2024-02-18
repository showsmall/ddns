package request

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Req(url string) (bool, []byte) {
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("请求接口失败", err)
		return false, nil
	}
	defer resp.Body.Close()
	red, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Println("请求接口返回数据异常：", resp.StatusCode, string(red))
		return false, nil
	}
	log.Println("对外请求接口成功：", resp.StatusCode, string(red))
	return true, red
}
