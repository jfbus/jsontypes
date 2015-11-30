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

func (i *NullInt64) Set(val int64) {
	i.Present = true
	i.Val = val
}

func (i *NullInt64) UnmarshalJSON(buf []byte) error {
	i.Present = true
	if buf[0] == 'n' {
		i.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &i.Val)
	}
}

func (i *NullInt64) MarshalJSON() ([]byte, error) {
	if !i.Present || i.Null {
		return []byte("null"), nil
	}
	return json.Marshal(i.Val)
}

func (i *NullInt64) Scan(value interface{}) error {
	if value == nil {
		i.Val, i.Null = 0, true
		return nil
	}
	i.Null = false
	return convertAssign(&i.Val, value)
}

func (i NullInt64) Value() (driver.Value, error) {
	if i.Null {
		return nil, nil
	}
	return i.Val, nil
}

func (i NullInt64) Missing() bool {
	return !i.Present
}

type Int64 struct {
	Val     int64
	Present bool
}

func (i *Int64) Set(val int64) {
	i.Present = true
	i.Val = val
}

func (i *Int64) UnmarshalJSON(buf []byte) error {
	i.Present = true
	return json.Unmarshal(buf, &i.Val)
}

func (i *Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Val)
}

func (i *Int64) Scan(value interface{}) error {
	if value == nil {
		i.Val = 0
		return nil
	}
	return convertAssign(&i.Val, value)
}

func (i Int64) Value() (driver.Value, error) {
	return i.Val, nil
}

func (i Int64) Missing() bool {
	return !i.Present
}

type RONullInt64 struct {
	Val     int64
	Null    bool
	Present bool
}

func (i *RONullInt64) Set(val int64) {
	i.Present = true
	i.Val = val
}

func (i *RONullInt64) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (i *RONullInt64) MarshalJSON() ([]byte, error) {
	if !i.Present || i.Null {
		return []byte("null"), nil
	}
	return json.Marshal(i.Val)
}

func (i *RONullInt64) Scan(value interface{}) error {
	if value == nil {
		i.Val, i.Null = 0, true
		return nil
	}
	i.Null = false
	return convertAssign(&i.Val, value)
}

func (i RONullInt64) Value() (driver.Value, error) {
	if i.Null {
		return nil, nil
	}
	return i.Val, nil
}

func (i RONullInt64) Missing() bool {
	return !i.Present
}

type ROInt64 struct {
	Val     int64
	Present bool
}

func (i *ROInt64) Set(val int64) {
	i.Present = true
	i.Val = val
}

func (i *ROInt64) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (i *ROInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Val)
}

func (i *ROInt64) Scan(value interface{}) error {
	if value == nil {
		i.Val = 0
		return nil
	}
	return convertAssign(&i.Val, value)
}

func (i ROInt64) Value() (driver.Value, error) {
	return i.Val, nil
}

func (i ROInt64) Missing() bool {
	return !i.Present
}
