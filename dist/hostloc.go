package dist

import (
	"log"

	"github.com/carseason/look"
)

func (this *CheckIn) Hostloc(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	if err := cl.Get("https://www.hostloc.com/forum.php"); err != nil {
		log.Println("Check in HostLoc :", err)
		return
	}
	links := []string{"25650", "7436", "22176", "23376", "132", "26477", "25285", "26532", "25728", "26440", "18756", "12368", "26564"}
	for i := 0; i < 12; i++ {
		url := "https://www.hostloc.com/space-uid-" + links[i] + ".html"
		cl.Get(url)
	}
	log.Println("Hostloc 已发送获取积分请求")
	return
}
