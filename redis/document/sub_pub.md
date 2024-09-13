# 发布和订阅 是一种消息通信模式：发送者(pub)发送消息，订阅者(sub)接收消息

## Redis有两种发布/订阅模式：
- 基于频道(Channel)的发布/订阅
- 基于模式(pattern)的发布/订阅
  通配符中?表示1个占位符，*表示任意个占位符(包括0)，?*表示1个以上占位符。

``` shell
subscribe channel(频道)

psubscribe c? c* c?*

publish channel(频道) vlaue


```
### 使用psubscribe命令可以重复订阅同一个频道，如客户端执行了psubscribe c? c?*。这时向c1发布消息客户端会接受到两条消息，而同时publish命令的返回值是2而不是1。同样的，如果有另一个客户端执行了subscribe c1 和psubscribe c?*的话，向c1发送一条消息该客户顿也会受到两条消息(但是是两种类型:message和pmessage)，同时publish命令也返回2.


### 为什么单线程的 Redis 能那么快？
- Redis的瓶颈主要在IO而不是CPU，所以为了省开发量，在6.0版本前是单线程模型；其次，Redis 是单线程主要是指 Redis 的网络 IO 和键值对读写是由一个线程来完成的，这也是 Redis 对外提供键值存储服务的主要流程。（但 Redis 的其他功能，比如持久化、异步删除、集群数据同步等，其实是由额外的线程执行的）。





