/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     chromedp-rod-test
 * @Date:        2023-05-30 09:13
 * @Description:
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/leafney/rose"
	"time"
)

const (
	WSID = "be444244-1ca1-4c13-89ba-a9ded25b4eea"
)

// 通过ws链接
func main4() {
	l := launcher.MustNewManaged("ws://127.0.0.1:7317")
	// 配置
	l.Set("")
	l.Headless(false)

	browser := rod.New().Client(l.MustClient()).MustConnect()
	defer browser.MustClose()
}

func main3() {
	rodHandle(WSID)
}

func main() {
	//wsId := "443fefad-a3ee-458f-b82a-a67b5263388e"
	wsUrl := fmt.Sprintf("ws://127.0.0.1:9222/devtools/browser/%s", WSID)

	browser := rod.New().ControlURL(wsUrl).MustConnect().NoDefaultDevice().Trace(true).SlowMotion(time.Second)
	//defer browser.MustClose()

	// 新建一个tab页
	page := browser.MustPage()

	//defer func() {
	//	fmt.Println("即将关闭Tab页")
	//	time.Sleep(3 * time.Second)
	//	page.MustClose()
	//}()

	// 禁止弹窗
	page.MustEvalOnNewDocument(`window.alert = () => {};window.prompt = () => {}`)

	fmt.Println("准备开始")
	//utils.Sleep(20)

	// 阻止请求图片、视频、字体文件等类型
	hijackVideoImage(page)

	//
	//page.
	//	MustNavigate("https://www.youtube.com/")

	//MustNavigate("https://www.iplaysoft.com/").
	//MustWaitLoad()
	//MustClose()

	// 等待页面加载完成 方式一
	//page.MustWaitLoad().MustWaitIdle()

	// 等待页面加载完成 方式二
	//wait := page.MustWaitRequestIdle()
	//
	//wait()

	fmt.Println("RequestIdle")
	utils.Sleep(5)

	//val1 := page.MustEval(`() => document.documentElement.scrollHeight`).Str()
	//fmt.Println(val1)

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
	//swipeUpToLoadMore(page, 0)

	// 获取网页内容
	//saveHtml(page)

	// 测试点击，输入
	AutoRunAction(page)

	swipePageLoadMore(page, 2)

	fmt.Println("success")

	time.Sleep(time.Hour)
	page.MustClose()
}

// ----------------------

// 独立方法，测试 defer 使用
func rodHandle(wsId string) {
	wsUrl := fmt.Sprintf("ws://127.0.0.1:9222/devtools/browser/%s", wsId)
	b := rod.New().ControlURL(wsUrl).MustConnect().Trace(true)
	//defer b.MustClose()

	page := b.MustPage()
	defer page.MustClose()

	page.MustNavigate("https://www.youtube.com/")
	page.MustWaitLoad()
	time.Sleep(time.Second * 20)
}

// ----------------------

type ActionModel struct {
	Action  string `json:"action"`
	Method  string `json:"method"`
	Element string `json:"element"`
	Content string `json:"content"`
}

func AutoRunAction(page *rod.Page) {
	//actionStr := `[{"action":"navigate","method":"","element":"","content":"https://www.baidu.com/"},{"action":"visible","method":"css","element":"#kw","content":""},{"action":"input","method":"css","element":"#kw","content":"golang go-rod"},{"action":"click","method":"css","element":"#su","content":""},{"action":"sleep","method":"","element":"","content":"10"}]`
	actionStr := `[{"action":"navigate","method":"","element":"","content":"https://www.youtube.com/"},{"action":"wait","method":"","element":"","content":""},{"action":"input","method":"css","element":"div#search-input #search","content":"golang go-rod"},{"action":"click","method":"css","element":"#search #search-icon-legac","content":""},{"action":"sleep","method":"","element":"","content":"10"}]`

	list := make([]ActionModel, 0)
	rose.JsonUnMarshalStr(actionStr, &list)

	err := rod.Try(func() {

		for _, m := range list {
			fmt.Println("run ", m.Action)

			switch m.Action {
			case "navigate":
				page.MustNavigate(m.Content)
			case "visible":
				page.Timeout(10 * time.Second).MustElement(m.Element).MustWaitVisible().CancelTimeout()
			case "input":
				page.Timeout(10 * time.Second).MustElement(m.Element).MustInput(m.Content).CancelTimeout()
			case "click":
				page.Timeout(10 * time.Second).MustElement(m.Element).MustClick().CancelTimeout()
			case "sleep":
				rose.TSleep(rose.StrToInt64(m.Content))
			case "wait":
				page.MustWaitIdle().MustWaitIdle()
			case "scroll":
				swipePageLoadMore(page, rose.StrToInt(m.Content))
			default:

			}
			// 执行完成一次操作，等待页面加载
			waitPageLoad(page)
			pageSleepRand()
		}
	})

	if errors.Is(err, context.DeadlineExceeded) {
		//	超时
		fmt.Printf("页面超时 [%v]", err)
	} else if errors.Is(err, context.Canceled) {
		//	cancel
		fmt.Printf("页面取消 [%v]", err)
	} else {
		//	other
		fmt.Printf("其他错误 [%v]", err)
	}
}

// ----------------------

// 保存网页
func pageGetHtml(page *rod.Page) string {
	return page.MustEval(`() => document.documentElement.outerHTML`).Str()
	//os.WriteFile("tmp1.txt", []byte(pageHtml), 0644)
}

// 上滑页面加载更多
func swipePageLoadMore(page *rod.Page, mode int) {
	var (
		defHeight = 0
		nowHeight = 0
		count     = 0
	)
	fmt.Println("开始自动上滑操作")

	// 先获取当前页面的正文高度
	defHeight = getPageHeight(page)
	fmt.Printf("页面当前高度 [%v]\n", defHeight)

	if mode > 100 {
		//	滑动指定距离
		slidePageDownDistance(page, mode)
	} else if mode == 0 {
		//	不滑动

	} else {
		// 循环上滑
		for {
			fmt.Println("将页面滑动到正文最大高度")
			slidePageDownEnd(page)

			fmt.Println("等待页面加载完成")
			// 等待页面加载
			page.MustWaitLoad().MustWaitIdle()

			pageSleepRand()

			fmt.Println("按下键盘End键")
			// 触发按键End，使页面滚动到最底部
			pressKeyEnd(page)

			//page.KeyActions().Type(input.End).MustDo()
			fmt.Println("等待页面加载完成")
			page.MustWaitLoad().MustWaitIdle()

			pageSleepRand()

			count += 1

			// 获取上滑后的页面正文高度
			nowHeight = getPageHeight(page)
			fmt.Printf("页面上滑前高度 [%v] 上滑后高度 [%v]\n", defHeight, nowHeight)

			if nowHeight <= defHeight {
				// 已滚动到最底部
				fmt.Println("已滑动到页面最底部，停止自动上滑")
				break
			}

			// 是否有滑动次数限制
			if mode > 0 && mode <= count {
				fmt.Println("达到自动滑动次数限制，停止自动滑动")
				break
			}

			//	滑动后重置页面最新正文高度
			defHeight = nowHeight

			fmt.Println("等待下一次自动滑动")
			pageSleepRand()
		}
	}

	fmt.Println("自动上滑操作结束")
}

func pageSleepRand() {
	rose.TSleepRand(5, 30)
}

func pageSleep(sec int64) {
	rose.TSleep(sec)
}

// 获取页面高度
func getPageHeight(page *rod.Page) int {
	return page.MustEval(`() => document.documentElement.scrollHeight`).Int()
}

// 页面向下滑动指定距离
func slidePageDownDistance(page *rod.Page, distance int) {
	jsStr := fmt.Sprintf(`() => window.scrollTo(0, %d)`, distance)
	page.MustEval(jsStr)
}

// 滑动到页面底部
func slidePageDownEnd(page *rod.Page) {
	page.MustEval(`() => window.scrollTo(0, document.documentElement.scrollHeight)`)
}

func slidePageDownEndSmooth(page *rod.Page) {
	page.MustEval(`() => window.scrollTo({top: document.documentElement.scrollHeight,left: 0,behavior: "smooth"})`)
}

// 按下键盘上的End键
func pressKeyEnd(page *rod.Page) {
	page.KeyActions().Type(input.End).MustDo()
}

func waitPageLoad(page *rod.Page) {
	page.MustWaitLoad().MustWaitIdle()
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
