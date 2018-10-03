package main

type config interface {
	Entry(string) (string, error)
	Entries() (map[string]string, error)
}

type tomlConfig struct {
	path string
}

func newTOMLC
