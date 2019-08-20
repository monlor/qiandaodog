/*#Comment
 *@Description: qiandaodog 2.1
 *@MethodAuthor:  Carseason
 *@Date: 2019-08-20 01:07:44
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"qiandaodog/customize"
	"qiandaodog/dist"
)

type Cookies struct {
	Name   string
	Cookie string
}
type Customize struct {
	Name   string
	Cookie string
	Url    string
	Types  string
	Ua     string
	Value  []struct {
		Name  string
		Value string
	}
}

func Openfile(name string) []byte {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		os.Exit(0)
	}
	return buf
}

func main() {
	PackGo()
	CustomizeGo()
}

func PackGo() {
	file := Openfile("cookie.txt")
	value := []Cookies{}
	if err := json.Unmarshal(file, &value); err != nil {
		fmt.Println("Cookie 格式异常:", err)
		os.Exit(0)
	}
	qd := new(dist.CheckIn)
	for v := range value {
		if value[v].Cookie == "" {
			continue
		}
		switch value[v].Name {
		case "acfun": //A站
			qd.Acfun(value[v].Cookie)
		case "baidu": //百度贴吧及文库
			qd.Baidu(value[v].Cookie)
		case "v2ex": //V2EX
			qd.V2ex(value[v].Cookie)
		case "hostloc": //hostloc
			qd.Hostloc(value[v].Cookie)
		case "163music": //网易云音乐
			qd.Music163(value[v].Cookie)
		case "miui": //miui论坛
			qd.Miui(value[v].Cookie)
		case "52pojie": //吾爱破解
			qd.Pojie52(value[v].Cookie)
		case "kafan": //卡饭
			qd.Kafan(value[v].Cookie)
		case "smzdm": //什么值得买
			qd.Smzdm(value[v].Cookie)
		case "zimuzu": //人人字幕组
			qd.Zimuzu(value[v].Cookie)
		case "gztown": //港知堂社区PT站
			qd.Gztown(value[v].Cookie)
		case "meizu": //魅族论坛
			qd.Meizu(value[v].Cookie)
		case "hdpfans": //高清范论坛
			qd.Hdpfans(value[v].Cookie)
		case "chh": //CHH
			qd.Chh(value[v].Cookie)
		case "koolshare": //Koolshare
			qd.Koolshare(value[v].Cookie)
		case "right": //恩山
			qd.Right(value[v].Cookie)
		case "huawei": //花粉俱乐部
			qd.Huawei(value[v].Cookie)
		}

	}
}

func CustomizeGo() {
	file := Openfile("customize.txt")
	value := []Customize{}
	if err := json.Unmarshal(file, &value); err != nil {
		fmt.Println("自定义签到 格式异常:", err)
		os.Exit(0)
	}
	cu := customize.CustomizeCheck{}
	for v := range value {
		if value[v].Cookie == "" || value[v].Url == "" {
			continue
		}
		if value[v].Ua == "" {
			value[v].Ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36"
		}
		if value[v].Types == "Post" || value[v].Types == "POST" || value[v].Types == "post" {
			if len(value[v].Value) > 0 {
				valueForm := url.Values{}
				for k := range value[v].Value {
					valueForm.Set(value[v].Value[k].Name, value[v].Value[k].Value)
				}
				cu.PostForm(value[v].Name, value[v].Url, value[v].Cookie, value[v].Ua, valueForm)
			} else {
				cu.Post(value[v].Name, value[v].Url, value[v].Cookie, value[v].Ua)
			}
		} else {
			cu.Get(value[v].Name, value[v].Url, value[v].Cookie, value[v].Ua)
		}

	}

}
