package sources

import (
	"os/exec"
)

func NewBrewSource() brew {
	return brew{}
}

type brew struct {
	packageName    string
	packageVersion string
}

func (brew *brew) Find() error {
	return exec.Command("brew", "search", brew.packageName).Run()
}

func (brew *brew) Name() string {
	return "brew"
}

func (brew *brew) Purge() error {
	return brew.Remove()
}

func (brew *brew) Remove() error {
	return exec.Command("brew", "remove", brew.packageName).Run()
}

func (brew *brew) Update() error {
	return nil
}

func (brew *brew) Install() error {
	return exec.Command("brew", "install", brew.packageName).Run()
}

func (brew *brew) SetPackage(name string, version string) {
	brew.packageName = name
	brew.packageVersion = version
}
