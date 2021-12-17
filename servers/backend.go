package servers

import(
	"context"
	"log"
	"net"
	"webAppProject/proto"
	"google.golang.org/grpc"
)

const(
	port = ":10021"
)



type accountInfo struct {
	user string
	pass string
}

var accounts = make([]accountInfo, 3)
type serverobj struct {
	proto.UnimplementedCommsServer
}

func (serv *serverobj) SendAccountInfo(ctx context.Context, req *proto.AccountInfo) (*proto.AccountResponse) {
		log.Printf("Received RPC %v", req)
		return &proto.AccountResponse{Message: true}
}

func (serv *serverobj) SendPost(ctx context.Context, req *proto.AccountInfo) ()


func BackendRun(){
	lis, err :=  net.Listen("tcp", port)
	if err != nil {
			log.Fatalf("Failure to listen: %v", err)
	}
	s:= grpc.NewServer()
	servstruct := proto.UnimplementedCommsServer{} 
	proto.RegisterCommsServer(s, servstruct)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	}




