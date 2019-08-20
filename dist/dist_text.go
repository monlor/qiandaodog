package dist

import (
	"regexp"
	"strings"
)

type CheckIn struct{}

/******************正则字符串查找字符串,返回数组*****************/
func RegexpFindAll(content, rule string) []string { //Regexp0
	r1 := regexp.MustCompile(rule)
	r2 := r1.FindAllString(content, -1)
	return r2
}

/******************正则查找匹配内容*****************/
func RegexpSubmatch(content []byte, rule string) []byte { //Regexp3,s2为正文，s1为规则
	r1 := regexp.MustCompile(rule).FindSubmatch(content)
	if len(r1) == 0 {
		return nil
	}
	return r1[1]
}

/******************正则查找内容，返回第一个匹配到的*****************/
func RegexFindString(content, rule string) string { //Regexp2
	r1 := regexp.MustCompile(rule)
	r2 := r1.FindString(content)
	return r2
}

/******************正则替换*****************/
func RegexpReplaceAllString(content, rule, key string) string { //Regexp1
	r1 := regexp.MustCompile(rule)
	r2 := r1.ReplaceAllString(content, key)
	return r2
}

func Replace(s1, s2 string) string {
	return strings.Replace(s1, s2, "", -1)

}
