package dist

import (
	"encoding/json"
	"log"

	"github.com/carseason/look"
)

type MeizuTemplate struct {
	Code    int    `json:code`
	Message string `json:message`
}

func (this *CheckIn) Meizu(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("https://bbs-act.meizu.cn/index.php?mod=signin&action=sign"); err != nil {
		log.Println("Check in Meizu :", err)
		return
	}
	template := MeizuTemplate{}
	if err := json.Unmarshal(cl.Body(), &template); err != nil {
		log.Println("Check in Meizu :", err)
		return
	}
	log.Println("魅族论坛", template.Message)
}
