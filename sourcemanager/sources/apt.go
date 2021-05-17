package sources

import "fmt"

func NewAptSource() apt {
	return apt{}
}

type apt struct{}

func (apt *apt) Find() error {
	return nil
}

func (apt *apt) Name() string {
	return "apt"
}

func (apt *apt) Purge() error {
	return nil
}

func (apt *apt) Remove() error {
	return nil
}

func (apt *apt) Update() error {
	return nil
}

func (apt *apt) Install() error {
	fmt.Println("called install on apt")
	return nil
}
