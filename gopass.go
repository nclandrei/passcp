package main

import (
	"os"
	"strings"
	"encoding/json"
	"io"

	"github.com/BurntSushi/toml"
)

const (
	jsonConfigType        configType = "json"
	tomlConfigType        configType = "toml"
	onePasswordConfigType configType = "op"
	lastPassConfigType    configType = "lp"
)

type configType string

type Passworder interface {
	Password(string) (string, error)
}

type PassCp struct {
	ConfigFilePath string
	Parsers []Passworder
}

func NewPassCp(configFilePath string) *PassCp {
	fi, err := os.Stat(configFilePath)
	if err != nil {

	}
}

//testing this

func (pc *PassCp) AddPass(v string) error {
	return nil
}

func parseTOML(r io.Reader) (map[string]string, error) {
	values := map[string]string{}
	metaData, err := toml.DecodeReader(r, &values)
	if err != nil {
		return nil, err
	}
	keys := metaData.Keys()
	for _, key := range keys {
		if strings.A(key.String())
	}
}
