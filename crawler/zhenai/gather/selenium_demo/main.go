package selenium_demo

import (
	"fmt"
	"github.com/tebeka/selenium"
	"io"
	"strings"
	"time"
)

const (
	//设置常量 分别设置chromedriver的地址和本地调用端口
	seleniumPath = "/Users/yuqiang/go/bin/chromedriver"
	port         = 9515
)

func GetPageContent() io.Reader {
	//1.开启selenium服务
	//设置selium服务的选项,设置为空。根据需要设置。
	ops := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//延迟关闭服务
	defer service.Stop()

	//2.调用浏览器
	//设置浏览器兼容性，我们设置浏览器名称为chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//调用浏览器urlPrefix: 测试参考：DefaultURLPrefix = "http://127.0.0.1:4444/wd/hub"
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}
	//延迟退出chrome
	defer wd.Quit()

	//3.对页面元素进行操作
	//获取百度页面
	if err := wd.Get("http://www.zhenai.com/zhenghun"); err != nil {
		panic(err)
	}
	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		for {
			dg, err := wd.FindElement(selenium.ByCSSSelector, "body > div.DIALOG.dialog.active")
			if err == nil {
				return true, nil
			}
			time.Sleep(10 * time.Millisecond)
			fmt.Println(dg)
		}
	})

	wd.Wait(func(wd selenium.WebDriver) (bool, error) {
		for {
			dialog, err := wd.FindElement(selenium.ByCSSSelector, "body > div.DIALOG.dialog.active > div > div.cancel-icon > img")
			err = dialog.Click()
			if err == nil {
				return true, nil
			}
		}
	})

	time.Sleep(3 * time.Second)

	//找到百度输入框id
	we, err := wd.FindElement(selenium.ByCSSSelector, "#app > div.tab-box > div > div:nth-child(2)")
	if err != nil {
		panic(err)
	}
	//向输入框发送“”
	err = we.Click()
	if err != nil {
		panic(err)
	}

	//fmt.Println(wd.PageSource())
	content, err := wd.PageSource()
	if err != nil {
		panic(err)
	}

	r := strings.NewReader(content)
	return r

	//睡眠20秒后退出
	//time.Sleep(20 * time.Second)
}
