package sqlparser

import (
	"github.com/xwb1989/sqlparser"
)

//-------------------------------------
//
//
//
//-------------------------------------
type deleteSql struct {
	result   *SQLParserResult
	node     *sqlparser.Delete
	tableMap map[string]*DBTable
}

//
//
//
//
//
func newDeleteSql(r *SQLParserResult, n *sqlparser.Delete) *deleteSql {
	return &deleteSql{
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
func (ss *deleteSql) doParser() error {
	ss.visitSimpleTable(ss.node.Table, "")
	return nil
}

//
//
//
//
//
func (ss *deleteSql) visitSimpleTable(table *sqlparser.TableName, alias string) {
	dbOwner := string(table.Qualifier)
	tableName := string(table.Name)
	dbTable := ss.result.AddTable(dbOwner, tableName, alias)
	ss.tableMap[alias] = dbTable
}
