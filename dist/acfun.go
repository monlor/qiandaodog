package dist

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/carseason/look"
)

type AcfuncTemplate struct {
	Message string `json:"message"`
}

func (this *CheckIn) Acfun(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	valueForm := url.Values{}
	valueForm.Set("channel", "0")
	valueForm.Set("date", fmt.Sprintf("%d", (time.Now().UnixNano())/1e6))
	if err := cl.PostForm("http://www.acfun.cn/webapi/record/actions/signin", valueForm); err != nil {
		log.Println("Check in Acfunc :", err)
		return
	}
	template := new(AcfuncTemplate)
	if err := json.Unmarshal(cl.Body(), template); err != nil {
		log.Println("Acfun	签到失败")
	}
	log.Println("Acfun	", template.Message)

}
