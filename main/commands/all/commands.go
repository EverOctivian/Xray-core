package all

import (
	"zgjzd.cn/guoqingjun/xray-core/main/commands/all/api"
	"zgjzd.cn/guoqingjun/xray-core/main/commands/all/tls"
	"zgjzd.cn/guoqingjun/xray-core/main/commands/base"
)

// go:generate go run zgjzd.cn/guoqingjun/xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
