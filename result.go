// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

type stubResult struct {
	affectedRows int64
	insertId     int64
}

func (res *stubResult) LastInsertId() (int64, error) {
	return res.insertId, nil
}

func (res *stubResult) RowsAffected() (int64, error) {
	return res.affectedRows, nil
}

func NewStubResult() *stubResult {
	return new(stubResult)
}
