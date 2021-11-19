package binding

import (
	"net/http"
)

var (
	Simple = simpleBinding{}
)

type simpleBinding struct{}

func (simpleBinding) Name() string {
	return "simple"
}

func (simpleBinding) Bind(req *http.Request, obj interface{}) error {
	values := req.URL.Query()
	if err := mapForm(obj, values); err != nil {
		return err
	}

	method := req.Method
	if method == "GET" || method == "DELETE" || method == "HEADER" {
		return validate(obj)
	}

	return nil
}

func (simpleBinding) BindExt(req *http.Request, obj interface{}, values map[string][]string) error {
	if err := mapForm(obj, req.URL.Query()); err != nil {
		return err
	}

	if err := mapURI(obj, values); err != nil {
		return err
	}

	method := req.Method
	if method == "GET" || method == "DELETE" || method == "HEADER" {
		return validate(obj)
	}

	return nil
}
