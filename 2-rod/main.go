/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     chromedp-rod-test
 * @Date:        2023-05-30 09:13
 * @Description:
 */

package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"time"
)

func main() {
	wsUrl := "ws://127.0.0.1:9222/devtools/browser/7e82126b-f67e-448b-942d-e2279575e019"

	browser := rod.New().ControlURL(wsUrl).MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	// tab页
	page := browser.MustPage()
	//
	page.MustEvalOnNewDocument(`window.alert = () => {};window.prompt = () => {}`)

	// 阻止请求图片、视频、字体文件等类型

	//
	page.
		MustNavigate("").
		MustWaitLoad().
		MustClose()

	fmt.Println("success")
	time.Sleep(time.Hour)
}
