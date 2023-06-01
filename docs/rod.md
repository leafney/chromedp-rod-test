## rod

- [go-rod/rod: A Devtools driver for web automation and scraping](https://github.com/go-rod/rod)
- [rod/lib/examples at main · go-rod/rod · GitHub](https://github.com/go-rod/rod/tree/main/lib/examples)

## 依赖

```
go get github.com/go-rod/rod
```

## 运行

### Macos本地启动chrome浏览器

```
默认启动命令：
sudo /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --remote-debugging-port=9222

其他启动参数：
--remote-debugging-port=9222 --no-default-browser-check --no-first-run --disable-infobars

```

然后可以访问 `http://127.0.0.1:9222/json/version` 找到启动的浏览器的ws链接地址：`webSocketDebuggerUrl` 的值。
访问 `http://127.0.0.1:9222/json/list` 可以看到当前打开的page页面的列表。

`ws://127.0.0.1:9222/devtools/browser/5293bae5-6cea-454b-9558-5889a1abb9ac`

----

## 测试

### 1-rod

经测试，`go-rod` 无法截取请求过程中的其他异步请求。通过分析 `ctx.MustLoadResponse()` 方法的源码后可知：其内部是通过 `http.Client` 对当前拦截到的url发起的另一次请求，跟当前拦截的请求并没有关系。
这样的话，当前被拦截的请求的响应数据是无法获取到的，而得到的请求响应数据是再次发起请求后得到的，结果可能并不是想要的。


### 步骤化动作

----

| 操作类型 | 选择方式 | 选择元素 | 动作类型 | 动作内容 | 说明 |
| -- |
| navigate |   |   |   |   | www.baidu.com | 打开网址 |
| find | element | "body > footer" | visible |     | 等待元素显示 |
| find | element | "#pkg-examples" | cick |     | 找到某元素并点击 |
| find | element | ".text"         | input | "hello" | 找到某元素然后输入内容 |
| wait |      |      |   wait |   | 等待页面加载 |
| sleep |     |    |


| 动作类型 | 选择方式 | 选择元素 | 动作内容 | 说明 |
| navigate |     |     | www.baidu.com | 打开网址 |
| visible | element | ".name" |    | 等待元素显示 |
| invisible | element | "" |  | 等待元素隐藏 |
| input | byId | "#title" |  "hello" | 找到某元素输入内容  |
| click |    |     |    |  找到某元素并点击 |
| wait |    |    |    | 等待页面加载 |
| sleep |    |    | 10 | 停止10s |
| eval |   |   | js script | 执行js代码 |
| console |    |   |  js script | 在 DevTools 中执行js script |

----


### 特殊方法

- `MustWaitLoad` 表示会一直等待直到页面加载完毕
- 


