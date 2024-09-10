# 5种基本类型 String(字符串)  List(列表) Hash(哈希) Set(集合) Zset(有序集合)

## String 
-  可以是字符串,整数,浮点型   整个字符串或字符串的一部分进行操作,对整数或浮点数有自增和自减
- String类型的值是二进制安全的，意思是 redis 的 string 可以包含任何数据。如数字，字符串，jpg图片或者序列化的对象。
- Redis String 适用于存储单一数据，操作简单、效率高。
| 命令 | 简述 | 使用 |
|-------|-------|-------|
| GET | 获取存储在给定键中的值 | GET KEY |
| SET | 设置存储在给定键中的值 | SET KEY VALUE|
| DEL | 删除存储在给定键中的值 | DEL KEY|
| INCR | 将键存储的值加1 | INCR KEY|
| DECR | 将键存储的值减1 | DECR KEY|
| INCRBY | 将键存储的值加上整数 | INCRBY KEY amount|
| DECRBY | 将键存储的值减去整数 | DECRBY KEY amount|

## List
- 一个链表，链表上的每个节点都包含一个字符串  对链表的两端进行push和pop操作，读取单个或多个元素；根据值查找或删除元素
- Redis中的List其实就是链表（Redis用双端链表实现List）。
- Lpush + Lpop = Stack 栈
- Lpush + Rpop = Queue  队列
- Lpush + ltrim =  Capped Collection（有限集合）
- Lpush + brpop = Message Queue（消息队列）

| 命令 | 简述 | 使用 |
|-------|-------|-------|
| LPUSH | 将给定值推入到列表左端 | LPUSH KEY VALUE |
| RPUSH | 将给定值推入到列表右端 | RPUSH KEY VALUE |
| RPOP | 从列表的右端弹出一个值，并返回被弹出的值 | RPOP KEY|
| LPOP | 从列表的左端弹出一个值，并返回被弹出的值 | LPOP KEY|
| LRANGE | 获取列表在给定范围上的所有值 | LRANGE KEY 0 -1|
| LINDEX | 通过索引获取列表中的元素。你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。 | LINDEX KEY INDEX|

## Hash
-  包含键值对的无序散列表  包含方法有添加、获取、删除单个元素
- 缓存： 能直观，相比string更节省空间，的维护缓存信息，如用户信息，视频信息等。
- Redis Hash 适用于存储多个字段的数据，内存使用更高效，支持批量操作和结构化数据管理。
| 命令 | 简述 | 使用 |
|-------|-------|-------|
| HSET | 添加键值对 | HSET KEY KEY1 VALUE |
| HGET | 获取指定散列键的值 | HGET KEY KEY1|
| HGETALL | 获取散列中包含的所有键值对 | HGETALL KEY|
| HDEL | 如果给定键存在于散列中，那么就移除这个键 | HDEL KEY KEY1|


## Set
- Redis 的 Set 是 String 类型的无序集合。集合成员是唯一的，这就意味着集合中不能出现重复的数据。
- Redis 中集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。
- 包含字符串的无序集合 字符串的集合，包含基础的方法有看是否存在添加、获取、删除；还包含计算交集、并集、差集等
- 标签（tag）,给用户添加标签，或者用户给消息添加标签，这样有同一标签或者类似标签的可以给推荐关注的事或者关注的人。
- 点赞，或点踩，收藏等，可以放到set中实现
| 命令 | 简述 | 使用 |
|-------|-------|-------|
| SADD | 向集合添加一个或多个成员 | SADD KEY VALUE1 VALUE2 |
| SCARD | 获取集合的成员数 | SCARD KEY|
| SMEMBERS | 返回集合中的所有成员    | SMEMBERS KEY|
| SISMEMBER | 判断 member 元素是否是集合 key 的成员 | SISMEMBER KEY VALUE1|

## Zset
- 和散列一样，用于存储键值对 字符串成员与浮点数分数之间的有序映射；元素的排列顺序由分数的大小决定；包含方法有添加、获取、删除单个元素以及根据分值范围或成员来获取元素
- 有序集合的成员是唯一的, 但分数(score)却可以重复。有序集合是通过两种数据结构实现：

- 压缩列表(ziplist): ziplist是为了提高存储效率而设计的一种特殊编码的双向链表。它可以存储字符串或者整数，存储整数时是采用整数的二进制而不是字符串形式存储。它能在O(1)的时间复杂
  下完成list两端的push和pop操作。但是因为每次操作都需要重新分配ziplist的内存，所以实际复杂度和ziplist的内存使用量相关

- 跳跃表（zSkiplist): 跳跃表的性能可以保证在查找，删除，添等操作的时候在对数期望时间内完成，这个性能是可以和平衡树来相比较的，而且在实现方面比平衡树要优雅，这是采用跳跃表的主要原因。跳跃表的复杂度是O(log(n))。

| 命令 | 简述 | 使用 |
|-------|-------|-------|
| ZADD | 将一个带有给定分值的成员添加到有序集合里面 | ZADD zset-key 178 member1 |
| ZRANGE | 根据元素在有序集合中所处的位置，从有序集合中获取多个元素 | ZRANGE zset-key 0-1 withccores|
| ZREM | 如果给定元素成员存在于有序集合中，那么就移除这个元素 | ZREM zset-key member1|

