package jsontypes

import (
	"database/sql/driver"
	"encoding/json"
	"log"
	"time"
)

type NullTime struct {
	Val     time.Time
	Null    bool
	Present bool
}

func (t *NullTime) Set(val time.Time) {
	t.Present = true
	t.Val = val
}

func (t *NullTime) UnmarshalJSON(buf []byte) error {
	t.Present = true
	if buf[0] == 'n' {
		t.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &t.Val)
	}
}

func (t *NullTime) MarshalJSON() ([]byte, error) {
	if !t.Present || t.Null || t.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Val)
}

func (t *NullTime) Scan(value interface{}) error {
	log.Println("NullTime.Scat is not implemented")
	return nil
}

func (t NullTime) Value() (driver.Value, error) {
	if t.Null || t.Val.IsZero() {
		return nil, nil
	}
	return t.Val, nil
}

func (t NullTime) WillUpdate() bool {
	return t.Present
}

type Time struct {
	Val     time.Time
	Present bool
}

func (t *Time) Set(val time.Time) {
	t.Present = true
	t.Val = val
}

func (t *Time) UnmarshalJSON(buf []byte) error {
	t.Present = true
	return json.Unmarshal(buf, &t.Val)
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if !t.Present || t.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Val)
}

func (t *Time) Scan(value interface{}) error {
	log.Println("Time.Scat is not implemented")
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return t.Val, nil
}

func (t Time) WillUpdate() bool {
	return t.Present
}

type RONullTime struct {
	Val     time.Time
	Null    bool
	Present bool
}

func (t *RONullTime) Set(val time.Time) {
	t.Present = true
	t.Val = val
}

func (t *RONullTime) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (t *RONullTime) MarshalJSON() ([]byte, error) {
	if !t.Present || t.Null || t.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Val)
}

func (t *RONullTime) Scan(value interface{}) error {
	log.Println("RONullTime.Scat is not implemented")
	return nil
}

func (t RONullTime) Value() (driver.Value, error) {
	if t.Null || t.Val.IsZero() {
		return nil, nil
	}
	return t.Val, nil
}

func (t RONullTime) WillUpdate() bool {
	return t.Present
}

type ROTime struct {
	Val     time.Time
	Present bool
}

func (t *ROTime) Set(val time.Time) {
	t.Present = true
	t.Val = val
}

func (t *ROTime) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (t *ROTime) MarshalJSON() ([]byte, error) {
	if !t.Present || t.Val.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Val)
}

func (t *ROTime) Scan(value interface{}) error {
	log.Println("ROTime.Scat is not implemented")
	return nil
}

func (t ROTime) Value() (driver.Value, error) {
	return t.Val, nil
}

func (t ROTime) WillUpdate() bool {
	return t.Present
}
