// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
	"database/sql/driver"
	"fmt"
	"io"
)

type stubRows struct {
	curNum int
}

func NewStubRows() *stubRows {
	return new(stubRows)
}

func (rows *stubRows) Columns() []string {
	fmt.Println("stubRows.Columns()")
	return []string{"Column1", "Column2"}
}

func (rows *stubRows) Close() (er error) {
	fmt.Println("stubRows.Close()")
	return
}

func (rows *stubRows) Next(dest []driver.Value) (err error) {
	fmt.Println("stubRows.Next()")
	rows.curNum++
	if rows.curNum < 5 {
		dest[0] = rows.curNum
		dest[1] = "ABC"
	} else {
		err = io.EOF
	}

	return
}
