# jsontypes

jsontypes enables you to decode json to a struct and be able to :

* know if a value was set to null or to a zero value,
* know if the corresponding attribute was present in the json object, or not.

jsontypes provides :

* standard, non nullable types : null maps to the zero value (0, "", ...)
* nullable types (null is different from zero)
* read-only types (unmarshal throws an error, marshal marshals the struct value)

Examples :

```go
type Foo struct {
	Bar jsontypes.NullString
	Baz jsontypes.ROString
}
foo := &Foo{}
// returns an error
err := json.Unmarshal([]byte(`{"Baz": "baz"}`), foo)
// foo.Bar.Present is set to false
err := json.Unmarshal([]byte(`{}`), foo)
```

TODO :
* add more types: Int, Float, time.Time (with null output for zero values)
* add String(), etc. helpers
* add sql support (mapping to sql.NullString, etc)
* add bson support