package Conf

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

type TomlInfo struct {
	Title string
	Mail  MailInfo
}

type MailInfo struct {
	Host     string
	Port     uint
	Mailer   string `toml:"Mailer"`
	Password string
	UseTls   bool
	From     string
}

var (
	conf *TomlInfo
	once sync.Once
)

func GetMailConf(path string) *TomlInfo {
	once.Do(func() {
		path, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(path, &conf); err != nil {
			panic(err)
		}
	})
	return conf
}
