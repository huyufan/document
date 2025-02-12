# 查询和聚合的基础使用

``` shell 
//match_all表示查询所有的数据，sort即按照什么字段排序 分页查询(from+size)
get /bank/_search 
{
 "query" :{"match_all":{}},
 "sort":[
    {"account_number":"desc"}
 ],
 "from":0,
 "size": 10
}

//指定字段查询：match 如果要在字段中搜索特定字词，可以使用match; 如下语句将查询address 字段中包含 mill 或者 lane的数据
get /bank/_search
{
    "query":{"match":{"address":"mill lane"}}
}

// 查询段落匹配：match_phrase 如果我们希望查询的条件是 address字段中包含 "mill lane"，则可以使用match_phrase
get /bank/_search 
{
    "query":{"match_phrase":{"address":"mill lane"}}
}

//多条件查询: bool 例如，以下请求在bank索引中搜索40岁客户的帐户，但不包括居住在爱达荷州（ID）的任何人
get /bank/_search
{
    "query": {
        "bool":{
            "must": [
              {"match":{"age":"40"}}
            ],
            "must_not": [
              {"match":{"state":"ID"}}
            ]
        }
    }
}

//查询条件：query or filter query 上下文的条件是用来给文档打分的，匹配越好 _score 越高；filter 的条件只产生两种结果：符合与不符合，后者被过滤掉
get /bank/_search 
{
    "query":{
        "bool":{
            "must":[
                {"match":{"state":"ND"}}
            ],
            "filter":[
                {"term": {"age":"40"}},
                {"range":{
                    "balance": {
                        "gte":20000,
                        "lte": 30000
                    }
                }}
            ]
        }
    }
}

GET /bank/_search 
{
    "query":{
        "bool":{
            "filter": [
              {
                "term":{"age":"40"}
              },
              {
                "range":{
                    "balance": {
                        "gte": 20000,
                        "lte": 30000
                    }
                }
              }
            ]
        }
    }
}

//简单聚合 比如我们希望计算出account每个州的统计数量， 使用aggs关键字对state字段聚合，被聚合的字段无需对分词统计，所以使用state.keyword对整个字段统计 计算每个州的平均结余。涉及到的就是在对state分组的基础上，嵌套计算avg(balance):
GET /bank/_search
{
    "size":0,
"aggs":{
    "group_by_state":{
        "terms":{
            "field":"state.keyword",
            "order": {
              "banlane": "desc"
            }
        },
        "aggs": {
          "banlane": {
            "avg": {
              "field": "balance"
            }
          }
        }
        }
        }
}

```