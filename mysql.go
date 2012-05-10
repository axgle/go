/**
install:
go get github.com/ziutek/mymysql/thrsafe
go get github.com/ziutek/mymysql/autorc
go get github.com/ziutek/mymysql/godrv

this hello world  sample no need to create table schema,just connect a database.
on windows, it is work fine!

more see:
https://github.com/ziutek/mymysql
*/
package main

import (
     "fmt"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    // _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

func main() {
    user:="root"
	pass:="root"
	dbname:="test"
	
    db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)

    err := db.Connect()
    if err != nil {
        panic(err)
    }

    rows, res, err := db.Query("select 'hello world' as  FirstColumn,22 as SecondColumn limit  %d", 20)
    if err != nil {
        panic(err)
    }

    for _, row := range rows {
        for _, col := range row {
            if col == nil {
                // col has NULL value
            } else {
                // Do something with text in col (type []byte)
            }
        }
 

        // You may get values by column name
        first := res.Map("FirstColumn")
        //second := res.Map("SecondColumn")
        fmt.Printf( row.Str(first)); 
    }
}