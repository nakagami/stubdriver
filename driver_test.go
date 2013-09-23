// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
    "testing"
    "fmt"
    "database/sql"
)

func TestConnect(t *testing.T) {
    fmt.Println("1")
    db, err := sql.Open("stubdriver", "dataSourceName")
    if err != nil {
        t.Fatalf("Error connecting: %v", err)
    }
    defer db.Close()

    fmt.Println("2")
    r, err2 := db.Exec("INSERT foo, bar")
    if err2 != nil {
        t.Fatalf("Error Exec(): %v", err2)
    }

    fmt.Println("3")
    tx, _ := db.Begin()
    r, _ = tx.Exec("UPDATE with transaction1")
    tx.Rollback()

    fmt.Println("4")
    r, _ = tx.Exec("UPDATE with transaction2")
    fmt.Println("\tResult:", r)

    fmt.Println("5")
    tx, _ = db.Begin()
    r, _ = tx.Exec("UPDATE with transaction3")
    fmt.Println("\tResult:", r)

    fmt.Println("6")
    tx.Commit()

    fmt.Println("7")
    var n int
    var s string
    rows, _ := db.Query("select column1, column2 from foo")

    columns, _ := rows.Columns()
    fmt.Println("\tColumns:", columns)

    for rows.Next() {
        rows.Scan(&n, &s)
        fmt.Println("\t", n, s)
    }

    fmt.Println("8")
    stmt, _ := db.Prepare("INSERT (?,?) ")
    stmt.Exec(1,2)

    fmt.Println("9")
    stmt, _ = db.Prepare("SELECT where id=? and name=? ")

    rows, _ = stmt.Query(1, "ABC")

    for rows.Next() {
        rows.Scan(&n, &s)
        fmt.Println("\t", n, s)
    }

    fmt.Println("------------")
}
