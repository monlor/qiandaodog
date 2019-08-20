package dist

import (
	"log"

	"github.com/carseason/look"
)

func (this *CheckIn) V2ex(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("https://www.v2ex.com/mission/daily"); err != nil {
		log.Println("Check in V2ex :", err)
		return
	}
	status := RegexFindString(cl.Content(), `(<input type="button" class="super normal button" value="领取 X 铜币" onclick="location.href = '[^']+|每日登录奖励已领取)`)
	switch status {
	case "每日登录奖励已领取":
		log.Println("V2ex	每日登录奖励已领取过")
	case "":
		log.Println("V2EX	未登录")
	default:
		url := "https://www.v2ex.com" + RegexFindString(status, `[^']+$`)
		if err := cl.Get(url); err != nil {
			log.Println("V2EX	签到失败")
			return
		}
		log.Println("V2EX	签到成功")
	}
}
