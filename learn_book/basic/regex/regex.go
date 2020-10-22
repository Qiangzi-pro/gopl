package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ccmouse@gmail.com@abc.com
email is abc@def.org
email2 is    kkk@qq.com
email3 is ddd@abd.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
}