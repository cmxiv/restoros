package handler

import (
	"fmt"
	"os"
	"restoros/models"
	"text/tabwriter"
)

// ListHandler -
type ListHandler struct {
	command *models.Command
}

// Handle -
func (listHandler *ListHandler) Handle(config *models.Config) *models.Config {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(writer, "Name\t| Version\t| Source\t|")
	for _, pkg := range config.Packages {
		fmt.Fprintln(writer, fmt.Sprintf("%s\t| %s\t| %s\t|", pkg.Name, pkg.Version, pkg.Source))
	}
	writer.Flush()
	return config
}
