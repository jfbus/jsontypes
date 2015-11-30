package jsontypes

import "errors"

var (
	ErrReadOnlyValue = errors.New("Read-only value")
)

type Missinger interface {
	Missing() bool
}
