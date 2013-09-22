// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

type stubTx struct {
}

func (tx *stubTx) Commit() (err error) {
    return
}

func (tx *stubTx) Rollback() (err error) {
    return
}
