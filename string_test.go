package jsontypes

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalString(t *testing.T) {

	type embed struct {
		NullString NullString
		String     String
	}

	type testcase struct {
		json     string
		obj      interface{}
		expected interface{}
	}

	utc := []testcase{

		{`null`, &NullString{}, &NullString{Present: true, Null: true}},
		{`""`, &NullString{}, &NullString{Present: true}},
		{`"foo"`, &NullString{}, &NullString{Present: true, Val: "foo"}},

		{`{}`, &embed{}, &embed{}},
		{`{"NullString": null}`, &embed{}, &embed{NullString: NullString{Present: true, Null: true}}},
		{`{"NullString": ""}`, &embed{}, &embed{NullString: NullString{Present: true}}},
		{`{"NullString": "foo"}`, &embed{}, &embed{NullString: NullString{Present: true, Val: "foo"}}},
		{`{"String": null}`, &embed{}, &embed{String: String{Present: true}}},
		{`{"String": ""}`, &embed{}, &embed{String: String{Present: true}}},
		{`{"String": "foo"}`, &embed{}, &embed{String: String{Present: true, Val: "foo"}}},
	}

	for _, tc := range utc {
		err := json.Unmarshal([]byte(tc.json), tc.obj)
		if err != nil {
			t.Errorf("%s expected to unmarshal to %#v but got error %s", tc.json, tc.expected, err)
		}
		if !reflect.DeepEqual(tc.obj, tc.expected) {
			t.Errorf("%s expected to unmarshal to %#v but got %#v", tc.json, tc.expected, tc.obj)
		}
	}
}

func TestUnmarshalROString(t *testing.T) {

	type embed struct {
		NullString RONullString
		String     ROString
	}

	utc := []string{
		`{"NullString": null}`,
		`{"NullString": ""}`,
		`{"NullString": "foo"}`,
		`{"String": null}`,
		`{"String": ""}`,
		`{"String": "foo"}`,
	}

	for _, tc := range utc {
		obj := embed{}
		err := json.Unmarshal([]byte(tc), &obj)
		if err == nil {
			t.Errorf("Unmarshaling of RO field in %s expected to generate a Read-Only error, nil found", tc)
		}
	}
}
