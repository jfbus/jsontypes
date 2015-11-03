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

func (s *NullString) Set(value string) {
	s.Val = value
	s.Present = true
}

func (s *NullString) UnmarshalJSON(buf []byte) error {
	s.Present = true
	if buf[0] == 'n' {
		s.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &s.Val)
	}
}

func (s *NullString) MarshalJSON() ([]byte, error) {
	if !s.Present || s.Null {
		return []byte("null"), nil
	}
	return json.Marshal(s.Val)
}

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		s.Val, s.Null = "", true
		return nil
	}
	s.Null = false
	return convertAssign(&s.Val, value)
}

func (s NullString) Value() (driver.Value, error) {
	if s.Null {
		return nil, nil
	}
	return s.Val, nil
}

func (s NullString) WillUpdate() bool {
	return s.Present
}

type String struct {
	Val     string
	Present bool
}

func (s *String) Set(value string) {
	s.Val = value
	s.Present = true
}

func (s *String) UnmarshalJSON(buf []byte) error {
	s.Present = true
	return json.Unmarshal(buf, &s.Val)
}

func (s *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Val)
}

func (s *String) Scan(value interface{}) error {
	if value == nil {
		s.Val = ""
		return nil
	}
	return convertAssign(&s.Val, value)
}

func (s String) Value() (driver.Value, error) {
	if !s.Present {
		return nil, nil
	}
	return s.Val, nil
}

func (s String) WillUpdate() bool {
	return s.Present
}

type RONullString struct {
	Val     string
	Null    bool
	Present bool
}

func (s *RONullString) Set(value string) {
	s.Val = value
	s.Present = true
}

func (s *RONullString) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (s *RONullString) MarshalJSON() ([]byte, error) {
	if s.Null {
		return []byte("null"), nil
	}
	return json.Marshal(s.Val)
}

func (s *RONullString) Scan(value interface{}) error {
	if value == nil {
		s.Val, s.Null = "", true
		return nil
	}
	s.Null = false
	return convertAssign(&s.Val, value)
}

func (s RONullString) Value() (driver.Value, error) {
	if s.Null {
		return nil, nil
	}
	return s.Val, nil
}

func (s RONullString) WillUpdate() bool {
	return s.Present
}

type ROString struct {
	Val     string
	Present bool
}

func (s *ROString) Set(value string) {
	s.Val = value
	s.Present = true
}

func (s *ROString) UnmarshalJSON(buf []byte) error {
	return ErrReadOnlyValue
}

func (s *ROString) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Val)
}
func (s *ROString) Scan(value interface{}) error {
	if value == nil {
		s.Val = ""
		return nil
	}
	return convertAssign(&s.Val, value)
}

func (s ROString) Value() (driver.Value, error) {
	return s.Val, nil
}

func (s ROString) WillUpdate() bool {
	return s.Present
}
