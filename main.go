package sqlparser

import (
	"errors"
	"fmt"
	"github.com/binlaniua/kitgo"
	"github.com/xwb1989/sqlparser"
	"reflect"
	"regexp"
)

var (
	pattern_chinese = regexp.MustCompile("[\u4E00-\u9FA5]+")
	pattern_as      = regexp.MustCompile(`(?i)as ('|")([^'"]+)('|")`)
)

//-------------------------------------
//
//
//
//-------------------------------------
type SQLParser struct {
	sql             string
	chineseFieldMap map[string]string
	result          *SQLParserResult
}

//
//
//
//
//
func NewSQLParser(sqlString string) *SQLParser {
	//
	sp := &SQLParser{
		sql:             sqlString,
		chineseFieldMap: map[string]string{},
		result:          NewSQLparserResult(),
	}

	//
	sp.cleanSql()
	return sp
}

//
//
// 获取解析结果
//
//
func (sp *SQLParser) GetResult() *SQLParserResult {
	return sp.result
}

//
//
//
//
//
func (sp *SQLParser) DoParser() (*SQLParserResult, error) {
	ast, err := sqlparser.Parse(sp.sql)
	if err != nil {
		return nil, err
	}

	switch node := ast.(type) {
	//
	case *sqlparser.Select:
		sp.result.sqlType = SQL_TYPE_SELECT
		err = newSelectSql(sp.result, "", node).doParser()

	//
	case *sqlparser.Union:
		sp.result.sqlType = SQL_TYPE_UNION
		err = newUnionSql(sp.result, "", node).doParser()

	//
	case *sqlparser.Insert:
		sp.result.sqlType = SQL_TYPE_INSERT
		err = newInsertSql(sp.result, node).doParser()

	//
	case *sqlparser.Update:
		sp.result.sqlType = SQL_TYPE_UPDATE
		err = newUpdateSql(sp.result, node).doParser()

	//
	case *sqlparser.Delete:
		sp.result.sqlType = SQL_TYPE_DEL
		err = newDeleteSql(sp.result, node).doParser()

	//
	default:
		return nil, errors.New(fmt.Sprintf("不支持类型 => %s", reflect.TypeOf(ast)))
	}
	return sp.result, nil
}

//
//
// 清理SQL
// 1. 不支持中文   not support chinese
// 2. 不支持函数   not support database function
// 3. AS 后面不能更引号 not support "as 'xxx'", support "as xxxx"
// 4. 不支持(+)连接 not support left join use (+)
//
//
func (sp *SQLParser) cleanSql() {
	count := 0
	sp.sql = pattern_chinese.ReplaceAllStringFunc(sp.sql, func(src string) string {
		alias := fmt.Sprintf("__r%d", count)
		sp.chineseFieldMap[alias] = src
		count++
		return alias
	})
	sp.sql = pattern_as.ReplaceAllStringFunc(sp.sql, func(src string) string {
		return kitgo.StringReplace(src, `"|'`, "")
	})
	sp.sql = kitgo.StringReplace(sp.sql, ";|\\(\\+\\)", "")
}
