package sources

import "fmt"

func NewBrewSource() brew {
	return brew{}
}

type brew struct {}

func (brew *brew) Find() error {
	return nil
}

func (brew *brew) Name() string {
	return "brew"
}

func (brew *brew) Purge() error {
	return nil
}

func (brew *brew) Remove() error {
	return nil
}

func (brew *brew) Update() error {
	return nil
}

func (brew *brew) Install() error {
	fmt.Println("called install on brew")
	return nil
}