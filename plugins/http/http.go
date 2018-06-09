package http

import (
	"io/ioutil"
	gohttp "net/http"
	"time"

	"github.com/plopezm/kaiser/core/context"
	"github.com/plopezm/kaiser/plugins"
	"github.com/robertkrimen/otto"
)

var netClient = &gohttp.Client{
	Timeout: time.Second * 10,
}

// Response The structure of an Http response
type Response struct {
	StatusCode int
	Body       string
	Headers    map[string][]string
}

type HttpPlugin struct {
	context context.JobContext
}

// New Used by Kaiser, returns new functionality for Kaiser
func New(context context.JobContext) plugins.KaiserPlugin {
	plugin := new(HttpPlugin)
	plugin.context = context
	return plugin
}

// GetFunctions returns the functions to be registered in the VM
func (plugin *HttpPlugin) GetFunctions() map[string]interface{} {
	functions := make(map[string]interface{})
	functions["http"] = map[string]interface{}{
		"get": plugin.Get,
	}
	return functions
}

// Get Performs an http get request
func (plugin *HttpPlugin) Get(call otto.FunctionCall) otto.Value {
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
