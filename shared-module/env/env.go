package env

type env struct {
	*raw
	*computed
}
type Env interface {
	Project() string
	BucketSuffix() string
}

var _ Env = (*env)(nil)

func NewEnv() (*env, error) {
	r, err := newRaw()
	if err != nil {
		return nil, err
	}
	return NewEnvWithRaw(r)
}

func NewEnvWithRaw(raw *raw) (*env, error) {
	c, err := newComputed(raw)
	if err != nil {
		return nil, err
	}
	return &env{
		raw:      raw,
		computed: c,
	}, nil
}
