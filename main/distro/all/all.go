package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "zgjzd.cn/guoqingjun/xray-core/app/dispatcher"
	_ "zgjzd.cn/guoqingjun/xray-core/app/proxyman/inbound"
	_ "zgjzd.cn/guoqingjun/xray-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "zgjzd.cn/guoqingjun/xray-core/app/commander"
	_ "zgjzd.cn/guoqingjun/xray-core/app/log/command"
	_ "zgjzd.cn/guoqingjun/xray-core/app/proxyman/command"
	_ "zgjzd.cn/guoqingjun/xray-core/app/stats/command"

	// Developer preview services
	_ "zgjzd.cn/guoqingjun/xray-core/app/observatory/command"

	// Other optional features.
	_ "zgjzd.cn/guoqingjun/xray-core/app/dns"
	_ "zgjzd.cn/guoqingjun/xray-core/app/dns/fakedns"
	_ "zgjzd.cn/guoqingjun/xray-core/app/log"
	_ "zgjzd.cn/guoqingjun/xray-core/app/metrics"
	_ "zgjzd.cn/guoqingjun/xray-core/app/policy"
	_ "zgjzd.cn/guoqingjun/xray-core/app/reverse"
	_ "zgjzd.cn/guoqingjun/xray-core/app/router"
	_ "zgjzd.cn/guoqingjun/xray-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "zgjzd.cn/guoqingjun/xray-core/app/observatory"

	// Inbound and outbound proxies.
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/blackhole"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/dns"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/dokodemo"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/freedom"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/http"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/loopback"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/mtproto"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/shadowsocks"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/socks"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/trojan"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/vless/inbound"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/vless/outbound"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/vmess/inbound"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/vmess/outbound"
	_ "zgjzd.cn/guoqingjun/xray-core/proxy/wireguard"

	// Transports
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/domainsocket"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/grpc"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/http"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/kcp"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/quic"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/reality"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/tcp"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/tls"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/udp"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/websocket"

	// Transport headers
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/headers/http"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/headers/noop"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/headers/srtp"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/headers/tls"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/headers/utp"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/headers/wechat"
	_ "zgjzd.cn/guoqingjun/xray-core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "zgjzd.cn/guoqingjun/xray-core/main/json"
	_ "zgjzd.cn/guoqingjun/xray-core/main/toml"
	_ "zgjzd.cn/guoqingjun/xray-core/main/yaml"

	// Load config from file or http(s)
	_ "zgjzd.cn/guoqingjun/xray-core/main/confloader/external"

	// Commands
	_ "zgjzd.cn/guoqingjun/xray-core/main/commands/all"
)
