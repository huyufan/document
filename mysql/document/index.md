# MySQL 单表不要超过 2000W 行，靠谱吗？

## 这张表数据，在硬盘上存储也是类似如此的，它实际是放在一个叫 person.ibd （innodb data）的文件中，也叫做表空间；虽然数据表中，他们看起来是一条连着一条，但是实际上在文件中它被分成很多小份的数据页，而且每一份都是 16K。
## 大概就像下面这样，当然这只是我们抽象出来的，在表空间中还有段、区、组等很多概念，但是我们需要跳出来看。

## 单表建议值
- 非叶子节点内指向其他页的数量为 x
- 叶子节点内能容纳的数据行数为 y
- B+ 数的层数为 z

### Total =x^(z-1) *y 也就是说总数会等于 x 的 z-1 次方 与 Y 的乘积

#### x =?
- 已经介绍了页的结构，索引也也不例外，都会有 File Header (38 byte)、Page Header (56 Byte)、Infimum + Supermum（26 byte）、File Trailer（8byte）, 再加上页目录，大概 1k 左右。
- 我们就当做它就是 1K, 那整个页的大小是 16K, 剩下 15k 用于存数据，在索引页中主要记录的是主键与页号，主键我们假设是 Bigint (8 byte), 而页号也是固定的（4Byte）, 那么索引页中的一条数据也就是 12byte。
##### **所以 x=15*1024/12≈1280 行。**

#### Y=?
- 叶子节点和非叶子节点的结构是一样的，同理，能放数据的空间也是 15k。
- 但是叶子节点中存放的是真正的行数据，这个影响的因素就会多很多，比如，字段的类型，字段的数量。每行数据占用空间越大，页中所放的行数量就会越少。
- 这边我们暂时按一条行数据 1k 来算，那一页就能存下 15 条，Y = 15*1024/1000  ≈15。
- 根据上述的公式，Total =x^(z-1) *y，已知 x=1280，y=15：

-  假设 B+ 树是两层，那就是 z = 2， Total = （1280 ^1 ）*15 = 19200
- 假设 B+ 树是两层，那就是 z = 2， Total = （1280 ^1 ）*15 = 19200

