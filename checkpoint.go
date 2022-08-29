package checkpoint

type Checkpoint struct {
	Callbacks *Callbacks

	Values map[string]interface{}
}

// NewCheckpoint returns a checkpoint struct initialized and ready to run.
func NewCheckpoint(callbacks *Callbacks) *Checkpoint {
	cp := &Checkpoint{
		Values: make(map[string]interface{}),
	}

	cp.initCallbacks(callbacks)

	return cp
}

func (c *Checkpoint) Execute(fn func() (interface{}, error)) (interface{}, error) {
	reqValues := make(map[string]interface{})

	c.Callbacks.BeforeRequest(c.Values, reqValues)
	res, err := fn()
	c.Callbacks.AfterRequest(c.Values, reqValues)

	if err != nil {
		c.Callbacks.OnFailure(c.Values, reqValues)
	} else {
		c.Callbacks.OnSuccess(c.Values, reqValues)
	}

	return res, err
}
