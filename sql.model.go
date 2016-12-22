package sqlparser

import (
	"encoding/json"
	"github.com/binlaniua/kitgo"
)

//-------------------------------------
//
//  表用户
//
//-------------------------------------
type DBUser struct {
	Name     string
	TableMap map[string]*DBTable
}

//-------------------------------------
//
//  表
//
//-------------------------------------
type DBTable struct {
	Name      string
	Alias     string
	DBUser    *DBUser `json:"-"`
	ColumnMap map[string]*DBTableColumn
}

//-------------------------------------
//
//  表字段
//
//-------------------------------------
type DBTableColumn struct {
	Name    string
	Alias   string
	DBTable *DBTable `json:"-"`
}

//-------------------------------------
//
//
//
//-------------------------------------
type SQLParserResult struct {
	//
	// SQL_TYPE
	//
	sqlType string

	//
	//
	//
	userMap map[string]*DBUser
}

//
//
//
//
//
func NewSQLparserResult() *SQLParserResult {
	return &SQLParserResult{
		sqlType: "",
		userMap: map[string]*DBUser{},
	}
}

//
//
//
//
//
func (spr *SQLParserResult) AddResult(dbOwner string, table string, col string) *DBUser {
	if IS_DEBUG {
		kitgo.DebugLog.Printf("add => [%s] [%s] [%s]", dbOwner, table, col)
	}

	// add db user
	dbUser, ok := spr.userMap[dbOwner]
	if !ok {
		dbUser = &DBUser{
			Name: dbOwner,
			TableMap: map[string]*DBTable{},
		}
		spr.userMap[dbOwner] = dbUser
	}

	// add db table
	dbTable, ok := dbUser.TableMap[table]
	if !ok {
		dbTable = &DBTable{
			DBUser: dbUser,
			Name: table,
			ColumnMap: map[string]*DBTableColumn{},
		}
		dbUser.TableMap[table] = dbTable
	}

	// add db table column
	if col != "" {
		dbTable.ColumnMap[col] = &DBTableColumn{
			Name: col,
			DBTable: dbTable,
		}
	}
	return dbUser
}

//
//
// 添加表
//
//
func (spr *SQLParserResult) AddTable(dbOwner string, table string, tableAlias string) *DBTable {
	if dbOwner == "" {
		dbOwner = "*"
	}
	if IS_DEBUG {
		kitgo.DebugLog.Printf("add table  => [%s] [%s] [%s]", dbOwner, table, tableAlias)
	}

	//
	dbUser := spr.AddResult(dbOwner, table, "")
	dbTable := dbUser.TableMap[table]
	if dbTable == nil {
		kitgo.ErrorLog.Printf("[%s] [%s] => nil", dbOwner, table)
	} else if tableAlias != "" {
		dbTable.Alias = tableAlias
	}
	return dbTable
}

//
//
// 添加表别名
//
//
func (spr *SQLParserResult) AddTableAlias(table string, tableAlias string) *SQLParserResult {
	if IS_DEBUG {
		kitgo.DebugLog.Printf("set table alias => [%s] [%s]", table, tableAlias)
	}
	if tableAlias != "" {
		for _, dbUser := range spr.userMap {
			for _, dbTable := range dbUser.TableMap {
				if dbTable.Name == table {
					dbTable.Alias = tableAlias
				}
			}
		}
	}
	return spr
}

//
//
// 添加列
//
//
func (spr *SQLParserResult) AddCol(col string, colAlias string, dbTable *DBTable) {
	if IS_DEBUG {
		kitgo.DebugLog.Printf("add table column  => [%s] [%s]", col, colAlias, dbTable)
	}

	isMatch := false
	if dbTable != nil {
		isMatch = true
		dbTable.ColumnMap[col] = &DBTableColumn{
			Name: col,
			Alias: colAlias,
			DBTable: dbTable,
		}
	}
	if !isMatch {
		kitgo.ErrorLog.Printf("表字段没有匹配到 => [%s] [%s]", col, colAlias)
	}
}

//
//
//
//
//
func (spr *SQLParserResult) String() string {
	r, _ := json.MarshalIndent(spr.userMap, "", "    ")
	return string(r)
}

//
//
//
//
//
func (spr *SQLParserResult) GetDBUser(key string) *DBUser {
	r := spr.userMap[key]
	return r
}

//
//
//
//
//
func (spr *SQLParserResult) IsEmpty() bool {
	return len(spr.userMap) == 0
}