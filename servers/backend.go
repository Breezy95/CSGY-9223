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
	posts []string
	followedusers []string
	
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
		//accountsSlice := append(accountsSlice, newAcct) 
		log.Printf("Contents of account  %v", newAcct)
		log.Printf("Contents of account array  %v", accountsSlice)
		
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
	ans, err := AccountVerify(req)
	log.Printf("Does this account with password: %s exist: %t", req.GetPassword(), ans)
	if err != nil{
		return &proto.AccountResponse{Message: ans}, errors.New("Account Does Not Exist")
	}
	return &proto.AccountResponse{Message: ans}, nil

}

func AccountVerify(req *proto.AccountInfo) (bool,error) {
	for _,s := range accountsSlice {
		log.Print(s)
		if req.GetUsername() == s.user && req.GetPassword() == s.pass {
			return true, nil
		}
	}
	return false, errors.New("Account Exists Already")
}

func RegisterAccount(ctx context.Context, req *proto.AccountInfo) (*proto.AccountResponse, error) {
	log.Print("Entering account verification method")
	
	ans ,err := AccountVerify(req)
	if err != nil{
		return &proto.AccountResponse{Message: ans}, errors.New("Account Does Not Exist")
	}
	return &proto.AccountResponse{Message: ans}, nil
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




