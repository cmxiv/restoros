package sourcehandler

import (
	"fmt"
	"restoros/argumentparser"
)

// SourceHandler -
type SourceHandler interface {
	Install()
	Remove()
	Purge()
	Find()
	Update()
}

var (
	apt  SourceHandler
	sdk  SourceHandler
	npm  SourceHandler
	snap SourceHandler
)

// GetSourceHandler -
func GetSourceHandler(source string) (SourceHandler, error) {
	switch source {
	case "apt":
		if apt == nil {
			apt = &AptSourceHandler{}
		}
		return apt, nil
	case "sdk":
		if sdk == nil {
			sdk = &SdkSourceHandler{}
		}
		return sdk, nil
	case "snap":
		if snap == nil {
			snap = &SnapSourceHandler{}
		}
		return snap, nil
	case "npm":
		if npm == nil {
			npm = &NpmSourceHandler{}
		}
		return npm, nil
	default:
		return nil, fmt.Errorf("Invalid source" + argumentparser.UsageMessage)
	}
}
