package main


import( "fmt"
	"webAppProject/servers"
	"html/template"
	"context"
	"log"
	"os"
	"time"
	"webAppProject/proto"
	"google.golang.org/grpc")


	const(
		address = "localhost:10021"
		defaultUsername = "user"
		defaultPassword = "pass"
		)

func main() {
	fmt.Println("Starting from project")
	//go servers.ClientRun()
	 go servers.BackendRun()
	fmt.Println("Creating connection to RPC")
	 conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c :=  proto.NewCommsClient(conn)//init client for 

        // Contact the server and print out its response.
        name := defaultUsername
        if len(os.Args) > 1 {
                name = os.Args[1]
        }
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.SendAccountInfo(ctx, &proto.AccountInfo{Username: name, Password: defaultPassword})//&pb.HelloRequest{Name: name})
		resp, err2 := c.SendPost(ctx, &proto.PostInfo{Post: "Test Post" ,Author: name, Date: "today"})
		r3, err3 := c.DoesAccountExist(ctx, &proto.AccountInfo{Username: name, Password: defaultPassword})
        if err != nil || err2 != nil  || err3 != nil{
                log.Fatalf("could not greet: %v", err)
        }
		//Prints response
        log.Println( r.GetMessage())
		log.Println(resp.GetMessage())
		log.Println(r3.GetMessage())




}
