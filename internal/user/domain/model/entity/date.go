package entity

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Date struct{ time.Time }

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("birthday must be YYYY-MM-DD format")
	}
	d.Time = t
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format("2006-01-02") + `"`), nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time.Format("2006-01-02"), nil
}

func (d *Date) Scan(v any) error {
	t, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("cannot scan %T into Date", v)
	}
	d.Time = t
	return nil
}
