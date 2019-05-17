package _db

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"strconv"
	"fmt"
)

func MysqlDB(dbUrl, dbName, user, password string, port int) {
	connURL := strings.Join([]string{user, ":", password, "@", "tcp(", dbUrl, ":", strconv.Itoa(port), ")/", dbName, "?charset=utf8"}, "")
	fmt.Println(connURL)
	db, err := sql.Open("mysql", connURL)
	checkErr(err)
	defer db.Close()

	insertData(db)

	deleteById(db, 21)

	update(db, "%忘然%")

	query(db)
}

func query(db *sql.DB) {
	rows, err := db.Query("select Id, Name from T_Person")
	defer rows.Close()
	checkErr(err)

	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}

func insertData(db *sql.DB) {
	stmt, err := db.Prepare("insert into T_Person(Name, Age, Address, UpdateBy) values (?, ?, ?, ?)")
	checkErr(err)

	rs, err := stmt.Exec("卓忘然", 33, "湾子晓月楼", "dante")
	checkErr(err)

	insertId, err := rs.LastInsertId()
	checkErr(err)

	fmt.Println("新增Id: ", insertId)
}

func deleteById(db *sql.DB, id int32) {
	stmt, err := db.Prepare("delete from T_Person where id = ?")
	checkErr(err)

	rs, err := stmt.Exec(id)
	checkErr(err)

	delRow, err := rs.RowsAffected()
	checkErr(err)

	fmt.Println("删除行数: ", delRow)
}

func update(db *sql.DB, name string) {
	stmt, err := db.Prepare("update T_Person set Remark = ? where Name like ?")
	checkErr(err)

	rs, err := stmt.Exec("我是通过Go语言操作！", name)
	checkErr(err)

	updateRow, err := rs.RowsAffected()
	checkErr(err)

	fmt.Println("更新行数: ", updateRow)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}