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
    db, err := sql.Open("stubdriver", "dataSourceName")
    if err != nil {
        t.Fatalf("Error connecting: %v", err)
    }
    defer db.Close()

    r, err2 := db.Exec("SELECT foo, bar")
    if err2 != nil {
        t.Fatalf("Error Exec(): %v", err2)
    }

    tx, _ := db.Begin()
    r, _ = tx.Exec("SELECT with transaction1")
    tx.Rollback()

    r, _ = tx.Exec("SELECT with transaction2")
    fmt.Println("\tResult:", r)
    tx.Commit()
}
