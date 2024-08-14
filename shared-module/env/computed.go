package env

type computed struct {
	bucketSuffix string
}

func newComputed(r *raw) (*computed, error) {
	c := &computed{
		bucketSuffix: r.stackName,
	}
	if err := c.validate(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *computed) validate() error {
	return nil
}

func (c *computed) BucketSuffix() string {
	return c.bucketSuffix
}
