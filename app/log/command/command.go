package command

//go:generate go run zgjzd.cn/guoqingjun/xray-core/common/errors/errorgen

import (
	"context"

	grpc "google.golang.org/grpc"
	"zgjzd.cn/guoqingjun/xray-core/app/log"
	"zgjzd.cn/guoqingjun/xray-core/common"
	"zgjzd.cn/guoqingjun/xray-core/core"
)

type LoggerServer struct {
	V *core.Instance
}

// RestartLogger implements LoggerService.
func (s *LoggerServer) RestartLogger(ctx context.Context, request *RestartLoggerRequest) (*RestartLoggerResponse, error) {
	logger := s.V.GetFeature((*log.Instance)(nil))
	if logger == nil {
		return nil, newError("unable to get logger instance")
	}
	if err := logger.Close(); err != nil {
		return nil, newError("failed to close logger").Base(err)
	}
	if err := logger.Start(); err != nil {
		return nil, newError("failed to start logger").Base(err)
	}
	return &RestartLoggerResponse{}, nil
}

func (s *LoggerServer) mustEmbedUnimplementedLoggerServiceServer() {}

type service struct {
	v *core.Instance
}

func (s *service) Register(server *grpc.Server) {
	ls := &LoggerServer{
		V: s.v,
	}
	RegisterLoggerServiceServer(server, ls)

	// For compatibility purposes
	vCoreDesc := LoggerService_ServiceDesc
	vCoreDesc.ServiceName = "v2ray.core.app.log.command.LoggerService"
	server.RegisterService(&vCoreDesc, ls)
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, cfg interface{}) (interface{}, error) {
		s := core.MustFromContext(ctx)
		return &service{v: s}, nil
	}))
}
