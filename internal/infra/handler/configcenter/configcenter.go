package configcenter

import (
	"github.com/goxtools/watcher"
	"github.com/jinfeijie/pangu/internal/constant"
	"github.com/jinfeijie/pangu/internal/dto"
	"github.com/jinfeijie/pangu/internal/infra/handler"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

const configName = "configCenter"

func init() {
	handler.Register(configName, newConfigCenter(), constant.GroupInit)
}

type configCenter struct{}

func newConfigCenter() *configCenter {
	return &configCenter{}
}

func (h *configCenter) Serve() {

	allData, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	var config dto.FrameworkConfig
	err = yaml.Unmarshal(allData, &config)
	if err != nil {
		panic(err)
	}

	conf = config

	tt := time.Second * 10
	t := time.NewTicker(tt)
	go watcher.NewWatcher(10, time.Second, time.Second*10).On(func(args ...interface{}) {
		for {
			select {
			case <-t.C:
				allData, err := os.ReadFile("config.yaml")
				if err != nil {
					panic(err)
				}
				var config dto.FrameworkConfig
				err = yaml.Unmarshal(allData, &config)
				if err != nil {
					panic(err)
				}

				conf = config
				t.Reset(tt)
			default:
				time.Sleep(time.Second)
			}
		}
	})
}

func (h *configCenter) Name() string {
	return configName
}

var conf dto.FrameworkConfig

func GetConfig() dto.FrameworkConfig {
	return conf
}
