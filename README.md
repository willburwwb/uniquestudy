# uniquestudy
## day 1
早上先安装archlinux，担心实装会出问题所以用的虚拟机，然后下午阅读《linux就该这么学》，了解到一些以前没看过的内容。

## day 2
从下午开始学习，接着看linux相关知识，大概学习linux的相关任务，同时了解了关于文本处理的相关命令，并开始阅读Learn-vim，然后学习shell的相关知识。

## day 3-4
昨天白天有点事，晚上开始学习，然后忘记提交了（
这两天主要复习了git，在archlinux下用bash-it配置了一下bash，然后学习了shell编程，和构建工具，晚上下载了docker，开始了解docker。

## day 5
今天主要学习了docker的用法，关于docker的镜像容器的相关命令，和dockerfile，然后看了下Go语言入门书籍。

## day 6
今天上午起床有点晚(),看了看cmake的内容，然后下午的话接着看go，晚上打了场acm的比赛有点累，接着看go
update：
发现自己linux的几个文本命令还没看，先补下

## Week two
## day 2-3
（day 1有事所以摸了...）主要学习了go语言，然后安装gin框架和postgrepsql数据库，简单学习了下，跟着gin的中文教程简单操作了下gin的路由操作
## day 4
大致写了gin的注册功能，用到go-email的发送功能，用viper对配置进行管理，然后验证码在redis中储存
## day 5
写了登录的功能，通过jwt中间件进行登录验证，同时增加了邮箱找回密码功能
## day 6
写了post的模块功能，一开始没有想好数据库的设计，后来增加了储存匿名名称的表，然后一共有post，Comment，nickname三个表
在删除表时通过中间件检验是否有权限（感觉可能要用casbin这样的权限管理框架，但没有想好怎么用...）
## day 7
写了comment的模块功能