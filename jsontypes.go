package jsontypes

import "errors"

var (
	ErrReadOnlyValue = errors.New("Read-only value")
)

type WillUpdater interface {
	WillUpdate() bool
}
