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

var accountsSlice = make([]accountInfo,0)

type server struct {
	proto.UnimplementedCommsServer
}

func (serv *server) SendAccountInfo(ctx context.Context, req *proto.AccountInfo) (*proto.AccountResponse,error) {
		log.Printf("Received RPC %v", req)
		//account := accountInfo{user: ,pass: }
		//accountsSlice = append(, *req)kkkkkkkk
		newAcct := accountInfo{user: req.GetUsername(), pass: req.GetPassword()}
		accountsSlice := append(accountsSlice, newAcct) 
		log.Println(accountsSlice)
		return &proto.AccountResponse{Message: true},nil
}

/*func (serv *server) SendPost(ctx context.Context, req *proto.AccountInfo) (*proto.AccountResponse,nil) {
	log.Printf("Received RPC %v", req.GetUsername())

	return  &proto.AccountResponse{Message: true} , nil

}

*/


func BackendRun(){
	lis, err :=  net.Listen("tcp", port)
	if err != nil {
			log.Fatalf("Failure to listen: %v", err)
	}
	s:= grpc.NewServer()
	 
	proto.RegisterCommsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	}




