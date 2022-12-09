package config

// Settings object for defined settings values (Should be overriden at implementation point to suite application)
type Settings struct {
	Hostname string `json:"hostname"`
	Hardfail bool   `json:"containerHardfail"`
}
