package dist

import (
	"log"
	"net/url"

	"github.com/carseason/look"
)

func (this *CheckIn) Pojie52(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	valueForm := url.Values{}
	valueForm.Set("mod", "task")
	valueForm.Set("do", "apply")
	valueForm.Set("id", "2")
	if err := cl.PostForm("https://www.52pojie.cn/home.php?", valueForm); err != nil {
		log.Println("Check in Pojie52 :", err)
		return
	}
	log.Println("吾爱破解论坛 已发送签到请求")
}
