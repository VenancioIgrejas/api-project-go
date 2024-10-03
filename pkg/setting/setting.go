package setting

import (
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Setting struct {
	Server ServerSetting `toml:"server"`
	Logger LogSetting    `toml:"logger"`
	Main   MainSetting   `toml:"main"`
}

type ServerSetting struct {
	Host         string        `toml:"host"`
	Port         int32         `toml:"port"`
	RunMode      string        `toml:"run_mode"`
	ReadTimeout  time.Duration `toml:"read_timeout"`
	WriteTimeout time.Duration `toml:"write_timeout"`
}

type LogSetting struct {
	SavePath   string `toml:"save_path"`
	SaveName   string `toml:"save_name"`
	FileExt    string `toml:"file_ext"`
	TimeFormat string `toml:"time_format"`
}

type MainSetting struct {
	RuntimeRootPath string `toml:"runtime_rootpath"`
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
