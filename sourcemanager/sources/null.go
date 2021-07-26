package sources

import "fmt"

func NewNullSource() null {
	return null{}
}

type null struct{}

func (null *null) Find() error {
	return fmt.Errorf("nothing can be found from null source")
}

func (null *null) Name() string {
	return "null"
}

func (null *null) Purge() error {
	return fmt.Errorf("nothing to purge from null source")
}

func (null *null) Remove() error {
	return fmt.Errorf("nothing can be removed from null source")
}

func (null *null) Update() error {
	return fmt.Errorf("nothing can be updated from null source")
}

func (null *null) Install() error {
	return fmt.Errorf("no sources found to install the provided package")
}

func (null *null) SetPackage(name string, version string) {}
