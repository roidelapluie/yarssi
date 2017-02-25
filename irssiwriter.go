package main

import (
	"github.com/pkg/errors"
	"os"
	"text/template"
)

func WriteIrssFile(configMapping ConfigMapping) error {
	tmpl, err := template.New("irssiConfig").Parse(irssiConfigTemplate)
	if err != nil {
		return errors.Wrap(err, "could not parse template")
	}
	err = tmpl.Execute(os.Stdout, configMapping)
	if err != nil {
		return errors.Wrap(err, "could not write template")
	}
	return nil
}
