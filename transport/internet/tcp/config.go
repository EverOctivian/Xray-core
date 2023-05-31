package tcp

import (
	"zgjzd.cn/guoqingjun/xray-core/common"
	"zgjzd.cn/guoqingjun/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
