package setting

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Setting struct {
	Server ServerSetting `toml:"server"`
}

type ServerSetting struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

var AppSetting = &Setting{}

func Setup() {
	var err error
	cfg, err := os.ReadFile("conf/app.toml")

	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	if _, err := toml.Decode(string(cfg), AppSetting); err != nil {
		panic(err)
	}

}
