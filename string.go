package jsontypes

import "encoding/json"

type NullString struct {
	Value   string
	Null    bool
	Present bool
}

func (n *NullString) UnmarshalJSON(buf []byte) error {
	n.Present = true
	if buf[0] == 'n' {
		n.Null = true
		return nil
	} else {
		return json.Unmarshal(buf, &n.Value)
	}
}

func (n *NullString) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

type String struct {
	Value   string
	Present bool
}

func (n *String) UnmarshalJSON(buf []byte) error {
	n.Present = true
	return json.Unmarshal(buf, &n.Value)
}

func (n *String) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Value)
}

type RONullString struct {
	Value   string
	Null    bool
	Present bool
}

func (n *RONullString) UnmarshalJSON(buf []byte) error {
	n.Present = true
	return nil
}

func (n *RONullString) MarshalJSON() ([]byte, error) {
	if !n.Present || n.Null {
		return []byte("null"), nil
	}
	return json.Marshal(n.Value)
}

type ROString struct {
	Value   string
	Present bool
}

func (n *ROString) UnmarshalJSON(buf []byte) error {
	n.Present = true
	return nil
}

func (n *ROString) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Value)
}
