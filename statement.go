// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
	"database/sql/driver"
	"fmt"
)

type stubStmt struct {
}

func (stmt *stubStmt) Close() (err error) {
	fmt.Println("stubStmt.Close()")

	return
}

func (stmt *stubStmt) NumInput() int {
	numInput := -1
	fmt.Println("stubStmt.NumInput()", numInput)

	return numInput
}

func (stmt *stubStmt) Exec(args []driver.Value) (driver.Result, error) {
	var err error
	fmt.Println("stubStmt.Exec()")
	r := NewStubResult()

	return r, err
}

func (stmt *stubStmt) Query(args []driver.Value) (driver.Rows, error) {
	var err error
	fmt.Println("stubStmt.Query()")

	r := NewStubRows()
	return r, err
}

func NewStubStatement() *stubStmt {
	return new(stubStmt)
}
