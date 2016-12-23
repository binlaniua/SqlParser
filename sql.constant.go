package sqlparser

//-------------------------------------
//
//  配置
//
//-------------------------------------
var (
	IS_DEBUG = false
)

//-------------------------------------
//
//  SQL TYPE
//
//-------------------------------------
const (
	SQL_TYPE_SELECT = "select"
	SQL_TYPE_UNION  = "union"
	SQL_TYPE_INSERT = "insert"
	SQL_TYPE_UPDATE = "update"
	SQL_TYPE_DEL    = "del"
)
