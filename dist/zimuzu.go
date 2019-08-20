package dist

import (
	"log"

	"github.com/carseason/look"
)

func (this *CheckIn) Zimuzu(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("http://www.zimuzu.tv/user/user"); err != nil {
		log.Println("Check in Zimuzu :", err)
		return
	}
	log.Println("人人影视字幕组	已发送签到请求")
}
