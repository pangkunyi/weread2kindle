微信读书转Kindle
=========
说明

通过微信的电脑客户端的微信读书公众号获取登录cookie，然后模拟登录微信读书并获取对应书籍的内容并输出到html文件，最后通过kindlegen或用kindle previewer来转成mobi格式的书籍。

安装
```
git clone https://github.com/pangkunyi/weread2kindle.git
go build
```

执行
```
> weread2kindle
Usage of weread2kindle:
  -b int
        book id, should be great than 0
  -c string
        weread login cookie
  -d string
        output directory (default ".")
```