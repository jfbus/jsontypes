package jsontypes

import (
	"database/sql/driver"
	"encoding/json"
)

type NullInt64 struct {
	Val     int64
	Null    bool
	Present bool
}

func (n *NullInt64) UnmarshalJSON(buf []byte) error {
	n.Present = true
	if buf[0] == 'n' {
		n.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &n.Val)
	}
}

func (n *NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *NullInt64) Scan(value interface{}) error {
	return nil
}

func (n NullInt64) Value() (driver.Value, error) {
	if n.Null {
		return nil, nil
	}
	return n.Val, nil
}

func (n NullInt64) WillUpdate() bool {
	return n.Present
}

type Int64 struct {
	Val     int64
	Present bool
}

func (n *Int64) UnmarshalJSON(buf []byte) error {
	n.Present = true
	return json.Unmarshal(buf, &n.Val)
}

func (n *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Val)
}

func (n *Int64) Scan(value interface{}) error {
	return nil
}

func (n Int64) Value() (driver.Value, error) {
	return n.Val, nil
}

func (n Int64) WillUpdate() bool {
	return n.Present
}

type RONullInt64 struct {
	Val     int64
	Null    bool
	Present bool
}

func (n *RONullInt64) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *RONullInt64) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *RONullInt64) Scan(value interface{}) error {
	return nil
}

func (n RONullInt64) Value() (driver.Value, error) {
	if n.Null {
		return nil, nil
	}
	return n.Val, nil
}

func (n RONullInt64) WillUpdate() bool {
	return false
}

type ROInt64 struct {
	Val     int64
	Present bool
}

func (n *ROInt64) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *ROInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Val)
}

func (n *ROInt64) Scan(value interface{}) error {
	return nil
}

func (n ROInt64) Value() (driver.Value, error) {
	return n.Val, nil
}

func (n ROInt64) WillUpdate() bool {
	return false
}
