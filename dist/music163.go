package dist

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/carseason/look"
)

func (this *CheckIn) Music163(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	valueForm := url.Values{}
	valueForm.Set("type", "1")
	if err := cl.PostForm("http://music.163.com/api/point/dailyTask", valueForm); err != nil {
		log.Println("Check in 163Music :", err)
		return
	}
	var music163 struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	if err := json.Unmarshal(cl.Body(), &music163); err != nil || music163.Code == 301 {
		log.Println("网易云音乐登录失败")
	}
	log.Println("网易云音乐签到", music163.Msg)

}
