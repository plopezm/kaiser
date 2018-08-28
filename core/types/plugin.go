package types

// Plugin This interface should be implemented for every plugin
type Plugin interface {
	// GetFunctions returns the functions to be registered in the VM
	GetFunctions() map[string]interface{}
	// GetInstance Creates a new plugin instance with a context
	GetInstance(context JobContext) Plugin
}
