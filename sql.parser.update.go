package sqlparser

import (
	"github.com/xwb1989/sqlparser"
)

//-------------------------------------
//
//
//
//-------------------------------------
type updateSql struct {
	result   *SQLParserResult
	node     *sqlparser.Update
	tableMap map[string]*DBTable
}

//
//
//
//
//
func newUpdateSql(r *SQLParserResult, n *sqlparser.Update) *updateSql {
	return &updateSql{
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
func (ss *updateSql) doParser() error {
	ss.visitSimpleTable(ss.node.Table, "")
	for _, column := range ss.node.Exprs {
		tableAlias := string(column.Name.Qualifier)
		dbTable := ss.tableMap[tableAlias]
		ss.result.AddCol(string(column.Name.Name), "", dbTable)
	}
	return nil
}

//
//
//
//
//
func (ss *updateSql) visitSimpleTable(table *sqlparser.TableName, alias string) {
	dbOwner := string(table.Qualifier)
	tableName := string(table.Name)
	dbTable := ss.result.AddTable(dbOwner, tableName, alias)
	ss.tableMap[alias] = dbTable
}
