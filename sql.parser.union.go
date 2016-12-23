package sqlparser

import (
	"github.com/xwb1989/sqlparser"
)

//-------------------------------------
//
//
//
//-------------------------------------
type unionSql struct {
	alias  string
	result *SQLParserResult
	node   *sqlparser.Union
}

//
//
//
//
//
func newUnionSql(r *SQLParserResult, alias string, n *sqlparser.Union) *unionSql {
	return &unionSql{
		alias,
		r,
		n,
	}
}

//
//
//
//
//
func (ss *unionSql) doParser() error {
	ss.visitQuery(ss.node.Left)
	ss.visitQuery(ss.node.Right)
	return nil
}

//
//
// 分析查询
//
//
func (sp *unionSql) visitQuery(node sqlparser.SelectStatement) {
	switch s := node.(type) {
	case *sqlparser.Select:
		newSelectSql(sp.result, sp.alias, s).doParser()
	case *sqlparser.Union:
		newUnionSql(sp.result, sp.alias, s).doParser()
	}
}
