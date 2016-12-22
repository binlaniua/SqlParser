

# 在做SQL审计的时候, 我们需要分析一条SQL访问了那些表，那些字段

以下是业务SQL的场景, 目前对中文支持不友好

## SELECT
```
p := sqlparser.NewSQLParser(`
    select
        t1.a,
        t1.b,
        t2.e,
        t2.f
    from
        table1 t1,
        table2 t2
    where
        t1.a = t2.a
`)
r, err := p.DoParser()
if err != nil {
    t.Error(err)
} else {
    t.Log(r.String())
}
```
```
{
    "*": {
        "table1": [
            "a",
            "b"
        ],
        "table2": [
            "e",
            "f"
        ]
    }
}
```

## INSERT
```
p := sqlparser.NewSQLParser(`
    insert into table1 (a,b,c,d) values (1,2,3,4)
`)
r, err := p.DoParser()
if err != nil {
    t.Error(err)
} else {
    t.Log(r.String())
}
```
```
{
    "*": {
        "table1": [
            "a",
            "b",
            "c",
            "d"
        ]
    }
}
```

## UPDATE
```
p := sqlparser.NewSQLParser(`
    update table1 set a = 1, b = 2, c = 3
`)
r, err := p.DoParser()
if err != nil {
    t.Error(err)
} else {
    t.Log(r.String())
}
```
```
{
    "*": {
        "table1": [
            "a",
            "b",
            "c"
        ]
    }
}
```

## UPDATE
```
p := sqlparser.NewSQLParser(`
    delete from table1 where a = 1
`)
r, err := p.DoParser()
if err != nil {
    t.Error(err)
} else {
    t.Log(r.String())
}
```
```
{
    "*": {
        "table1": []
    }
}
```