package dist

import (
	"fmt"
	"log"
	"strings"

	"github.com/carseason/look"
)

func (this *CheckIn) Baidu(cookie string) {
	this.Tieba(cookie)
	this.Wenku(cookie)
}
func (this *CheckIn) Tieba(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	cl.SetHeader("Referer", "https://wapp.baidu.com/")
	if err := cl.Get("https://wapp.baidu.com/mo/q----,sz@320_240-1-3---2/m?tn=bdFBW&tab=favorite"); err != nil {
		log.Println("Check in Tieba :", err)
		return
	}
	links := RegexpFindAll(cl.Content(), `>\d+\.<a href="[^"]+">[^<]+`)
	if len(links) == 0 {
		log.Println("未关注贴吧或百度Cookie错误")
		return
	}
	fmt.Println("--------------------------------正在签到百度贴吧：--------------------------------")
	var number int //签到计数器
	for v := range links {
		url := "https://" + string(RegexpSubmatch([]byte(links[v]), `"//([^"]+)`))
		name := RegexFindString(links[v], `[^>]+$`)
		if err := cl.Get(url); err != nil {
			continue
		}
		if strings.Contains(cl.Content(), ">签到<") == true {
			/**************筛选出签到链接******************/
			tmpUrl := string(RegexpSubmatch(cl.Body(), `style="text-align:right;"><a href="([^"]+)`))
			urlChild := "https://tieba.baidu.com" + strings.Replace(tmpUrl, `amp;`, "", -1)
			if err := cl.Get(urlChild); err != nil {
				log.Println("Check in Tieba-"+name+": ", err)
				continue
			}
			log.Println("Check in Tieba-" + name + ": 	签到成功")
			number++
		} else if strings.Contains(cl.Content(), ">已签到<") {
			log.Println("Check in Tieba-" + name + ": 	已签到")
		} else {
			log.Println("Check in Tieba-" + name + ": 	签到失败")
		}
	}
	fmt.Println("--------------------------------百度贴吧签到完成,本次签到", number, "个吧")
}

func (this *CheckIn) Wenku(cookie string) {
	cl := look.NewLook()
	cl.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
	cl.SetHeader("Cookie", cookie)
	cl.SetHeader("Referer", "https://wenku.baidu.com/task/browse/daily")
	fmt.Println("--------------------------------正在签到百度文库：--------------------------------")
	if err := cl.Get("https://wenku.baidu.com/task/submit/signin"); err != nil {
		log.Println("Check in Wenku :", err)
		return
	}
	if err := cl.Get("https://tanbi.baidu.com/home/task/taskIndex"); err != nil {
		log.Println("Check in Wenku :", err)
		return
	}
	if day := string(RegexpSubmatch(cl.Body(), `<div class="countIcon">([^<]+)</div>`)); day != "" {
		log.Println("百度文库已连续签到", day, "天")
	} else {
		log.Println("百度文库签到失败")
	}
	return
}
