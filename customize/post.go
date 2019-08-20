package customize

import (
	"log"
	"net/url"
	"qiandaodog/dist"

	"github.com/carseason/look"
)

func (this *CustomizeCheck) Post(name, url, cookie, ua string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", ua)
	cl.SetHeader("Cookie", cookie)
	referer := dist.RegexFindString(url, `^https?://[^/]+`)
	if referer != "" {
		cl.SetHeader("Referer", referer)
	}

	if err := cl.Post(url); err != nil {
		log.Println("Check in "+name+" :", err)
		return
	}
	log.Println("对自定义签到 :" + name + " 已发送请求")
	return
}

func (this *CustomizeCheck) PostForm(name, url, cookie, ua string, value url.Values) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", ua)
	cl.SetHeader("Cookie", cookie)
	referer := dist.RegexFindString(url, `^https?://[^/]+`)
	if referer != "" {
		cl.SetHeader("Referer", referer)
	}
	if err := cl.PostForm(url, value); err != nil {
		log.Println("Check in "+name+" :", err)
		return
	}
	log.Println("对自定义签到 :" + name + " 已发送请求")
	return
}
