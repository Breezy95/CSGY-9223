package main


import( "fmt"
	"webAppProject/servers"
	"html/template"
	"net/http"
	"strings"
	"context"
	"log"
	"os"
	"time"
	"webAppProject/proto"
	"google.golang.org/grpc")

	type accountInfo struct {
		user string
		pass string
	}

	//var conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//var c =  proto.NewCommsClient(conn)
	//var ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	const(
		address = "localhost:10021"
		defaultUsername = "user"
		defaultPassword = "pass"
		)

	

func runServers() {
    servers.BackendRun()
	fmt.Println("Creating connection to RPC")
	 conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c :=  proto.NewCommsClient(conn)//init client for 
		log.Printf("\n\nvariable c is type: %T\n\n",c)
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

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "AAAAAAAAAAAAAAAAAAAAVVVVVVVVVVVVVVVVVVVVVVVVXXXXXXXXXXXXXXXXXXXXXXXXXXXX") // send data to client side
}


func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("pages/login.gtpl")
		t.Execute(w, nil)
	} else{
		r.ParseForm()
		log.Printf("username submitted: %s \npassword submitted: %s", r.FormValue("username"),r.FormValue("password"))
		var conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		var c =  proto.NewCommsClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		responseRPC,err := c.SendAccountInfo(ctx,&proto.AccountInfo{Username: r.FormValue("username"), Password: r.FormValue("password") })

		if err != nil {
			log.Fatalf("Could not send RPC in login")
		}
		log.Println("Successful message")
		log.Println(responseRPC.GetMessage())
	}
}

	


func main() {
	fmt.Println("Starting from project")
	go runServers()
	
	
	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login", login)
	err:= http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
	runServers()
	
}
