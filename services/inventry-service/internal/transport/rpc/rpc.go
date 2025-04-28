package rpc

import (
	"fmt"
	"net"

	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/pkg/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func StartUnsecureRPCServer(hdl *Handler) {
	address := fmt.Sprintf("0.0.0.0:%d", hdl.cfg.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	r := setupUnsecureRouter(hdl)

	zap.S().Infof("grpc listining: %s", address)

	if err = r.Serve(listener); err != nil {
		panic("failed run grpc server! " + err.Error())
	}
}

func setupUnsecureRouter(hdl api.InventryServer) *grpc.Server {
	var s *grpc.Server
	s = grpc.NewServer()

	api.RegisterInventryServer(s, hdl)

	printServiceInterface(s.GetServiceInfo())

	return s
}

func printServiceInterface(m map[string]grpc.ServiceInfo) {
	for service, info := range m {
		fmt.Printf("[gRPC Info] %v:\n", service)

		for _, method := range info.Methods {
			fmt.Printf("[gRPC Info]   %v\n", method.Name)
		}
	}
}
