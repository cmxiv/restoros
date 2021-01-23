package models

// Command -
type Command struct {
	Primary   string
	Secondary string
	Arguments []string
}

// Config -
type Config struct {
	Modified bool
	Sources  []string  `json:"sources"`
	Packages []Package `json:"packages"`
}

// Package -
type Package struct {
	Skip    bool   `json:"skip"`
	Name    string `json:"name"`
	Flags   string `json:"flags"`
	Source  string `json:"source"`
	Version string `json:"version"`
	Command string `json:"command"`
}
