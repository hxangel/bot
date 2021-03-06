package db

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	utils "github.com/hxangel/bot/libs/utils"
	"log"
	"os"
	"reflect"
)

var loger = log.New(os.Stdout, "[mysql] ", log.Ldate|log.Ltime)

//mysql
type Mysql struct {
	Config string
}

func NewMysql(config string) *Mysql {
	return &Mysql{Config: config}
}

//query all result
func (m *Mysql) FetchAll(query string, args ...interface{}) (results []map[string]interface{}, err error) {
	dsn := m.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}

	for rows.Next() {
		result, err := m.RowsToMap(rows)
		if err == nil {
			results = append(results, result)
		}
	}
	err = rows.Err()
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}
	return results, nil
}

//query all result
func (m *Mysql) Fetch(query string, args ...interface{}) (results map[string]interface{}, err error) {
	dsn := m.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}

	if !rows.Next() {
		loger.Print(rows.Err())
		return nil, rows.Err()
	}

	result, err := m.RowsToMap(rows)
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}
	return result, nil
}

func (m *Mysql) FetchRow(query string, args ...interface{}) (*sql.Row, error) {
	dsn := m.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(query, args...)
	return row, nil
}

//update and delete
func (m *Mysql) Execute(query string, args ...interface{}) (int64, error) {
	dsn := m.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		loger.Print(err.Error())
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		loger.Print(err.Error())
		return 0, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		loger.Print(err.Error())
		return 0, err
	}
	id, err := res.LastInsertId()
	if id > 0 {
		return id, err
	}
	affect, err := res.RowsAffected()
	return affect, err
}

func (m *Mysql) RowsToMap(rows *sql.Rows) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	fields, err := rows.Columns()
	if err != nil {
		loger.Print(err.Error())
		return nil, err
	}

	var containers []interface{}
	for i := 0; i < len(fields); i++ {
		var container interface{}
		containers = append(containers, &container)
	}
	if err := rows.Scan(containers...); err != nil {
		return nil, err
	}

	for index, field := range fields {
		rawValue := reflect.Indirect(reflect.ValueOf(containers[index]))

		if rawValue.Interface() != nil {
			result[field] = rawValue.Interface()
		}
	}
	return result, nil
}

func (m *Mysql) DSN() string {
	username := utils.GetConfig("mysql", "username")
	password := utils.GetConfig("mysql", "password")
	host := utils.GetConfig("mysql", "host")
	port := utils.GetConfig("mysql", "port")
	dbname := utils.GetConfig("mysql", "dbname")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, dbname)
}
