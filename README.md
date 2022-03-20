[toc]

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

