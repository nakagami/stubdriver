// Stub (dummy) Driver for Go's database/sql package
// Written by Hajime Nakagami<nakagami@gmail.com>
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
	"database/sql"
	"database/sql/driver"
    "fmt"
    "reflect"
)

type stubDriver struct{}

var globalStubDriver = &stubDriver{}

func (d *stubDriver) Open(dsn string) (driver.Conn, error) {
    fmt.Println("stubDriver.Open()", dsn)

	return nil, nil
}

func init() {
    fmt.Println("init()", reflect.TypeOf(globalStubDriver))

    // for stub test: not call in init() usually
    db, err := globalStubDriver.Open("DataSourceName")
    fmt.Println("db, err=\t", db, err)

	sql.Register("stubdriver", globalStubDriver)
}
