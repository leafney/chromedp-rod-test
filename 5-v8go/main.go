/**
 * @Author:      leafney
 * @GitHub:      https://github.com/leafney
 * @Project:     chromedp-rod-test
 * @Date:        2023-05-31 17:24
 * @Description:
 */

package main

import (
	"fmt"
	v8 "rogchap.com/v8go"
	"time"
)

func main() {
	ctx := v8.NewContext()

	vals := make(chan *v8.Value, 1)
	errs := make(chan error, 1)

	//ctx.RunScript(TPL_CHEERIO_JS(), "base.js")

	go func() {
		val, err := ctx.RunScript(TPL_CHEERIO_JS(), "base.js")
		if err != nil {
			errs <- err
			return
		}
		vals <- val
	}()

	select {
	case val := <-vals:
		// success
		fmt.Printf("sucdess %v", val)
	case err := <-errs:
		// javascript error
		fmt.Printf("err %v", err)
	case <-time.After(200 * time.Millisecond):
		vm := ctx.Isolate()     // get the Isolate from the context
		vm.TerminateExecution() // terminate the execution
		err := <-errs           // will get a termination error back from the running script
		fmt.Printf("after err %v", err)
	}

	time.Sleep(time.Hour)

	//
	//	data := `
	//const html="<div><a href='/post/7231834021055250490' target='_blank' class='title'>Node.js版本管理工具，我选择n</a><p>哈哈哈</p></div>";
	//
	//const $ = cheerio.load(html);
	//
	//const text= $(".title").text();
	//console.log(text);
	//`

	//data := `<div><a href="/post/7231834021055250490" target="_blank" class="title">Node.js版本管理工具，我选择n</a><p>哈哈哈</p></div>`
	//
	//val, _ := ctx.RunScript(`Jiexi("`+data+`")`, "export.js")
	//fmt.Printf("result : %v", val)
}
