# in
- In 用于检查左边的表达式是否存在右边的列表或子查询的结果集中，如果存在 则IN 返回True  否则返回FALSE
``` mysql
SELECT column_name(s) FROM table_name WHERE column_name IN(value1,value2,....)
```
``` mysql
SELECT column_name(s) FROM table_name where column_name IN(select column_name(s) from another_table where condition)
```

# EXISTS
-  用于判断子查询是否至少能返回一行数据。它不关心子查询返回什么数据，只关心是否有结果。如果子查询有结果，则 EXISTS  返回 TRUE，否则返回 FALSE
``` mysql
SELECT * FROM Customers WHERE EXISTS(SELECT 1 FROM Orders Where Orders.CustomerID = Customers.CustomerID)
```
## 性能差异：在很多情况下，EXISTS  的性能优于 IN ，特别是当子查询的表很大时。这是因为  一旦找到匹配项就会立即停止查询，而 IN 可能会扫描整个子查询结果集。
## 使用场景: 如果子查询结果集较小且不频繁变动，IN, 可能更直观易懂。而当子查询涉及外部查询的每一行判断，并且子查询的效率较高时，EXISTS 更为合适
## NULL值处理: IN  能够正确处理子查询中包含NULL值的情况，而 EXISTS 不受子查询结果中NULL值的影响，因为它关注的是行的存在性，而不是具体值。