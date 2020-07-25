package main

import (
	"OTPService/Logging"
	"OTPService/RegisterNewToken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	port = ":50051"
)

type server struct {}

	func (s  *server) RegisterNewToken(req RegisterNewToken.SetTokenIDReq, resp RegisterNewToken.RegisterNewTokenServer) error{
	key_otp := req.GetYubikeyOtp()
	userid := req.GetUserid()
    RegisterNewToken.RegisterNewYubikey(key_otp, int(userid), resp)

	return nil
}

func  main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		Logging.LogForMSInfo(err)
	}
	s := grpc.NewServer()
	RegisterNewToken.RegisterNewTokenServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve %v", err)
	}

}