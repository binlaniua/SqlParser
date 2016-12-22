package sqlparser

import (
	"encoding/json"
	"github.com/binlaniua/kitgo"
)

//-------------------------------------
//
//
//
//-------------------------------------
type SQLParserResult struct {

	//
	//
	// SQL_TYPE
	//
	//
	sqlType    string

	//
	//
	// 表用户
	//  	表
	//	  表字段
	//	  表字段
	//	表
	result     map[string]map[string][]string

	//
	//
	// 表别名
	//	表用户, 表名(或者表别名)
	//	表用户, 表名(或者表别名)
	//
	tableAlias map[string][]string
}

//
//
//
//
//
func NewSQLparserResult() *SQLParserResult {
	return &SQLParserResult{
		sqlType: "",
		result: map[string]map[string][]string{},
		tableAlias: map[string][]string{},
	}
}

//
//
//
//
//
func (spr *SQLParserResult) AddResult(dbOwner string, table string, col string) *SQLParserResult {
	//kitgo.DebugLog.Printf("添加信息 => %s %s %s", dbOwner, table, col)
	m, ok := spr.result[dbOwner]
	if !ok {
		m = map[string][]string{}
		spr.result[dbOwner] = m
	}
	t, ok := m[table]
	if !ok {
		t = []string{}
	}

	if col != "" {
		m[table] = append(t, col)
	} else {
		m[table] = []string{}
	}
	return spr
}

//
//
// 添加表
//
//
func (spr *SQLParserResult) AddTable(dbOwner string, table string, tableAlias string) *SQLParserResult {
	if dbOwner == "" {
		dbOwner = "*"
	}
	kitgo.DebugLog.Printf("添加表信息 => %s %s %s", dbOwner, table, tableAlias)
	spr.AddResult(dbOwner, table, "")
	if tableAlias != "" {
		spr.tableAlias[tableAlias] = []string{
			dbOwner,
			table,
		}
	}
	return spr
}

//
//
// 添加表别名
//
//
func (spr *SQLParserResult) AddTableAlias(table string, tableAlias string) *SQLParserResult {
	kitgo.DebugLog.Printf("添加表别名 => %s %s", table, tableAlias)
	if tableAlias != "" {
		spr.tableAlias[tableAlias] = []string{
			"",
			table,
		}
	}
	return spr
}

//
//
// 添加列
//
//
func (spr *SQLParserResult) AddCol(col string, colAlias string, tableAlias string) {
	isMatch := false
	if tableAlias == "" {
		//找第一个表用户的第一个表...
		for k1, v1 := range spr.result {
			for k2 := range v1 {
				isMatch = true
				spr.AddResult(k1, k2, col)
			}
		}
	} else {
		var currentAlias []string
		for {
			ts, ok := spr.tableAlias[tableAlias]
			if ok {
				//因为这个取出来有可能是表别名, 需要再取一次, 直到取不到为止
				currentAlias = ts
				tableAlias = ts[1]
			} else {
				break
			}
		}
		if currentAlias != nil {
			isMatch = true
			spr.AddResult(currentAlias[0], currentAlias[1], col)
		}
	}

	if !isMatch {
		kitgo.ErrorLog.Printf("表字段没有匹配到 => %s %s %s", col, colAlias, tableAlias)
	}
}

//
//
//
//
//
func (spr *SQLParserResult) String() string {
	r, _ := json.MarshalIndent(spr.result, "", "    ")
	return string(r)
}

//
//
//
//
//
func (spr *SQLParserResult) GetOwner(key string) map[string][]string {
	r := spr.result[key]
	return r
}

//
//
//
//
//
func (spr *SQLParserResult) IsEmpty() bool {
	return len(spr.result) == 0
}