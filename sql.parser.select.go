package sqlparser

import (
	"github.com/binlaniua/kitgo"
	"github.com/xwb1989/sqlparser"
	"reflect"
	"strings"
)

//-------------------------------------
//
//
//
//-------------------------------------
type selectSql struct {
	alias         string
	result        *SQLParserResult
	node          *sqlparser.Select
	tableMap      map[string]*DBTable
	tableAliasMap map[string]string
}

//
//
//
//
//
func newSelectSql(r *SQLParserResult, alias string, n *sqlparser.Select) *selectSql {
	return &selectSql{
		alias,
		r,
		n,
		map[string]*DBTable{},
		map[string]string{},
	}
}

//
//
//
//
//
func (ss *selectSql) doParser() error {
	// 先解表
	for _, table := range ss.node.From {
		ss.visitForm(table, ss.alias)
	}

	// 再解字段
	for _, field := range ss.node.SelectExprs {
		newColumn(ss.result, field, ss.tableMap).doParser()
	}
	return nil
}

//
//
// 分析表
//
//
func (ss *selectSql) visitForm(table sqlparser.TableExpr, nodeAlias string) {
	kitgo.DebugLog.Print(strings.Repeat("=", 20))
	switch t := table.(type) {

	//父表
	case *sqlparser.ParenTableExpr:
		kitgo.DebugLog.Print("ParenTableExpr")
		ss.visitForm(t.Expr, nodeAlias)

	//左右表
	case *sqlparser.JoinTableExpr:
		kitgo.DebugLog.Print("JoinTableExpr")
		ss.visitForm(t.LeftExpr, nodeAlias)
		ss.visitForm(t.RightExpr, nodeAlias)

	//真实表
	case *sqlparser.AliasedTableExpr:
		kitgo.DebugLog.Print("AliasedTableExpr")
		// 加上表别名映射
		alias := string(t.As)
		if nodeAlias != "" {
			ss.tableAliasMap[alias] = nodeAlias
		}
		ss.visitTable(t.Expr, string(t.As))

		kitgo.DebugLog.Printf("[%s] => [%s]", alias, nodeAlias)

	//
	default:
		kitgo.ErrorLog.Print(reflect.TypeOf(table))
	}
}

//
//
// 分析查询表
//
//
func (ss *selectSql) visitTable(table sqlparser.SimpleTableExpr, alias string) {
	switch t := table.(type) {
	//简单的表
	case *sqlparser.TableName:
		ss.visitSimpleTable(t, alias)

	//子查询
	case *sqlparser.Subquery:
		ss.visitQuery(t.Select, alias)
	}
}

//
//
//
//
//
func (ss *selectSql) visitSimpleTable(table *sqlparser.TableName, alias string) {
	dbOwner := string(table.Qualifier)
	tableName := string(table.Name)
	dbTable := ss.result.AddTable(dbOwner, tableName, alias)
	ss.tableMap[alias] = dbTable

	//反向找表别名, 把表别名指向真实表
	for {
		newAlias, ok := ss.tableAliasMap[alias]
		if ok {
			ss.tableMap[newAlias] = dbTable
			alias = newAlias
		} else {
			break
		}
	}

	kitgo.DebugLog.Print(ss.tableAliasMap)
	kitgo.DebugLog.Print(ss.tableMap)
}

//
//
// 分析查询
//
//
func (ss *selectSql) visitQuery(node sqlparser.SelectStatement, nodeAlias string) {
	switch s := node.(type) {
	case *sqlparser.Select:
		nss := newSelectSql(ss.result, nodeAlias, s)
		nss.tableMap = ss.tableMap
		nss.tableAliasMap = ss.tableAliasMap
		nss.doParser()
	case *sqlparser.Union:
		newUnionSql(ss.result, nodeAlias, s).doParser()
	}
}
