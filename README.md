# goshorturl

    Golang 实现非官方短链接(Sina、Baidu)程序
#Usage
```
baidu:
    shorturl.Make("baidu", "https://996.icu/#/zh_CN")
sina:
    shorturl.Make("sina", "https://996.icu/#/zh_CN")
```

#Return
```
baidu:
    {"err_code":0,"short_url":"https://dwz.cn/BiIs5JVV","long_url":"https://996.icu/#/zh_CN"}
sina:
    {"err_code":0,"short_url":"http://t.cn/EJtKDOM","long_url":"https://996.icu/"}
```

#Api Doc

> * https://dwz.cn/console/apidoc
> * https://open.weibo.com/wiki/Short_url/shorten?retcode=6102

#Feature

> 应用Golang interface 来约束接入端统一实现、保持了返回值的统一性、易调用、易扩展
------