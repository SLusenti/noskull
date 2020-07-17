package base

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/slusenti/noskull/lib/utils"
)

type Config struct {
	RootUser        string
	RootPWD         string
	ChunkMaxSize_MB int
	ChunkMaxOBJ     int
	DBBindAddress   string
	DBBindPort      int
}

func NewConfig() *Config {
	return &Config{
		RootUser:        "nskroot",
		RootPWD:         "W" + utils.GenerateID() + "==",
		ChunkMaxOBJ:     300,
		ChunkMaxSize_MB: 10,
		DBBindAddress:   "127.0.0.1",
		DBBindPort:      58888,
	}
}

// go binary encoder
func (conf *Config) ToGOB64() error {
	var file *os.File
	var err error
	file, err = os.Open("data/basesys.nsk")
	if os.IsNotExist(err) {
		file, err = os.Create("data/basesys.nsk")
		if err != nil {
			panic(fmt.Sprintln(`failed to create basesys.nsk file`, err))
		}
	} else if err != nil {
		panic(fmt.Sprintln(`failed to open basesys.nsk file`, err))
	}
	e := gob.NewEncoder(file)
	err = e.Encode(conf)
	if err != nil {
		panic(fmt.Sprintln(`failed gob Encode`, err))
	}
	return nil
}

// go binary decoder
func ConfigFromGOB64() *Config {
	m := Config{}
	file, err := os.Open("data/basesys.nsk")
	if err != nil {
		panic(fmt.Sprintln(`failed to open basesys.nsk file`, err))
	}
	d := gob.NewDecoder(file)
	err = d.Decode(&m)
	if err != nil {
		panic(fmt.Sprintln(`failed gob Decode`, err))
	}
	return &m
}
