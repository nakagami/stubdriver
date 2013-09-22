// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
    "database/sql/driver"
)

type stubConn struct{}

func (sc *stubConn) Begin() (driver.Tx, error) {
    var err error
    return nil, err
}


func (sc *stubConn) Close() (err error) {
    return
}

func (sc *stubConn) Prepare(query string) (driver.Stmt, error) {
    var err error
    return nil, err
}

func (sc *stubConn) Exec(query string, args []driver.Value) (driver.Result, error) {
    var err error
    return nil, err
}

func (sc *stubConn) Query(query string, args []driver.Value) (driver.Rows, error) {
    var err error
    return nil, err
}
