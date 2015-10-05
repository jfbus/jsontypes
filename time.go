package jsontypes

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type NullTime struct {
	Val     time.Time
	Null    bool
	Present bool
}

func (n *NullTime) UnmarshalJSON(buf []byte) error {
	n.Present = true
	if buf[0] == 'n' {
		n.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &n.Val)
	}
}

func (n *NullTime) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null || n.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *NullTime) Scan(value interface{}) error {
	return nil
}

func (n NullTime) Value() (driver.Value, error) {
	if n.Null || n.Val.IsZero() {
		return nil, nil
	}
	return n.Val, nil
}

func (n NullTime) WillUpdate() bool {
	return n.Present
}

type Time struct {
	Val     time.Time
	Present bool
}

func (n *Time) UnmarshalJSON(buf []byte) error {
	n.Present = true
	return json.Unmarshal(buf, &n.Val)
}

func (n *Time) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *Time) Scan(value interface{}) error {
	return nil
}

func (n Time) Value() (driver.Value, error) {
	return n.Val, nil
}

func (n Time) WillUpdate() bool {
	return n.Present
}

type RONullTime struct {
	Val     time.Time
	Null    bool
	Present bool
}

func (n *RONullTime) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *RONullTime) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null || n.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *RONullTime) Scan(value interface{}) error {
	return nil
}

func (n RONullTime) Value() (driver.Value, error) {
	if n.Null || n.Val.IsZero() {
		return nil, nil
	}
	return n.Val, nil
}

func (n RONullTime) WillUpdate() bool {
	return false
}

type ROTime struct {
	Val     time.Time
	Present bool
}

func (n *ROTime) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (n *ROTime) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(n.Val)
}

func (n *ROTime) Scan(value interface{}) error {
	return nil
}

func (n ROTime) Value() (driver.Value, error) {
	return n.Val, nil
}

func (n ROTime) WillUpdate() bool {
	return false
}
