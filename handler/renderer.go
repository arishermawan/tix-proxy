package handler

import (
	"github.com/unrolled/render"
	"net/http"
	"tix-proxy/view"
)

var renderer *render.Render

func init() {
	renderer = render.New(render.Options{})
}

func ResponseRenderer(w http.ResponseWriter, httpCode int, data interface{}, success bool, errResp *view.ErrorResponse) {
	errors := []view.ErrorResponse{}
	if errResp != nil {
		errors = []view.ErrorResponse{*errResp}
	}
	renderer.JSON(w, httpCode, view.Response{
		Data:    data,
		Success: success,
		Errors:  errors,
	})
}

func RawErrorResponseRenderer(w http.ResponseWriter, httpCode int, err error) {
	errResp := &view.Response{
		Errors: []view.ErrorResponse{
			{Message: err.Error(), Severity: "error", Code: "GoTixProxyError"},
		},
	}
	renderer.JSON(w, httpCode, errResp)
}
