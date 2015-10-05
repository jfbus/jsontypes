package jsontypes

import (
	"database/sql/driver"
	"encoding/json"
)

type NullString struct {
	Val     string
	Null    bool
	Present bool
}

func (n *NullString) UnmarshalJSON(buf []byte) error {
	n.Present = true
	if buf[0] == 'n' {
		n.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &n.Val)
	}
}

func (n *NullString) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *NullString) Scan(value interface{}) error {
	return nil
}

func (n NullString) Value() (driver.Value, error) {
	if n.Null {
		return nil, nil
	}
	return n.Val, nil
}

func (n NullString) WillUpdate() bool {
	return n.Present
}

type String struct {
	Val     string
	Present bool
}

func (n *String) UnmarshalJSON(buf []byte) error {
	n.Present = true
	return json.Unmarshal(buf, &n.Val)
}

func (n *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Val)
}

func (n *String) Scan(value interface{}) error {
	return nil
}

func (n String) Value() (driver.Value, error) {
	if !n.Present {
		return nil, nil
	}
	return n.Val, nil
}

func (n String) WillUpdate() bool {
	return n.Present
}

type RONullString struct {
	Val     string
	Null    bool
	Present bool
}

func (n *RONullString) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *RONullString) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *RONullString) Scan(value interface{}) error {
	return nil
}

func (n RONullString) Value() (driver.Value, error) {
	if !n.Present {
		return nil, nil
	}
	return n.Val, nil
}

func (n RONullString) WillUpdate() bool {
	return false
}

type ROString struct {
	Val     string
	Present bool
}

func (n *ROString) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *ROString) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Val)
}
func (n *ROString) Scan(value interface{}) error {
	return nil
}

func (n ROString) Value() (driver.Value, error) {
	if !n.Present {
		return nil, nil
	}
	return n.Val, nil
}

func (n ROString) WillUpdate() bool {
	return false
}
