//go:build freebsd
// +build freebsd

package tcp

import (
	"zgjzd.cn/guoqingjun/xray-core/common/net"
	"zgjzd.cn/guoqingjun/xray-core/transport/internet"
	"zgjzd.cn/guoqingjun/xray-core/transport/internet/stat"
)

// GetOriginalDestination from tcp conn
func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	la := conn.LocalAddr()
	ra := conn.RemoteAddr()
	ip, port, err := internet.OriginalDst(la, ra)
	if err != nil {
		return net.Destination{}, newError("failed to get destination").Base(err)
	}
	dest := net.TCPDestination(net.IPAddress(ip), net.Port(port))
	if !dest.IsValid() {
		return net.Destination{}, newError("failed to parse destination.")
	}
	return dest, nil
}
