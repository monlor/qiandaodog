package dist

import (
	"log"

	"github.com/carseason/look"
)

type Gztown struct{}

func (this *CheckIn) Gztown(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("https://pt.gztown.net/attendance.php"); err != nil {
		log.Println("Check in Gztown :", err)
		return
	}
	log.Println("港知堂社区PT站	已发送签到请求")
}
