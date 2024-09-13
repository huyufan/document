# 状态信息 - info

## Redis提供的INFO命令不仅能够查看实时的吞吐量(ops/sec)，还能看到一些有用的运行时信息。
``` shell
127.0.0.1:6379> info
# Server
redis_version:3.2.3 #redis版本号
redis_git_sha1:00000000 #git sha1摘要值
redis_git_dirty:0  #git dirty标识
redis_build_id:443e50c39cbcdbe0 #redis构建id
redis_mode:standalone  #运行模式：standalone、sentinel、cluster
os:Linux 3.10.0-514.16.1.el7.x86_64 x86_64 #服务器宿主机操作系统
arch_bits:64 服务器宿主机CUP架构（32位/64位）
multiplexing_api:epoll #redis IO机制
gcc_version:4.8.5  #编译 redis 时所使用的 GCC 版本
process_id:1508  #服务器进程的 PID
run_id:b4ac0f9086659ce54d87e41d4d2f947e19c28401 #redis 服务器的随机标识符 （用于 Sentinel 和集群）
tcp_port:6380  #redis服务监听端口
uptime_in_seconds:520162 #redis服务启动以来经过的秒数
uptime_in_days:6 #redis服务启动以来经过的天数
hz:10  #redis内部调度（进行关闭timeout的客户端，删除过期key等等）频率，程序规定serverCron每秒运行10次
lru_clock:16109450 #自增的时钟，用于LRU管理,该时钟100ms(hz=10,因此每1000ms/10=100ms执行一次定时任务)更新一次
executable:/usr/local/bin/redis-server
config_file:/data/redis-6380/redis.conf 配置文件的路径

# Clients
connected_clients:2   #已连接客户端的数量（不包括通过从属服务器连接的客户端）
client_longest_output_list:0 #当前连接的客户端当中，最长的输出列表
client_biggest_input_buf:0 #当前连接的客户端当中，最大输入缓存
blocked_clients:0 #正在等待阻塞命令（BLPOP、BRPOP、BRPOPLPUSH）的客户端的数量

# Memory
used_memory:426679232 #由 redis 分配器分配的内存总量，以字节（byte）为单位
used_memory_human:406.91M   #以可读的格式返回 redis 分配的内存总量（实际是used_memory的格式化）
used_memory_rss:443179008 #从操作系统的角度，返回 redis 已分配的内存总量（俗称常驻集大小）。这个值和 top 、 ps等命令的输出一致
used_memory_rss_human:422.65M # redis 的内存消耗峰值(以字节为单位) 
used_memory_peak:426708912
used_memory_peak_human:406.94M
total_system_memory:16658403328
total_system_memory_human:15.51G
used_memory_lua:37888   # Lua脚本存储占用的内存
used_memory_lua_human:37.00K
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
mem_fragmentation_ratio:1.04 # used_memory_rss/ used_memory
mem_allocator:jemalloc-4.0.3

# Persistence
loading:0 #服务器是否正在载入持久化文件，0表示没有，1表示正在加载
rdb_changes_since_last_save:3164272  #离最近一次成功生成rdb文件，写入命令的个数，即有多少个写入命令没有持久化
rdb_bgsave_in_progress:0 #服务器是否正在创建rdb文件，0表示否
rdb_last_save_time:1559093160  #离最近一次成功创建rdb文件的时间戳。当前时间戳 - rdb_last_save_time=多少秒未成功生成rdb文件
rdb_last_bgsave_status:ok  #最近一次rdb持久化是否成功
rdb_last_bgsave_time_sec:-1  #最近一次成功生成rdb文件耗时秒数
rdb_current_bgsave_time_sec:-1 #如果服务器正在创建rdb文件，那么这个域记录的就是当前的创建操作已经耗费的秒数
aof_enabled:0 #是否开启了aof
aof_rewrite_in_progress:0 #标识aof的rewrite操作是否在进行中
aof_rewrite_scheduled:0  #rewrite任务计划，当客户端发送bgrewriteaof指令，如果当前rewrite子进程正在执行，那么将客户端请求的bgrewriteaof变为计划任务，待aof子进程结束后执行rewrite
aof_last_rewrite_time_sec:-1 #最近一次aof rewrite耗费的时长
aof_current_rewrite_time_sec:-1 #如果rewrite操作正在进行，则记录所使用的时间，单位秒
aof_last_bgrewrite_status:ok #上次bgrewriteaof操作的状态
aof_last_write_status:ok #上次aof写入状态

# Stats
total_connections_received:10   #服务器已经接受的连接请求数量
total_commands_processed:9510792   #redis处理的命令数
instantaneous_ops_per_sec:1   #redis当前的qps，redis内部较实时的每秒执行的命令数
total_net_input_bytes:1104411373   #redis网络入口流量字节数
total_net_output_bytes:66358938 #redis网络出口流量字节数
instantaneous_input_kbps:0.04  #redis网络入口kps
instantaneous_output_kbps:3633.35  #redis网络出口kps
rejected_connections:0  #拒绝的连接个数，redis连接个数达到maxclients限制，拒绝新连接的个数
sync_full:0  #主从完全同步成功次数
sync_partial_ok:0  #主从部分同步成功次数
sync_partial_err:0  #主从部分同步失败次数
expired_keys:0   #运行以来过期的key的数量
evicted_keys:0  #运行以来剔除(超过了maxmemory后)的key的数量
keyspace_hits:87  #命中次数
keyspace_misses:17   #没命中次数
pubsub_channels:0  #当前使用中的频道数量
pubsub_patterns:0  #当前使用的模式的数量
latest_fork_usec:0   #最近一次fork操作阻塞redis进程的耗时数，单位微秒
migrate_cached_sockets:0   #是否已经缓存了到该地址的连接

# Replication
role:master  #实例的角色，是master or slave
connected_slaves:0  #连接的slave实例个数
master_repl_offset:0 #主从同步偏移量,此值如果和上面的offset相同说明主从一致没延迟，与master_replid可被用来标识主实例复制流中的位置
repl_backlog_active:0   #复制积压缓冲区是否开启
repl_backlog_size:1048576  #复制积压缓冲大小
repl_backlog_first_byte_offset:0  #复制缓冲区里偏移量的大小
repl_backlog_histlen:0   #此值等于 master_repl_offset - repl_backlog_first_byte_offset,该值不会超过repl_backlog_size的大小

# CPU
used_cpu_sys:507.00  #将所有redis主进程在核心态所占用的CPU时求和累计起来
used_cpu_user:280.48   #将所有redis主进程在用户态所占用的CPU时求和累计起来
used_cpu_sys_children:0.00  #将后台进程在核心态所占用的CPU时求和累计起来
used_cpu_user_children:0.00  #将后台进程在用户态所占用的CPU时求和累计起来

# Cluster
cluster_enabled:0

# Keyspace
db0:keys=5557407,expires=362,avg_ttl=604780497
db15:keys=1,expires=0,avg_ttl=0

```

# 监控执行命令 - monitor

##  monitor用来监视服务端收到的命令。

# 监控延迟

## 监控延迟 - latency

``` shell
 redis-cli --latency -h 127.0.0.1
 ```
 ## 服务端内部的延迟监控稍微麻烦一些，因为延迟记录的默认阈值是0。尽管空间和时间耗费很小，Redis为了高性能还是默认关闭了它。所以首先我们要开启它，设置一个合理的阈值，例如下面命令中设置的100ms：

 ``` shell
 CONFIG SET latency-monitor-threshold 100
 ```

 # redis 优化

 ## 业务服务器到 Redis 服务器之间的网络存在问题，例如网络线路质量不佳，网络数据包在传输时存在延迟、丢包等情况
 ## Redis 本身存在问题，需要进一步排查是什么原因导致 Redis 变慢

 ## 什么是基准性能？
 ## 为了避免业务服务器到 Redis 服务器之间的网络延迟，你需要直接在 Redis 服务器上测试实例的响应延迟情况。执行以下命令，就可以测试出这个实例 60 秒内的最大响应延迟

 ``` shell
redis-cli -h 127.0.0.1 -p 6379 --intrinsic-latency 60
 ```
 ## 从输出结果可以看到，这 60 秒内的最大响应延迟为 72 微秒（0.072毫秒）。

 ## 你还可以使用以下命令，查看一段时间内 Redis 的最小、最大、平均访问延迟：
 ``` shell
redis-cli -h 127.0.0.1 -p 6379 --latency-history -i 1
 ```

 ### 第一步，你需要去查看一下 Redis 的慢日志（slowlog）。

 ``` shell
# 命令执行耗时超过 5 毫秒，记录慢日志
CONFIG SET slowlog-log-slower-than 5000
# 只保留最近 500 条慢日志
CONFIG SET slowlog-max-len 500

 ```

 ### 可以执行以下命令，就可以查询到最近记录的慢日志：
 ``` shell
SLOWLOG get 5
1) 1) (integer) 32693       # 慢日志ID
   2) (integer) 1593763337  # 执行时间戳
   3) (integer) 5299        # 执行耗时(微秒)
   4) 1) "LRANGE"           # 具体执行的命令和参数
      2) "user_list:2000"
      3) "0"
      4) "-1"
2) 1) (integer) 32692
   2) (integer) 1593763337
   3) (integer) 5044
   4) 1) "GET"
      2) "user_info:1000"
 ```

 #### 通过查看慢日志，我们就可以知道在什么时间点，执行了哪些命令比较耗时。
 #### 如果你的应用程序执行的 Redis 命令有以下特点，那么有可能会导致操作延迟变大：
 - 经常使用 O(N) 以上复杂度的命令，例如 SORT、SUNION、ZUNIONSTORE 聚合类命令
 - 使用 O(N) 复杂度的命令，但 N 的值非常大
 - 第一种情况导致变慢的原因在于，Redis 在操作内存数据时，时间复杂度过高，要花费更多的 CPU 资源。
 - 第二种情况导致变慢的原因在于，Redis 一次需要返回给客户端的数据过多，更多时间花费在数据协议的组装和网络传输过程中。
 - 另外，我们还可以从资源使用率层面来分析，如果你的应用程序操作 Redis 的 OPS 不是很大，但 Redis 实例的 CPU 使用率却很高，那么很有可能是使用了复杂度过高的命令导致的。除此之外，我们都知道，Redis 是单线程处理客户端请求的，如果你经常使用以上命令，那么当 Redis 处理客户端请求时，一旦前面某个命令发生耗时，就会导致后面的请求发生排队，对于客户端来说，响应延迟也会变长。

 ####  操作bigkey
 - 如果你查询慢日志发现，并不是复杂度过高的命令导致的，而都是 SET / DEL 这种简单命令出现在慢日志中，那么你就要怀疑你的实例否写入了 bigkey。

 ``` shell 
 redis-cli -h 127.0.0.1 -p 6379 --bigkeys -i 0.01 -a 密码
 ```

 - 对线上实例进行 bigkey 扫描时，Redis 的 OPS 会突增，为了降低扫描过程中对 Redis 的影响，最好控制一下扫描的频率，指定 -i 参数即可，它表示扫描过程中每次扫描后休息的时间间隔，单位是秒
 - 扫描结果中，对于容器类型（List、Hash、Set、ZSet）的 key，只能扫描出元素最多的 key。但一个 key 的元素多，不一定表示占用内存也多，你还需要根据业务情况，进一步评估内存占用情况

##### 那针对 bigkey 导致延迟的问题，有什么好的解决方案呢？这里有两点可以优化：
- 业务应用尽量避免写入 bigkey
- 如果你使用的 Redis 是 4.0 以上版本，用 UNLINK 命令替代 DEL，此命令可以把释放 key 内存的操作，放到后台线程中去执行，从而降低对 Redis 的影响
- 如果你使用的 Redis 是 6.0 以上版本，可以开启 lazy-free 机制（lazyfree-lazy-user-del = yes），在执行 DEL 命令时，释放内存也会放到后台线程中执行
###### 但即便可以使用方案 2，我也不建议你在实例中存入 bigkey。这是因为 bigkey 在很多场景下，依旧会产生性能问题。例如，bigkey 在分片集群模式下，对于数据的迁移也会有性能影响，以及我后面即将讲到的数据过期、数据淘汰、透明大页，都会受到 bigkey 的影响
