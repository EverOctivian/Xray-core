package conf_test

import (
	"testing"

	"zgjzd.cn/guoqingjun/xray-core/common/net"
	. "zgjzd.cn/guoqingjun/xray-core/infra/conf"
	"zgjzd.cn/guoqingjun/xray-core/proxy/dokodemo"
)

func TestDokodemoConfig(t *testing.T) {
	creator := func() Buildable {
		return new(DokodemoConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"address": "8.8.8.8",
				"port": 53,
				"network": "tcp",
				"timeout": 10,
				"followRedirect": true,
				"userLevel": 1
			}`,
			Parser: loadJSON(creator),
			Output: &dokodemo.Config{
				Address: &net.IPOrDomain{
					Address: &net.IPOrDomain_Ip{
						Ip: []byte{8, 8, 8, 8},
					},
				},
				Port:           53,
				Networks:       []net.Network{net.Network_TCP},
				Timeout:        10,
				FollowRedirect: true,
				UserLevel:      1,
			},
		},
	})
}
