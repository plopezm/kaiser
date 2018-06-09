package http

import (
	"io/ioutil"
	gohttp "net/http"
	"time"

	"github.com/robertkrimen/otto"
)

// New Used by Kaiser, returns new functionality for Kaiser
func New() (functions map[string]interface{}) {
	functions = make(map[string]interface{})
	functions["http"] = map[string]interface{}{
		"get": Get,
	}
	return functions
}

var netClient = &gohttp.Client{
	Timeout: time.Second * 10,
}

// Response The structure of an Http response
type Response struct {
	StatusCode int
	Body       string
	Headers    map[string][]string
}

// Get Performs an http get request
func Get(call otto.FunctionCall) otto.Value {
	url, err := call.Argument(0).ToString()
	resp, err := netClient.Get(url)
	if err != nil {
		res, _ := call.Otto.ToValue(err.Error())
		return res
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res, _ := call.Otto.ToValue(err.Error())
		return res
	}

	var jsResponse = &Response{
		StatusCode: resp.StatusCode,
		Body:       string(body),
		Headers:    resp.Header,
	}

	valueToReturn, err := call.Otto.ToValue(jsResponse)
	if err != nil {
		res, _ := call.Otto.ToValue(err.Error())
		return res
	}
	return valueToReturn
}
