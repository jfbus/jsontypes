package jsontypes

import (
	"database/sql/driver"
	"encoding/json"
)

type NullBool struct {
	Val     bool
	Null    bool
	Present bool
}

func (n *NullBool) UnmarshalJSON(buf []byte) error {
	n.Present = true
	if buf[0] == 'n' {
		n.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &n.Val)
	}
}

func (n *NullBool) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *NullBool) Scan(value interface{}) error {
	return nil
}

func (n NullBool) Value() (driver.Value, error) {
	if n.Null {
		return nil, nil
	}
	return n.Val, nil
}

func (n NullBool) WillUpdate() bool {
	return n.Present
}

type Bool struct {
	Val     bool
	Present bool
}

func (n *Bool) UnmarshalJSON(buf []byte) error {
	n.Present = true
	return json.Unmarshal(buf, &n.Val)
}

func (n *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Val)
}

func (n *Bool) Scan(value interface{}) error {
	return nil
}

func (n Bool) Value() (driver.Value, error) {
	return n.Val, nil
}

func (n Bool) WillUpdate() bool {
	return n.Present
}

type RONullBool struct {
	Val     Bool
	Null    bool
	Present bool
}

func (n *RONullBool) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *RONullBool) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *RONullBool) Scan(value interface{}) error {
	return nil
}

func (n RONullBool) Value() (driver.Value, error) {
	if n.Null {
		return nil, nil
	}
	return n.Val, nil
}

func (n RONullBool) WillUpdate() bool {
	return false
}

type ROBool struct {
	Val     bool
	Present bool
}

func (n *ROBool) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *ROBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Val)
}

func (n *ROBool) Scan(value interface{}) error {
	return nil
}

func (n ROBool) Value() (driver.Value, error) {
	return n.Val, nil
}

func (n ROBool) WillUpdate() bool {
	return false
}
