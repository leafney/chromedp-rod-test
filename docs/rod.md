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

