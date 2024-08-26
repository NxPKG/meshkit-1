package patterns

import "github.com/khulnasoft/meshkit/errors"

const (
	ErrInvalidVersionCode = ""
)

func ErrInvalidVersion(err error) error {
	return errors.New(ErrInvalidVersionCode, errors.Alert, []string{"invalid/incompatible semver version"}, []string{err.Error()}, []string{"version history for the content has been tampered outside meshplay"}, []string{"rolllback to one of the previous version"})
}
