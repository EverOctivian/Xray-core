package command_test

import (
	"context"
	"testing"

	"zgjzd.cn/guoqingjun/xray-core/app/dispatcher"
	"zgjzd.cn/guoqingjun/xray-core/app/log"
	. "zgjzd.cn/guoqingjun/xray-core/app/log/command"
	"zgjzd.cn/guoqingjun/xray-core/app/proxyman"
	_ "zgjzd.cn/guoqingjun/xray-core/app/proxyman/inbound"
	_ "zgjzd.cn/guoqingjun/xray-core/app/proxyman/outbound"
	"zgjzd.cn/guoqingjun/xray-core/common"
	"zgjzd.cn/guoqingjun/xray-core/common/serial"
	"zgjzd.cn/guoqingjun/xray-core/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
