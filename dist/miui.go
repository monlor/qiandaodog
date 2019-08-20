package dist

import (
	"log"
	"net/url"

	"github.com/carseason/look"
)

func (this *CheckIn) Miui(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	valueForm := url.Values{}
	valueForm.Set("mod", "sign/index")
	valueForm.Set("op", "sign")
	if err := cl.PostForm("https://www.miui.com/extra.php", valueForm); err != nil {
		log.Println("Check in Miui :", err)
		return
	}
	log.Println("MIUI已发送签到请求!")

}
