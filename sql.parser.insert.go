package sqlparser

import (
	"github.com/xwb1989/sqlparser"
)

//-------------------------------------
//
//
//
//-------------------------------------
type insertSql struct {
	result   *SQLParserResult
	node     *sqlparser.Insert
	tableMap map[string]*DBTable
}

//
//
//
//
//
func newInsertSql(r *SQLParserResult, n *sqlparser.Insert) *insertSql {
	return &insertSql{
		r,
		n,
		map[string]*DBTable{},
	}
}

//
//
//
//
//
func (ss *insertSql) doParser() error {
	ss.visitSimpleTable(ss.node.Table, "")
	for _, column := range ss.node.Columns {
		newColumn(ss.result, column, ss.tableMap).doParser()
	}
	return nil
}

//
//
//
//
//
func (ss *insertSql) visitSimpleTable(table *sqlparser.TableName, alias string) {
	dbOwner := string(table.Qualifier)
	tableName := string(table.Name)
	dbTable := ss.result.AddTable(dbOwner, tableName, alias)
	ss.tableMap[alias] = dbTable
}
