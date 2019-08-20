package dist

import (
	"log"

	"github.com/carseason/look"
)

func (this *CheckIn) Hdpfans(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Referers", "http://www.hdpfans.com/")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("http://www.hdpfans.com/qiandao"); err != nil {
		log.Println("Check in Hdpfans :", err)
		return
	}
	formhash := RegexpSubmatch(cl.Body(), `type="hidden" name="formhash" value="([^"]+)`)
	if err := cl.Get("http://www.hdpfans.com/plugin.php?id=k_misign:sign&operation=qiandao&formhash=" + string(formhash) + "&format=empty&inajax=1&ajaxtarget=JD_sign"); err != nil {
		log.Println("Check in Hdpfans :", err)
		return
	}
	log.Println("高清范论坛	已发送签到请求")
}
