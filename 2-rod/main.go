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
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"os"
	"time"
)

func main() {
	wsId := "7e0ad1a0-6a63-4c12-b8f1-045c74752767"
	wsUrl := fmt.Sprintf("ws://127.0.0.1:9222/devtools/browser/%s", wsId)

	browser := rod.New().ControlURL(wsUrl).MustConnect().NoDefaultDevice()
	//defer browser.MustClose()

	// tab页
	page := browser.MustPage()
	// 禁止弹窗
	page.MustEvalOnNewDocument(`window.alert = () => {};window.prompt = () => {}`)

	fmt.Println("准备开始")
	utils.Sleep(20)

	// 阻止请求图片、视频、字体文件等类型
	hijackVideoImage(page)

	//
	page.
		MustNavigate("https://www.youtube.com/")
	//MustNavigate("https://www.iplaysoft.com/").
	//MustWaitLoad()
	//MustClose()

	// 等待页面加载完成 方式一
	page.MustWaitLoad().MustWaitIdle()

	// 等待页面加载完成 方式二
	//wait := page.MustWaitRequestIdle()
	//
	//wait()

	fmt.Println("RequestIdle")
	utils.Sleep(10)

	val1 := page.MustEval(`() => document.documentElement.scrollHeight`).Str()
	fmt.Println(val1)

	// 第二种方式向下滚动页面，直到滚动到底部
	//page.Mouse.MustScroll(0, 300)
	//utils.Sleep(10)
	//page.Mouse.MustScroll(0, 500)

	// 测试第二种方式滚动
	//page.MustEval(`() => window.scrollTo(0, document.documentElement.scrollHeight)`)
	//
	//utils.Sleep(10)
	//// 按下end键
	//page.KeyActions().Type(input.End).MustDo()
	//utils.Sleep(10)
	//val2 := page.MustEval(`() => document.documentElement.scrollHeight`).Str()
	//fmt.Println(val2)

	// 第三次滚动
	swipeUpToLoadMore(page, 2)

	// 获取网页内容
	pageHtml := page.MustEval(`() => document.documentElement.outerHTML`).Str()
	os.WriteFile("tmp1.txt", []byte(pageHtml), 0644)

	fmt.Println("success")

	time.Sleep(time.Hour)
	page.MustClose()
}

// 上滑页面加载更多
func swipeUpToLoadMore(page *rod.Page, times int64) {
	var (
		defHeight       = 0
		nowHeight       = 0
		count     int64 = 0
	)

	// 先获取当前页面的正文高度
	defHeight = page.MustEval(`() => document.documentElement.scrollHeight`).Int()
	fmt.Printf("页面当前高度 [%v]\n", defHeight)

	fmt.Println("开始自动上滑操作")
	// 循环上滑
	for {
		fmt.Println("上滑页面到最大高度")
		// 将页面滑动到正文最大高度
		page.MustEval(`() => window.scrollTo(0, document.documentElement.scrollHeight)`)
		fmt.Println("等待页面加载完成")
		// 等待页面加载
		page.MustWaitLoad().MustWaitIdle()
		utils.Sleep(10)
		fmt.Println("按下键盘End键")
		// 触发按键End，使页面滚动到最底部
		page.KeyActions().Type(input.End).MustDo()
		fmt.Println("等待页面加载完成")
		page.MustWaitLoad().MustWaitIdle()
		utils.Sleep(10)

		count += 1

		// 获取上滑后的页面正文高度
		nowHeight = page.MustEval(`() => document.documentElement.scrollHeight`).Int()
		fmt.Printf("页面上滑前高度 [%v] 上滑后高度 [%v]\n", defHeight, nowHeight)

		if nowHeight <= defHeight {
			// 已滚动到最底部
			fmt.Println("已滑动到页面最底部，停止自动上滑")
			break
		}

		// 是否有滑动次数限制
		if times > 0 && times <= count {
			fmt.Println("达到自动滑动次数限制，停止自动滑动")
			break
		}

		//	滑动后页面正文高度增加了，说明页面上滑了
		defHeight = nowHeight

		fmt.Println("等待一下再次上滑")
		utils.Sleep(15)
	}

	fmt.Println("自动上滑操作结束")
}

// 禁止加载图片、视频、字体等资源
func hijackVideoImage(page *rod.Page) {
	r := page.HijackRequests()
	r.MustAdd("*", func(ctx *rod.Hijack) {
		resType := ctx.Request.Type()
		if resType == proto.NetworkResourceTypeImage || resType == proto.NetworkResourceTypeMedia || resType == proto.NetworkResourceTypeFont {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
		} else {
			ctx.ContinueRequest(&proto.FetchContinueRequest{})
		}
	})
	go r.Run()
}
