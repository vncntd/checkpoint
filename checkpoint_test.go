package checkpoint

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCallbacks(t *testing.T) {
	t.Run("Custom working flow", func(t *testing.T) {
		var beforeReq, afterReq, successReq, failureReq int
		callbacks := Callbacks{
			BeforeRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				beforeReq++
			},
			AfterRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				afterReq++
			},
			OnSuccess: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				successReq++
			},
			OnFailure: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				failureReq++
			},
		}

		cp := NewCheckpoint(&callbacks)
		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})

		assert.Equal(t, 1, beforeReq)
		assert.Equal(t, 1, afterReq)
		assert.Equal(t, 1, successReq)
		assert.Equal(t, 0, failureReq)

		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})
		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})
		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})

		assert.Equal(t, 4, beforeReq)
		assert.Equal(t, 4, afterReq)
		assert.Equal(t, 4, successReq)
		assert.Equal(t, 0, failureReq)
	})

	t.Run("Custom failing flow", func(t *testing.T) {
		var beforeReq, afterReq, successReq, failureReq int
		callbacks := Callbacks{
			BeforeRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				beforeReq++
			},
			AfterRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				afterReq++
			},
			OnSuccess: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				successReq++
			},
			OnFailure: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				failureReq++
			},
		}

		cp := NewCheckpoint(&callbacks)
		cp.Execute(func() (interface{}, error) {
			return nil, errors.New("failing")
		})

		assert.Equal(t, 1, beforeReq)
		assert.Equal(t, 1, afterReq)
		assert.Equal(t, 0, successReq)
		assert.Equal(t, 1, failureReq)

		cp.Execute(func() (interface{}, error) {
			return nil, errors.New("failing")
		})
		cp.Execute(func() (interface{}, error) {
			return nil, errors.New("failing")
		})
		cp.Execute(func() (interface{}, error) {
			return nil, errors.New("failing")
		})

		assert.Equal(t, 4, beforeReq)
		assert.Equal(t, 4, afterReq)
		assert.Equal(t, 0, successReq)
		assert.Equal(t, 4, failureReq)
	})

	t.Run("Custom global flow", func(t *testing.T) {
		var beforeReq, afterReq, successReq, failureReq int
		callbacks := Callbacks{
			BeforeRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				beforeReq++
			},
			AfterRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				afterReq++
			},
			OnSuccess: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				successReq++
			},
			OnFailure: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				failureReq++
			},
		}

		cp := NewCheckpoint(&callbacks)
		for i := 0; i < 20; i++ {
			cp.Execute(func() (interface{}, error) {
				return nil, errors.New("failing")
			})
		}
		for i := 0; i < 10; i++ {
			cp.Execute(func() (interface{}, error) {
				return nil, nil
			})
		}

		assert.Equal(t, 30, beforeReq)
		assert.Equal(t, 30, afterReq)
		assert.Equal(t, 10, successReq)
		assert.Equal(t, 20, failureReq)
	})
}

func TestValuesScope(t *testing.T) {
	t.Run("Test working flow", func(t *testing.T) {
		id := uuid.New()

		callbacks := Callbacks{
			BeforeRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				if val, ok := globalValues[id.String()]; !ok {
					globalValues[id.String()] = 1
				} else {
					globalValues[id.String()] = val.(int) + 1
				}
				assert.Empty(t, reqValues)
				reqValues[id.String()] = "Local values"
			},
			AfterRequest: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				assert.NotEmpty(t, reqValues)
				assert.Equal(t, reqValues[id.String()], "Local values")
			},
			OnSuccess: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
				assert.NotEmpty(t, reqValues)
				assert.Equal(t, reqValues[id.String()], "Local values")
			},
			OnFailure: func(globalValues map[string]interface{}, reqValues map[string]interface{}) {
			},
		}

		cp := NewCheckpoint(&callbacks)

		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})
		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})
		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})
		cp.Execute(func() (interface{}, error) {
			return nil, nil
		})

		assert.Equal(t, cp.Values[id.String()], 4)
	})
}
