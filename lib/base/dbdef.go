package base

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/slusenti/noskull/lib/utils"
)

const (
	SK_TYPE_BOOL   string = "bool"
	SK_TYPE_STRING string = "string"
	SK_TYPE_INT    string = "int"
	SK_TYPE_FLOAT  string = "float"
	SK_TYPE_DATE   string = "date"
)

type Schema struct {
	sck string
}

func (s *Schema) GetSchema() string {
	return s.sck
}

func (s *Schema) SetSchema(sk string) error {
	var inter interface{}
	json, err := json.Unmarshal([]byte(sk), &inter)
	if err != nil {
		return err
	}
	maproot, b := inter.(map[string]interface{})
	if !b {
		return errors.New("not a root map json schema.")
	}
	check := checkMap(maproot)
	if check {
		s.sck = sk
		return nil
	} else {
		return errors.New("not a valid json schema.")
	}
}

func checkMap(m map[string]interface{}) bool {
	for _, v := range m {
		val, b := v.(string)
		if b {
			if val != SK_TYPE_BOOL &&
				val != SK_TYPE_STRING &&
				val != SK_TYPE_INT &&
				val != SK_TYPE_FLOAT &&
				val != SK_TYPE_DATE {
				return false
			}
		} else {
			val, b = v.([]interface{})
			if b {
				return checkArr(val)
			} else {
				val, b = v.(map[string]interface{})
				if b {
					return checkMap(val)
				} else {
					return false
				}
			}
		}
	}
	return true
}

func checkArr(m []interface{}) bool {
	for _, v := range m {
		val, b := v.(string)
		if b {
			if val != SK_TYPE_BOOL &&
				val != SK_TYPE_STRING &&
				val != SK_TYPE_INT &&
				val != SK_TYPE_FLOAT &&
				val != SK_TYPE_DATE {
				return false
			}
		} else {
			val, b = v.([]interface{})
			if b {

			} else {
				val, b = v.(map[string]interface{})
				if b {
					return checkMap(val)
				} else {
					return false
				}
			}
		}
	}
	return true
}

type DBStruct struct {
	ID      string
	Name    string
	Schema  Schema
	Commits []string
	Entrys  []string
	Users   map[string]string
}

func NewDBStruct(n string, s Schema) *DBStruct {
	return &DBStruct{
		ID:      utils.GenerateID(),
		Name:    n,
		Schema:  s,
		Commits: []string{},
		Entrys:  []string{},
		Users:   map[string]string{},
	}
}

type DBconf []DBStruct

// go binary encoder
func (conf *DBconf) ToGOB64() error {
	var file *os.File
	var err error
	file, err = os.Open("data/definition.nsk")
	if os.IsNotExist(err) {
		file, err = os.Create("data/definition.nsk")
		if err != nil {
			panic(fmt.Sprintln(`failed to create definition.nsk file`, err))
		}
	} else if err != nil {
		panic(fmt.Sprintln(`failed to open definition.nsk file`, err))
	}
	e := gob.NewEncoder(file)
	err = e.Encode(conf)
	if err != nil {
		panic(fmt.Sprintln(`failed gob Encode`, err))
	}
	return nil
}

// go binary decoder
func DBconfFromGOB64() *Config {
	m := DBconf{}
	file, err := os.Open("data/definition.nsk")
	if err != nil {
		panic(fmt.Sprintln(`failed to open definition.nsk file`, err))
	}
	d := gob.NewDecoder(file)
	err = d.Decode(&m)
	if err != nil {
		panic(fmt.Sprintln(`failed gob Decode`, err))
	}
	return &m
}
