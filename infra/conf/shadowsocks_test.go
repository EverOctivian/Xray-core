package conf_test

import (
	"testing"

	"zgjzd.cn/guoqingjun/xray-core/common/net"
	"zgjzd.cn/guoqingjun/xray-core/common/protocol"
	"zgjzd.cn/guoqingjun/xray-core/common/serial"
	. "zgjzd.cn/guoqingjun/xray-core/infra/conf"
	"zgjzd.cn/guoqingjun/xray-core/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "xray-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				Users: []*protocol.User{{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "xray-password",
					}),
				}},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
