package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullString sql.NullString

// MarshalJSON customizes JSON output for NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil) // Returns "null" in JSON
}

type NullTime sql.NullTime

func (ns NullTime) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.Time)
	}
	return json.Marshal(nil) // Returns "null" in JSON
}

func (n *NullTime) Scan(value interface{}) error {
	return (*sql.NullTime)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}
