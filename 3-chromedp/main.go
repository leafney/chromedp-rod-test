/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     chromedp-rod-test
 * @Date:        2023-05-30 09:14
 * @Description:
 */

package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"strings"
	"time"
)

func main() {
	wsUrl := "ws://127.0.0.1:9222/devtools/browser/7e82126b-f67e-448b-942d-e2279575e019"

	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), wsUrl)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// 接收响应数据
	responseChan := make(chan string)

	// 设置监听
	chromedp.ListenTarget(ctx, func(event interface{}) {
		if ev, ok := event.(*network.EventResponseReceived); ok {

			if ev.Type != "XHR" {
				return
			}
			rsp := ev.Response

			if strings.HasPrefix(rsp.URL, "https://www.kuaidi100.com/query") {
				go func() {
					if body, err := network.GetResponseBody(ev.RequestID).Do(cdp.WithExecutor(ctx, chromedp.FromContext(ctx).Target)); err == nil {
						data := string(body)
						//fmt.Println("响应： ", data)

						responseChan <- data
						return
					}
				}()
			}
		}
	})

	var x string
	chromedp.Run(ctx,
		chromedp.Navigate("https://www.kuaidi100.com/"),

		chromedp.WaitVisible(`div#news`),

		//chromedp.Evaluate(),
		chromedp.EvaluateAsDevTools(`window.scroll(0,500) || "a";`, &x),

		chromedp.SendKeys(`#input`, "JT3032192857289", chromedp.ByID),

		chromedp.Click(`#query`, chromedp.NodeVisible),

		chromedp.Sleep(2*time.Second),

		// 让输入框失去焦点，隐藏下拉浮层
		chromedp.Click(`#internal`, chromedp.NodeVisible),

		// 关闭当前tab页
		page.Close(),
	)

	fmt.Println(<-responseChan)

	fmt.Println("success")
	time.Sleep(time.Hour)

}
