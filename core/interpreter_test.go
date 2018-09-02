package core

import (
	"testing"

	"github.com/plopezm/kaiser/core/types"
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
)

type MockedPlugin struct {
	JobName string
}

func (p *MockedPlugin) GetFunctions() map[string]interface{} {
	return map[string]interface{}{
		"testFunction": func() {},
	}
}

func (p *MockedPlugin) GetInstance(context types.JobContext) types.Plugin {
	p.JobName = context.GetJobName()
	return p
}

func Test_RegisterPlugin(t *testing.T) {
	// Given
	plugin := new(MockedPlugin)
	// When
	RegisterPlugin(plugin)
	// Then
	assert := assert.New(t)
	assert.NotNil(plugins)
	assert.Equal(1, len(plugins))
	assert.Equal(plugin, plugins[0])
}

func Test_registerPlugin(t *testing.T) {
	// Given
	plugin := new(MockedPlugin)
	vm := otto.New()

	// When
	registerPlugin(vm, plugin)

	// Then
	assert := assert.New(t)
	assert.NotNil(vm.Get("testFunction"))

}

func Test_addRegistedPlugins(t *testing.T) {
	// Given
	plugin := new(MockedPlugin)
	testContext := &types.JobInstanceContext{
		JobName: "testJob",
	}
	RegisterPlugin(plugin)
	vm := otto.New()

	// When
	addRegistedPlugins(vm, testContext)

	// Then
	assert := assert.New(t)
	assert.NotNil(vm.Get("testFunction"))
	assert.Equal("testJob", plugin.JobName)
}

func TestNewVMWithPlugins(t *testing.T) {
	// Given
	plugin := new(MockedPlugin)
	testContext := &types.JobInstanceContext{
		JobName: "testJob",
	}
	RegisterPlugin(plugin)

	// When
	vm := NewVMWithPlugins(testContext)

	// Then
	assert := assert.New(t)
	assert.NotNil(vm.Get("testFunction"))
	assert.Equal("testJob", plugin.JobName)
}
