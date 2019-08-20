package dist

import (
	"log"
	"net/url"

	"github.com/carseason/look"
)

type Kafan struct{}

func (this *CheckIn) Kafan(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("https://bbs.kafan.cn/"); err != nil {
		log.Println("Check in Kafan :", err)
		return
	}
	formhash := string(RegexpSubmatch(cl.Body(), `name="formhash" value="([^"]+)`))
	valueForm := url.Values{}
	valueForm.Set("id", "dsu_amupper")
	valueForm.Set("ppersubmit", "true")
	valueForm.Set("formhash", formhash)
	valueForm.Set("infloat", "yes")
	valueForm.Set("handlekey", "dsu_amupper")
	valueForm.Set("inajax", "1")
	valueForm.Set("ajaxtarget", "fwin_content_dsu_amupper")
	if err := cl.PostForm("https://bbs.kafan.cn/plugin.php", valueForm); err != nil {
		log.Println("卡饭论坛签到失败!")
		return
	}
	log.Println("卡饭论坛	已发送签到请求")
}
