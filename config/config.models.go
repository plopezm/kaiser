package config

// ConfigurationData Represents kaiser configuration, this data will be taken from kaiser.config.json
type ConfigurationData struct {
	LogFolder string `json:"logfile"`
	Workspace string `json:"workspace"`
}
