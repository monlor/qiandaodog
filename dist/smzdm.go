package dist

import (
	"encoding/json"
	"log"

	"github.com/carseason/look"
)

type SmzdmTemplate struct {
	Error_code int    `json:error_code`
	Error_msg  string `json:error_msg`
	Data       struct {
		Add_point int    `json:add_point`
		Slogan    string `json:slogan`
	} `json:"data"`
}

func (this *CheckIn) Smzdm(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Referer", "https://www.smzdm.com/")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("https://zhiyou.smzdm.com/user/checkin/jsonp_checkin"); err != nil {
		log.Println("Check in Smzdm :", err)
		return
	}
	template := SmzdmTemplate{}
	if err := json.Unmarshal(cl.Body(), &template); err != nil {
		log.Println("Check in Smzdm :", err)
		return
	}

	if template.Error_code == 0 {
		log.Println("什么值得买	", RegexpReplaceAllString(template.Data.Slogan, `<[^>]+>`, ""))
	} else {
		log.Println("什么值得买	", template.Error_msg)
	}

}
