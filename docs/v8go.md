# v8go

- [rogchap/v8go: Execute JavaScript from Go](https://github.com/rogchap/v8go) 


```
import v8 "rogchap.com/v8go"
```

## 测试一

尝试通过 `v8go` 加载三方的js库，然后调用三方库来执行js方法。测试后报错 `require is not defined`

参考了 - [rss-can/moment.go at main · soulteary/rss-can · GitHub](https://github.com/soulteary/rss-can/blob/main/internal/jssdk/moment.go)  的实现，
但结果却未达到预期。

后来看到这个讨论 - [Node modules · Issue #148 · rogchap/v8go](https://github.com/rogchap/v8go/issues/148) 

了解到 `v8go` 并不支持调用 `nodejs` 相关的api，比如 `require` 。所以得采用其他的方法来执行 `nodejs` 相关的方法。

----

