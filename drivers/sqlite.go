package drivers

import (
	"database/sql/driver"

	"github.com/pkg/errors"
	"modernc.org/sqlite"
)

type SqliteDriver struct {
	*sqlite.Driver
}

func (d SqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		conn.Close()
		return nil, errors.Wrap(err, "failed to enable enable foreign keys")
	}
	return conn, nil
}
