// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
    "database/sql/driver"
    )

type stubStmt struct {
}

func (stmt *stubStmt) Close() (err error) {
    return
}

func (stmt *stubStmt) NumInput() int {
    return 0
}

func (stmt *stubStmt) Exec(args []driver.Value) (driver.Result, error) {
    var err error
    return nil, err
}

func (stmt *stubStmt) Query(args []driver.Value) (driver.Rows, error) {
    var err error
    return nil, err
}
