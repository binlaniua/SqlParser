package sqlparser

import (
	"github.com/binlaniua/kitgo"
	"github.com/xwb1989/sqlparser"
	"reflect"
)

//-------------------------------------
//
// 字段的解析
//
//-------------------------------------
type column struct {
	result   *SQLParserResult
	expr     sqlparser.SelectExpr
	tableMap map[string]*DBTable
}

//
//
//
//
//
func newColumn(r *SQLParserResult, n sqlparser.SelectExpr, tableMap map[string]*DBTable) *column {
	return &column{
		r,
		n,
		tableMap,
	}
}

//
//
//
//
//
func (ss *column) doParser() error {
	tableMap := ss.tableMap
	switch f := ss.expr.(type) {

	// 非 * 号
	case *sqlparser.NonStarExpr:
		switch c := f.Expr.(type) {
		//表达式
		case *sqlparser.ColName:
			ss.addToResult(string(c.Name), string(f.As), string(c.Qualifier))

		//带函数的表达式
		case *sqlparser.FuncExpr:
			for _, e := range c.Exprs {
				newColumn(ss.result, e, tableMap).doParser()
			}
		//
		case sqlparser.ValTuple:
			for _, v := range c {
				ss.visitValExp(v)
			}

		// case when
		case *sqlparser.CaseExpr:
			ss.visitValExp(c.Else)
			ss.visitValExp(c.Expr)

		//非表达式, 不管
		case sqlparser.StrVal:
		case sqlparser.NumVal:
		case *sqlparser.BinaryExpr:
		default:
			kitgo.ErrorLog.Println(reflect.TypeOf(f.Expr))
		}

	// * 号
	case *sqlparser.StarExpr:
		ss.addToResult("*", "*", string(f.TableName))
	}
	return nil
}

//
//
// 解析表达式
//
//
func (sp *column) visitValExp(exp sqlparser.ValExpr) {
	if exp == nil {
		return
	}
	switch f := exp.(type) {
	case *sqlparser.BinaryExpr:
	default:
		kitgo.ErrorLog.Println(f)
	}
}

//
//
//
//
//
func (sp *column) addToResult(col, colAlias, tableAlias string) {
	dbTable := sp.tableMap[tableAlias]
	if IS_DEBUG {
		tableName := ""
		if dbTable != nil {
			tableName = dbTable.Name
		}
		kitgo.DebugLog.Printf("添加表列 => [%s] [%s] [%s] [%s]", col, colAlias, tableAlias, tableName)
	}
	sp.result.AddCol(col, colAlias, dbTable)
}
