package configurationmanager

type RestorosConfiguration struct {
	Packages []Package `json:"packages"`
}

type Package struct {
	Skip    bool   `json:"skip"`
	Name    string `json:"name"`
	Flags   string `json:"flags"`
	Source  string `json:"source"`
	Version string `json:"version"`
	Command string `json:"command"`
}
