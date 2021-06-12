第三周作业
=======

#### 问题描述
> 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。


#### 运行
```bazaar
./week3
```
- 手动 `ctrl + c` 或者 `kill {pid}` 即可退出
> 注意： 若直接 ```go run server.go```, 在发送 `kill {pid}` 命令的时候，子进程（时间打印）将不会停止

- [参考](https://github.com/golang/go/issues/15553) 