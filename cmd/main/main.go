package main

import (
	"OTPService/Logging"
	"OTPService/RegisterNewToken"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {}

func (s *server) YubikeyRegister(ctx context.Context, req *RegisterNewToken.SetTokenIDReq) (*RegisterNewToken.SetTokenIDResp, error) {
	key_otp := req.GetYubikeyOtp()
	userid := req.GetUserid()
	RegisterNewToken.RegisterNewYubikey(key_otp, int(userid))
	Response  := &RegisterNewToken.SetTokenIDResp{
		TokenCreated:         "True",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	return  Response, nil
}



func  main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		Logging.LogForMSInfo(err)
	}
	s := grpc.NewServer()
	RegisterNewToken.RegisterTokenServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve %v", err)
	}

}