package config

// UserSettings contains user-specific settings
type UserSettings struct {
	// Timezone is the timezone to use for the application
	Timezone string `yaml:"timezone,omitempty"`
}
