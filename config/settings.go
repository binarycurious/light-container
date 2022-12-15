package config

// Settings object for defined and custom settings values
type Settings struct {
	Environment string `json:"environment"`
	Hostname    string `json:"hostname"`
	Hardfail    bool   `json:"containerHardfail"`
	appended    map[string]interface{}
}

// SettingsProvider -
type SettingsProvider interface {
	GetObject() Settings
	Put(key string, value interface{})
	Get(key *string) interface{}
}

// GetObject - returns the underlying Settings struct
func (s *Settings) GetObject() Settings {
	return *s
}

// Put - add or override appended settings on the global settings object
func (s *Settings) Put(key string, value interface{}) {
	s.appended[key] = value
}

// Get - get a setting value
func (s *Settings) Get(key *string) interface{} {
	return s.appended[*key]
}
