package env

import (
	"errors"
	"os"
)

type raw struct {
	project   string
	stackName string
}

func newRaw() (*raw, error) {
	r := &raw{
		project:   os.Getenv("GCP_PROJECT_ID"),
		stackName: os.Getenv("STACK_NAME"),
	}
	if err := r.validate(); err != nil {
		return nil, err
	}
	return r, nil
}

var (
	ErrGoogleProjectNotSet = errors.New("GCP_PROJECT_ID must be set")
	ErrStackNameNotSet     = errors.New("STACK_NAME must be set")
)

func (r *raw) validate() error {
	var err error
	if r.project == "" {
		err = errors.Join(err, ErrGoogleProjectNotSet)
	}
	if r.stackName == "" {
		err = errors.Join(err, ErrStackNameNotSet)
	}
	return err
}

func (r *raw) Project() string {
	return r.project
}
