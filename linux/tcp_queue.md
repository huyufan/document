# TCP 半连接队列和全连接队列

## 如何知道应用程序的 TCP 全连接队列大小？
- 在服务端可以使用 **ss**  命令，来查看 TCP 全连接队列的情况：  在「LISTEN 状态」和「非 LISTEN 状态」所表达的含义是不同的。

### 示例  ss -lnt | grep 80
``` shell
# -l 显示正在监听(listening)的socket
# -n 不解析服务名称
# -t 只显示tcp socket

```

- Recv-Q：当前全连接队列的大小，也就是当前已完成三次握手并等待服务端  **accept()**  的 TCP 连接
- Send-Q：当前全连接最大队列长度，上面的输出结果说明监听 8088 端口的 TCP 服务，最大全连接长度为 128；


### 查看TCP 全连接队列溢出情况

``` shell

netstat -s | grep overflowed

### 返回数值 41150 times the listen queue of a socket overflowed
#上面看到的 41150 times ，表示全连接队列溢出的次数，注意这个是累计值。可以隔几秒钟执行下，如果这个数字一直在增加的话肯定全连接队列偶尔满了。
```

#### Linux 有个参数可以指定当 TCP 全连接队列满了会使用什么策略来回应客户端。

``` SHELL
  cat /proc/sys/net/ipv4/tcp_abort_on_overflow
```
- 0 ：如果全连接队列满了，那么 server 扔掉 client  发过来的 ack ；
- 1 ：如果全连接队列满了，server 发送一个 **reset**  包给 client，表示废掉这个握手过程和这个连接；

#### 修改全队列
``` shell

echo 5000 > /proc/sys/net/core/somaxcom
```
#### nginx 也可以设置
``` shell
# /usr/local/nginx/nginx.conf

server {
   listen 8088 default backlog = 5000
}
```

## TCP 半连接队列溢出


``` shell
### 查看当前半连接队列
 netstat -natp | grep SYN_RECV | wc -l 

### 查看半连接溢出情况
netstat -s | grep "SYNs to LISTEN"
```

### 如何防御 SYN 攻击？
- 增大半连接队列
- 开启 tcp_syncookies 功能
- 减少 SYN + ACK 重传次数

#### 增大半连接队列
- 增大 tcp_max_syn_backlog 和 somaxconn 的方法是修改 Linux 内核参数：

``` shell
#输入重定向  
#将 echo 命令的输出重定向到 /root/e.log 文件中。 
#> 表示输出重定向。它将 echo 5 产生的输出 5 写入文件 /root/e.log 中。
#如果文件 /root/e.log 已经存在，它会被覆盖。
#如果文件不存在，则会创建该文件。
echo 5000 > /proc/sys/net/ipv4/tcp_max_syn_backlog
echo 5000 > /proc/sys/net/core/somaxconn
server {
   listen 8088 default backlog = 5000
}
```

- 开启 tcp_syncookies 功能

``` shell
#切当SYN 半连接队列放不下时，再启用他
echo 1 > /proc/sys/net/ipv4/tcp_syncookies
```

- 减少 SYN + ACK 重传次数

``` shell
echo 1 > /proc/sys/net/ipv4/tcp_synack_retries
```
# 最后，改变了如上这些参数后，要重启 Nginx 服务，因为 SYN 半连接队列和 accept 队列都是在  **listen()** 初始化的。
