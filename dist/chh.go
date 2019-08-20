package dist
import (
	"github.com/carseason/look"
	"log"
)

func (this *CheckIn)Chh(cookie string){
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("https://www.chiphell.com/");err != nil {
		log.Println("Check in Chh :", err)
		return
	}
	log.Println("CHH 已发送签到请求")
}