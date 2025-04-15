# tcp
## 先看tcp报文头格式
![TCP 包头格式](https://raw.githubusercontent.com/huyufan/document/refs/heads/master/network/image/8.webp)

### 源端口号 和 目标端口 是不可少的，如果没有这两个端口号，数据就不知道应该发给哪个应用。

### 序号 解决包乱序的问题

### 确认号 目的是确认发出去对方是否有收到。如果没有收到就应该重新发送，直到送达，这个是为了解决丢包的问题。

### 状态码 SYN 发送一个连接  ACK是回复  RST 是重新连接 FIN 是结束连接

### 窗口大小 做流量控制 ，通信双方各声明一个窗口（缓存大小），标识自己当前能够的处理能力，别发送的太快，撑死我，也别发的太慢，饿死我。

### 阻塞控制，对于真正的通路堵车不堵车，它无能为力，唯一能做的就是控制自己，也即控制发送的速度。不能改变世界，就改变自己嘛。


## tcp  三次握手
![TCP 三次握手](https://raw.githubusercontent.com/huyufan/document/refs/heads/master/network/image/TCP%E4%B8%89%E6%AC%A1%E6%8F%A1%E6%89%8B.drawio.webp)

### 一开始 客户端 和服务端都处于CLOSED 状态。先是服务器主动监听某个端口 处于LISTEN状态
### 然后客户端主动发起连接SYN


