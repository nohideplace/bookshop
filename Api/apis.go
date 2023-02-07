package Api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

/*
1.Api层是逻辑上的入口，由一个APIs()函数统筹所有api，然后提供给main直接调用
2.在这一层内，将有相似性质的接口放在同一个文件内，如example，然后在一个文件内将同类api整合后在APIs()函数中调用即可
3.Service层负责逻辑上的处理，以及除了数据调用外的一切细节处理
4.DataHandler层负责数据库的查询，从文件中调取数据等操作，一切数据的获取必经层
5.Model层存放项目中所有的结构体和接口类型
6.ErrorHandler层用于编写错误类型，提供独立的错误处理
7.testfile包用于平常测试用，没有实际意义
*/

// APIs 注册api接口
func APIs(db *sql.DB, router *gin.Engine) *gin.Engine {
	router = Books(db, router)
	router = Comments(db, router)
	router = Users(db, router)
	router = UserOperation(db, router)
	return router
}
