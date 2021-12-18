package servers

import(
	"errors"
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
		log.Printf("Contents of accounts array %v",accountsSlice)
		return &proto.AccountResponse{Message: true},nil
}


func (serv *server) SendPost(ctx context.Context, req *proto.PostInfo) (*proto.PostReply, error) {
	log.Printf("Entering send post")
	log.Printf(`Received RPC "%s" , "%s", "%s"`, req.GetPost(),req.GetAuthor(),req.GetDate())

	return  &proto.PostReply{Message: true} , nil

}


func (serv *server) DoesAccountExist(ctx context.Context, req *proto.AccountInfo) (*proto.AccountResponse, error) {
	log.Print("Entering Does Acoount Exist")
	log.Printf("Received RPC for account info")
	for i,s := range accountsSlice {
		log.Printf("iteration: %d",i)
		if req.GetUsername() == s.user && req.GetPassword() == s.pass {
			return &proto.AccountResponse{Message: true}, errors.New("Account Does not Exist")
		}
	}
	return &proto.AccountResponse{Message: true},nil
	


}




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




