package customize

import (
	"log"
	"qiandaodog/dist"

	"github.com/carseason/look"
)

type CustomizeCheck struct {
}

func (this *CustomizeCheck) Get(name, url, cookie, ua string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", ua)
	cl.SetHeader("Cookie", cookie)
	referer := dist.RegexFindString(url, `^https?://[^/]+`)
	if referer != "" {
		cl.SetHeader("Referer", referer)
	}
	if err := cl.Get(url); err != nil {
		log.Println("Check in "+name+" :", err)
		return
	}
	log.Println("对自定义签到 :" + name + " 已发送请求")
	return
}
