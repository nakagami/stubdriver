// Stub (dummy) Driver for Go's database/sql package
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/
package stubdriver

import (
    "testing"
    "database/sql"
)

func TestConnect(t *testing.T) {
    db, err := sql.Open("stubdriver", "dataSourceName")
    if err != nil {
        t.Fatalf("Error connecting: %v", err)
    }
    defer db.Close()

    _, err = db.Exec("SELECT foo, bar")
    if err == nil {
        t.Fatalf("Error Exec(): %v", err)
    }
}
