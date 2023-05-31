package conf

import (
	"strings"

	"zgjzd.cn/guoqingjun/xray-core/app/commander"
	loggerservice "zgjzd.cn/guoqingjun/xray-core/app/log/command"
	observatoryservice "zgjzd.cn/guoqingjun/xray-core/app/observatory/command"
	handlerservice "zgjzd.cn/guoqingjun/xray-core/app/proxyman/command"
	statsservice "zgjzd.cn/guoqingjun/xray-core/app/stats/command"
	"zgjzd.cn/guoqingjun/xray-core/common/serial"
)

type APIConfig struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}

func (c *APIConfig) Build() (*commander.Config, error) {
	if c.Tag == "" {
		return nil, newError("API tag can't be empty.")
	}

	services := make([]*serial.TypedMessage, 0, 16)
	for _, s := range c.Services {
		switch strings.ToLower(s) {
		case "reflectionservice":
			services = append(services, serial.ToTypedMessage(&commander.ReflectionConfig{}))
		case "handlerservice":
			services = append(services, serial.ToTypedMessage(&handlerservice.Config{}))
		case "loggerservice":
			services = append(services, serial.ToTypedMessage(&loggerservice.Config{}))
		case "statsservice":
			services = append(services, serial.ToTypedMessage(&statsservice.Config{}))
		case "observatoryservice":
			services = append(services, serial.ToTypedMessage(&observatoryservice.Config{}))
		}
	}

	return &commander.Config{
		Tag:     c.Tag,
		Service: services,
	}, nil
}
