package plugins

// KaiserPlugin This interface should be implemented for every plugin
type KaiserPlugin interface {
	// GetFunctions returns the functions to be registered in the VM
	GetFunctions() map[string]interface{}
}
