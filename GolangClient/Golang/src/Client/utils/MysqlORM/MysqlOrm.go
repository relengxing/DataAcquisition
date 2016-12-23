package MysqlORM

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
)

var Db *sql.DB

func Connect(username string,password string,table string)  {
    var err error
    Db, err = sql.Open("mysql", username+":"+password+"@/"+table+"?charset=utf8")
    if err != nil {
        fmt.Println("failed to open database:", err.Error())
        return
    }
}

func Close(){
    defer Db.Close()
}
