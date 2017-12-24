# restsvr
基于Golang的Unix(Linux)任务调度系统

## 基本实现
* HTTP Server  
使用Golang自带net/http包创建，端口8888，后续考虑优化路由功能引用gorilla/mux
* * *
* 任务分发  
给予Redis的LIST实现任务队列机制(LPUSH/RPOP)，由redigo包实现
* * *
* 前端框架  
Jquery + UIKit
* * *
