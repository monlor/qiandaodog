package dist

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/carseason/look"
)

type HuaFen struct {
	Credit int    `json:"credit"`
	Url    string `json:url`
}

func (this *CheckIn) Huawei(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)

	if err := cl.Get("https://club.huawei.com/dsu_paulsign-sign.html"); err != nil {
		log.Println("Check in Huawei :", err)
		return
	}

	formhash := string(RegexpSubmatch(cl.Body(), `name="formhash" value="([^"]+)`))
	valueForm := url.Values{}
	valueForm.Set("operation", "qiandao")
	valueForm.Set("formhash", formhash)

	if err := cl.PostForm("https://club.huawei.com/plugin.php?id=dsu_paulsign:sign", valueForm); err != nil {
		log.Println("Check in Huawei :", err)
		return
	}
	template := HuaFen{}

	if err := json.Unmarshal(cl.Body(), &template); err != nil {
		log.Println("Check in Huawei :", err)
		return
	}
	if template.Credit != 0 {
		log.Println("花粉俱乐部签到成功")
	} else {
		log.Println("花粉俱乐部	签到失败")
	}
}
