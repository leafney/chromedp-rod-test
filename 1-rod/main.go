/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     chromedp-rod-test
 * @Date:        2023-05-30 09:08
 * @Description:
 */

package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"strings"
	"time"
)

func main() {

	wsUrl := "ws://127.0.0.1:9222/devtools/browser/7e82126b-f67e-448b-942d-e2279575e019"

	browser := rod.New().ControlURL(wsUrl).MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	//var sessionReqID proto.NetworkRequestID
	//go page.EachEvent(
	//	func(e *proto.NetworkRequestWillBeSent) {
	//		if e.Request.URL == "https://www.kuaidi100.com/query" {
	//			sessionReqID = e.RequestID
	//		}
	//
	//		if e.Type == proto.NetworkResourceTypeXHR {
	//
	//		}
	//
	//	}, func(e *proto.NetworkResponseReceived) {
	//		//if e.RequestID==sessionReqID{
	//		//
	//		//}
	//		//	e.Response.URL
	//
	//	},
	//)()

	router := browser.HijackRequests()
	defer router.MustStop()
	router.MustAdd("*", func(ctx *rod.Hijack) {
		ctx.MustLoadResponse()

		resType := ctx.Request.Type()
		if resType == proto.NetworkResourceTypeXHR {
			u := ctx.Request.URL().String()
			fmt.Println("请求： ", u)
			if strings.HasPrefix(u, "https://www.kuaidi100.com/query") {
				//ctx.MustLoadResponse()

				body := ctx.Response.Body()
				fmt.Println("响应 ", body)
			}
		}

		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})

	go router.Run()

	page := browser.MustPage("https://www.kuaidi100.com/")
	page.MustWaitLoad()
	utils.Sleep(10)

	wait := page.MustWaitRequestIdle()

	// 输入
	page.MustElement("#input").MustInput("JT3032192857289")

	// 模拟手动输入
	//page.MustElement("#input").MustFocus()
	//page.MustInsertText("300535118665")

	page.MustElement("#query").MustClick()
	wait()

	//page.MustWaitLoad().MustScreenshot("a.png")
	utils.Sleep(10)
	//page.MustClose()
	fmt.Println("success")
	time.Sleep(time.Hour)

}
