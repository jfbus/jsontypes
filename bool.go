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

func (b *NullBool) Set(val bool) {
	b.Present = true
	b.Val = val
}

func (b *NullBool) UnmarshalJSON(buf []byte) error {
	b.Present = true
	if buf[0] == 'n' {
		b.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &b.Val)
	}
}

func (b *NullBool) MarshalJSON() ([]byte, error) {
	if !b.Present || b.Null {
		return []byte("null"), nil
	}
	return json.Marshal(b.Val)
}

func (b *NullBool) Scan(value interface{}) error {
	if value == nil {
		b.Val, b.Null = false, true
		return nil
	}
	b.Null = false
	return convertAssign(&b.Val, value)
}

func (b NullBool) Value() (driver.Value, error) {
	if b.Null {
		return nil, nil
	}
	return b.Val, nil
}

func (b NullBool) WillUpdate() bool {
	return b.Present
}

type Bool struct {
	Val     bool
	Present bool
}

func (b *Bool) Set(val bool) {
	b.Present = true
	b.Val = val
}

func (b *Bool) UnmarshalJSON(buf []byte) error {
	b.Present = true
	return json.Unmarshal(buf, &b.Val)
}

func (b *Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Val)
}

func (b *Bool) Scan(value interface{}) error {
	if value == nil {
		b.Val = false
		return nil
	}
	return convertAssign(&b.Val, value)
}

func (b Bool) Value() (driver.Value, error) {
	return b.Val, nil
}

func (b Bool) WillUpdate() bool {
	return b.Present
}

type RONullBool struct {
	Val     bool
	Null    bool
	Present bool
}

func (b *RONullBool) Set(val bool) {
	b.Present = true
	b.Val = val
}

func (b *RONullBool) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (b *RONullBool) MarshalJSON() ([]byte, error) {
	if !b.Present || b.Null {
		return []byte("null"), nil
	}
	return json.Marshal(b.Val)
}

func (b *RONullBool) Scan(value interface{}) error {
	if value == nil {
		b.Val, b.Null = false, true
		return nil
	}
	b.Null = false
	return convertAssign(&b.Val, value)
}

func (b RONullBool) Value() (driver.Value, error) {
	if b.Null {
		return nil, nil
	}
	return b.Val, nil
}

func (b RONullBool) WillUpdate() bool {
	return false
}

type ROBool struct {
	Val     bool
	Present bool
}

func (b *ROBool) Set(val bool) {
	b.Present = true
	b.Val = val
}

func (b *ROBool) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (b *ROBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Val)
}

func (b *ROBool) Scan(value interface{}) error {
	if value == nil {
		b.Val = false
		return nil
	}
	return convertAssign(&b.Val, value)
}

func (b ROBool) Value() (driver.Value, error) {
	return b.Val, nil
}

func (b ROBool) WillUpdate() bool {
	return b.Present
}
