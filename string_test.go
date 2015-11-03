package jsontypes

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalString(t *testing.T) {

	type embed struct {
		NullString   NullString
		String       String
		RONullString RONullString
		ROString     ROString
	}

	type testcase struct {
		json     string
		obj      interface{}
		expected interface{}
	}

	utc := []testcase{

		{`null`, &NullString{}, &NullString{Present: true, Null: true}},
		{`""`, &NullString{}, &NullString{Present: true}},
		{`"foo"`, &NullString{}, &NullString{Present: true, Value: "foo"}},

		{`{}`, &embed{}, &embed{}},
		{`{"NullString": null}`, &embed{}, &embed{NullString: NullString{Present: true, Null: true}}},
		{`{"NullString": ""}`, &embed{}, &embed{NullString: NullString{Present: true}}},
		{`{"NullString": "foo"}`, &embed{}, &embed{NullString: NullString{Present: true, Value: "foo"}}},
		{`{"String": null}`, &embed{}, &embed{String: String{Present: true}}},
		{`{"String": ""}`, &embed{}, &embed{String: String{Present: true}}},
		{`{"String": "foo"}`, &embed{}, &embed{String: String{Present: true, Value: "foo"}}},
		{`{"RONullString": null}`, &embed{}, &embed{RONullString: RONullString{Present: true}}},
		{`{"RONullString": ""}`, &embed{}, &embed{RONullString: RONullString{Present: true}}},
		{`{"RONullString": "foo"}`, &embed{}, &embed{RONullString: RONullString{Present: true}}},
		{`{"ROString": null}`, &embed{}, &embed{}},
		{`{"ROString": ""}`, &embed{}, &embed{}},
		{`{"ROString": "foo"}`, &embed{}, &embed{}},
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
