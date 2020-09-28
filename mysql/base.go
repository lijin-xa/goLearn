package mysql

// 使用sqlx库
// 在Go Modules 里面找到sqlx.go源文件
// sqlx_test.go 测试用例

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // _ 匿名引用格式 只希望调用初始化的init函数
	"github.com/jmoiron/sqlx"          // sqlx 的官方文档
)

// sqlx.go文件
func Start() {
	initClient()
}

var Db *sqlx.DB

//type user struct {
//	Uid int 'db:"uid"',
//}

// 初始化连接
func initClient() {
	// localhost  127.0.0.1  本机Ip地址
	// localhost 等于 127.0.0.1 不过localhost是一个域名 可以理解为url地址 127.0.0.1是IP地址
	// localhost 和 127.0.0.1 不需要联网，都可以本机访问
	// 本机Ip需要联网，本机ip是本机对外开放访问的ip地址，也是和物理网卡绑定的IP地址
	// mysql账号的授权 - 可以允许root用户在特定ip进行登录 通过设置操作权限 - 来远程操作数据库表
	db, err := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	Db = db
	fmt.Println("connect success ", Db.DriverName())

	operateCURD()
}

// 接收玩家变量结构体
type user struct {
	Uid  int    `db:"uid"`
	Name string `db:"name"`
}

// 简单增删查改
func operateCURD() {
	//selectDemo()
	//insertDemo()
	//deleteDemo()
	//updateDemo()
	transactionDemo()
	//selectAllData()
}

// 查询数据练习
func selectDemo() {
	u := &user{}

	// 查询单行数据
	//sqlStr := "SELECT * FROM user limit 1"
	sqlStr := "SELECT * FROM user WHERE uid = ?"

	/*
	* func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {}
	* @param dest  接收查询结果 指针类型
	* @param query 查询的sql语句 SELECT * FROM table_name where id = ?
	* @param args 绑定参数的赋值 对应sql语句的?
	 */
	err := Db.Get(u, sqlStr, 2)
	if err != nil {
		fmt.Println("query failed: ", err)
		return
	}
	fmt.Printf("retrieve result uid:%d name:%s \n", u.Uid, u.Name)

	// 查询多行数据
	//var uList []user
	//uList := new([]user)
	uList := &[]user{}

	// sql语句 既可以使用占位符? 在程序中设置参数 也可以直接使用数字
	sqlStr = "SELECT * FROM user WHERE uid < ?"

	// 这里不再使用Get方法 dest参数 需要传递一个指针类型数据
	err = Db.Select(uList, sqlStr, 3)
	if err != nil {
		fmt.Println("retrieve multiple failed ", err)
		return
	}

	fmt.Println("retrieve multiple result")
	for _, v := range *uList {
		fmt.Printf("uid:%d %s\n", v.Uid, v.Name)
	}
}

// 插入数据练习
func insertDemo() {
	sqlStr := "INSERT INTO user (uid, name) VALUES (?, ?)"
	result, err := Db.Exec(sqlStr, 5, "lua")
	if err != nil {
		fmt.Printf("insert failed %s\n", err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed, err %s\n", err)
	}
	fmt.Printf("insert success LastInsertId: %d\n", insertId)
}

// 删除数据练习
func deleteDemo() {
	//sqlStr := "DELETE FROM user where uid = 5 limit 3"
	sqlStr := "DELETE FROM user where uid = ?"
	result, err := Db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err%+v\n", err)
	}
	/*
		type Result interface {
			LastInsertId() (int64, error)
			RowsAffected() (int64, error)
		}
	*/
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get lastInsertId falied, err %s\n", err)
	}
	fmt.Println("LastInsertId", id)

	row, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected falied, err %s\n", err)
	}
	fmt.Println("row", row)

	// 查看删除后数据
	selectAllData()
}

// 更新数据练习
func updateDemo() {
	sqlStr := "UPDATE user SET name = ? Where uid = ?"
	result, err := Db.Exec(sqlStr, "c++", 2)
	if err != nil {
		fmt.Printf("update failed, err %s\n", err)
		return
	}
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get row affected, err %s\n", err)
		return
	}
	fmt.Println(row, "row affected")
	selectAllData()
}

// 查询表中全部数据
func selectAllData() {
	sqlStr := "SELECT * FROM user order by uid"
	uList := new([]user)
	err := Db.Select(uList, sqlStr)
	if err != nil {
		fmt.Printf("retrieve data failed, err %s\n", err)
	}
	for _, v := range *uList {
		fmt.Println(v.Uid, v.Name)
	}
}

// 事务练习
func transactionDemo() {
	var result sql.Result
	var err error
	var tx *sql.Tx
	
	// db.Mustxx 有些必要操作必须宕机的可以考虑使用这类方法 MustExec和Exec实现功能类似
	tx, err = Db.Begin()
	if err != nil {
		fmt.Printf("start transaction failed, err %s\n", err)
	}

	// 设计一个延时函数 - 执行事务的提交Commit 回滚rollback 宕机时打印上下文信息
	defer func() {
		// 宕机中恢复
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil { // err 拿的值是从中断函数后执行这个 闭包匿名函数
			// 事务回滚
			_ = tx.Rollback()
			fmt.Println("Transaction Rollback")
		} else {
			// 事务提交
			tx.Commit()
			fmt.Println("Transaction Commit")
		}
	}()

	sqlStr1 := "UPDATE user Set uid = ? WHERE name = ?"
	result, err = tx.Exec(sqlStr1, 3, "go")
	if err != nil {
		fmt.Printf("1 update failed, err %s\n", err)
		return
	}
	rows1, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get rows affected failed")
		return
	}
	fmt.Printf("update rows %d\n", rows1)

	//sqlxStr2 := "UPDATE user Set uid = ? WHERE name = ?"
	result, err = tx.Exec(sqlStr1, 4, "lua")
	if err != nil {
		fmt.Printf("2 update failed, err %s\n", err)
		return
	}
	rows2, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get rows affected failed")
		return
	}
	fmt.Printf("update rows %d\n", rows2)
	if rows1 > 0 && rows2 > 0 {
		fmt.Println("update success")
	}
	return
}

// TODO 后续了解 参数绑定
// NameQuery NameExec
func selectNameQuery() {
	sqlStr := "SELECT * FROM user WHERE name = :name"
	Db.NamedQuery(sqlStr, map[string]interface{}{
		"name": "lijin",
	})
}

//type PlacePerson struct {
//	uid sql.NullInt32 `db:"uid"`
//	Name  sql.NullString `db:"name"`
//}

// 增删改查CURD操作 提示引入的事务 - 先暂时放下
//func operateCURD1() {
//	// 插入数据
//	tx := Db.MustBegin()
//
//	// 事务
//	tx.MustExec("insert into user (uid, name) values(?,?)", 3, "lijin-2")
//
//	tx.Commit()
//	tx.Rollback()
//}
