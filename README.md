
## Gorilla Context
 
##### 转载

> Go语言经典库使用分析，未完待续，欢迎扫码关注公众号flysnow_org或者网站http://www.flysnow.org/，第一时间看后续系列。觉得有帮助的话，顺手分享到朋友圈吧，感谢支持。

> 在Go1.7之前，Go标准库还没有内置Context的时候，如果我们想在一个Http.Request里附加值，怎么做呢？一般都是Map对象，存储对应的Request以及附加的值，然后在需要的时候取出来，今天我们介绍的这个就是实现了一个类似于这样功能的库，因为比较简单，而且实用，所以就先选择它来分析。

##### 原文链接
[Go语言经典库使用分析（二）| Gorilla Context](!http://www.flysnow.org/2017/07/29/go-classic-libs-gorilla-context.html)