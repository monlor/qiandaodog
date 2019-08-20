package dist

import (
	"log"

	"github.com/carseason/look"
)

func (this *CheckIn) Koolshare(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("http://koolshare.cn/forum.php"); err != nil {
		log.Println("Check in Chh :", err)
		return
	}
	log.Println("Koolshare 已发送签到请求")
}
