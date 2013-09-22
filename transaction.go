// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
    "fmt"
)

type stubTx struct {}

func (tx *stubTx) Commit() (err error) {
    fmt.Println("stubTx.Commit()")
    return
}

func (tx *stubTx) Rollback() (err error) {
    fmt.Println("stubTx.Rollback()")
    return
}

func NewStubTransaction() *stubTx {
    return new(stubTx)
}
