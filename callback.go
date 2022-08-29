package checkpoint

// Callbacks should be used only for custom metrics, logging and monitoring purpose.
type Callbacks struct {
	BeforeRequest func(map[string]interface{}, map[string]interface{})
	AfterRequest  func(map[string]interface{}, map[string]interface{})
	OnSuccess     func(map[string]interface{}, map[string]interface{})
	OnFailure     func(map[string]interface{}, map[string]interface{})
}

func (c *Checkpoint) initCallbacks(callbacks *Callbacks) {
	c.Callbacks = &Callbacks{
		BeforeRequest: DefaultAfterRequest,
		AfterRequest:  DefaultAfterRequest,
		OnSuccess:     DefaultOnSuccess,
		OnFailure:     DefaultOnFailure,
	}

	if callbacks.BeforeRequest != nil {
		c.Callbacks.BeforeRequest = callbacks.BeforeRequest
	}
	if callbacks.AfterRequest != nil {
		c.Callbacks.AfterRequest = callbacks.AfterRequest
	}
	if callbacks.OnSuccess != nil {
		c.Callbacks.OnSuccess = callbacks.OnSuccess
	}
	if callbacks.OnFailure != nil {
		c.Callbacks.OnFailure = callbacks.OnFailure
	}
}

func DefaultBeforeRequest(values map[string]interface{}, reqValues map[string]interface{}) {

}

func DefaultAfterRequest(values map[string]interface{}, reqValues map[string]interface{}) {

}

func DefaultOnSuccess(values map[string]interface{}, reqValues map[string]interface{}) {

}

func DefaultOnFailure(values map[string]interface{}, reqValues map[string]interface{}) {

}
