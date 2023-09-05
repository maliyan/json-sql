[toc]

# 思路

## 实现方式：

> 本质上是需要构造一种数据结构，实现查找极快的查询速度。

- 将json内容读取出来，其实不需要关注json如何转成Go的结构体，因为，我们最终的目的是**要查找**，所以，构造的数据结构应该能够满足极快的查找速度。





## 语法：

- 查询：
  - `select`：
    - `select id, foo.id`：按照
    - **ToDo:** alias支持

  - `from`：
    - 感觉这个没啥用途

  - `where`：
    - `and`：
    - `or`：

  - `limit`：

- 输出：?



# json-sql

> 挖个坑，埋点土，数个12345，种个JSON-SQL.
>
> `select foo from JSONData where bar = "xxx"`

```json
{
    "err_code": 0,
    "err_msg": "OK",
    "data":
    [
        {
            "id": 1,
            "name": "aaa"
        },
        {
            "id": 2,
            "name": "bbb"
        }
    ]
}
```

```sql
---
select err_code from 数据源名字 
```

