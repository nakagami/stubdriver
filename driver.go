// Stub (dummy) Driver for Go's database/sql package
// Public Domain http://creativecommons.org/publicdomain/zero/1.0/

package stubdriver

import (
	"database/sql"
	"database/sql/driver"
    "fmt"
)

type stubDriver struct{}

func (d *stubDriver) Open(dsn string) (driver.Conn, error) {
    fmt.Println("stubDriver.Open()")
    fmt.Println(dsn)

	return nil, nil
}

func init() {
	sql.Register("stubdriver", &stubDriver{})
}
