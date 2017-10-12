package chaos

import (
	"github.com/go-xorm/xorm"
	"github.com/kafrax/logx"
)

//mysql connect
var __ENGINE *xorm.Engine
var MYSQL_HOST=""
var MYSQL_SHOW=false
func MYSQLInstance() *xorm.Engine {
	if __ENGINE == nil {
		var err error
		__ENGINE, err = xorm.NewEngine("mysql", MYSQL_HOST)
		if err != nil {
			logx.Errorf("数据库连接初始化 |message=%v", err)
			return nil
		}
		__ENGINE.ShowSQL(MYSQL_SHOW)
		return __ENGINE
	}
	return __ENGINE
}
