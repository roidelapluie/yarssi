package main

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type WindowType (uint8)

type Window struct {
	Name   string
	Server string
	Show   bool
}

type Server struct {
	Name        string
	Password    string
	Address     string
	Port        int
	Autoconnect bool
	Ssl         bool
	Autosendcmd string
}

type ConfigMapping struct {
	Windows  []Window
	Servers  []Server
	Hilights []string
	Nickname string
	Realname string
	Username string
}

func (window *Window) IsChannel() bool {
	return strings.HasPrefix(window.Name, "#")
}

func (window *Window) Inc(i int) int {
	return i + 2
}

func ParseInputFile(filename string) (*ConfigMapping, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "could not read file")
	}

	var configMapping ConfigMapping

	err = yaml.Unmarshal(source, &configMapping)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse file")
	}

	return &configMapping, nil
}
