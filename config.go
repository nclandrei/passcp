package main

const (
	jsonConfigType configType = "json"
)

type configType string

type configurator func(string) (map[string]string, error)

type jsonConfig func(string) (map[string]string, error)
type tomlConfig func(string) (map[string]string, error)
type onePasswordConfig func(string) (map[string]string, error)
type lastPassConfig func(string) (map[string]string, error)
