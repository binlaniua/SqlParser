

# 在做SQL审计的时候, 分析一条SQL访问了那些表，那些字段
* 可以做敏感字段过滤
* 可以做表权限
* 等等.....

以下是业务SQL的场景, 目前对中文支持不友好

## 使用方法
```
p := sqlparser.NewSQLParser(`
    select
        t1.a,
        t1.b,
        t3.e ccc,
        t3.f ddd
    from
        xx.table1 t1,
        (select t2.b as e, t2.d as f from (select a as b, c as d from yy.table2) t2) t3
`)
r, err := p.DoParser()
if err != nil {
    t.Error(err)
} else {
    t.Log(r.String())
}
```

## 结果
```
意思查询
    表用户 xx
        表 table1
            表字段
                a 无别名
                b 无别名
    别用户 yy
        表 table2
            表字段
                a -> b -> e -> ccc (最终查询出来是ccc)
                c -> d -> f -> ddd (最终查询出来是ddd)
{
    "xx": {
        "Name": "xx",
        "TableMap": {
            "table1": {
                "Name": "table1",
                "Alias": "t1",
                "ColumnMap": {
                    "a": {
                        "Name": "a",
                        "Alias": {
                            "Name": "",
                            "Alias": null
                        }
                    },
                    "b": {
                        "Name": "b",
                        "Alias": {
                            "Name": "",
                            "Alias": null
                        }
                    }
                }
            }
        }
    },
    "yy": {
        "Name": "yy",
        "TableMap": {
            "table2": {
                "Name": "table2",
                "Alias": "t3",
                "ColumnMap": {
                    "a": {
                        "Name": "a",
                        "Alias": {
                            "Name": "b",
                            "Alias": {
                                "Name": "e",
                                "Alias": {
                                    "Name": "ccc",
                                    "Alias": null
                                }
                            }
                        }
                    },
                    "c": {
                        "Name": "c",
                        "Alias": {
                            "Name": "d",
                            "Alias": {
                                "Name": "f",
                                "Alias": {
                                    "Name": "ddd",
                                    "Alias": null
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```