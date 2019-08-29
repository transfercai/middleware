package service

import (
	"net/http"
	"transfercai/middleware/error"
)

func Test(r *http.Request) (interface{}, *error.SError) {
	err := r.ParseForm()
	if err != nil {
		return nil, &error.SError{
			Message:   "failed to parse form",
			ErrorCode: -1,
		}
	}
	a := r.Form.Get("a")
	return map[string]interface{}{
		"a": a,
	}, nil
}
