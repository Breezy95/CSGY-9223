package main


import( 
	_ "embed"
	"fmt"
	"webAppProject/servers"
	"html/template"
	"net/http"
	//"strings"
	"context"
	"log"
	"os"
	"time"
	"webAppProject/proto"
	"google.golang.org/grpc"
	_ "github.com/dgrijalva/jwt-go"

	//"github.com/bradrydzewski/go.auth"
)

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


func mainPage(w http.ResponseWriter, r *http.Request){
	
	t,_ := template.ParseFiles("pages/home.gtpl")
	t.Execute(w, nil)
}


func register(w http.ResponseWriter, r *http.Request) {
	total:= 3
	if r.Method == "GET"{
		t, _ := template.ParseFiles("pages/registration.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		var conn, _ = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		var c =  proto.NewCommsClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err :=c.RegisterAccount(ctx,&proto.AccountInfo{Username: r.FormValue("username"), Password: r.FormValue("password") })
		if err != nil{
			log.Println(err)
			time.Sleep(4)
			http.Redirect(w, r, "/login", 301)
		}

		if resp.GetMessage() == true{
			http.Redirect(w, r, "/login", 303)
		} else {
			http.Redirect(w, r, "/registration", 303)
		}


	}

}


func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t,_ := template.ParseFiles("pages/login.gtpl")
		t.Execute(w, nil)
	} else{

		r.ParseForm()
		log.Printf("username submitted: %s \npassword submitted: %s", r.FormValue("username"),r.FormValue("password"))
		var conn, _ = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		var c =  proto.NewCommsClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		log.Printf("after rpc creation")
		defer cancel()
		//reorder two methods below
		responseRPC,rpcerror := c.SendAccountInfo(ctx,&proto.AccountInfo{Username: r.FormValue("username"), Password: r.FormValue("password") })
		resp2, rpcerr2 := c.DoesAccountExist(ctx, &proto.AccountInfo{Username: r.FormValue("username"), Password: r.FormValue("password") })
		if rpcerr2 != nil{
			log.Println("Account does not exist")
			log.Println(resp2.GetMessage())
			fmt.Fprintf(w,rpcerr2.Error())
			http.Redirect(w, r, "/login", 301)
		}
		if rpcerror != nil {
			http.Redirect(w, r, "/login", 301)
			fmt.Fprintf(w, rpcerror.Error())
			
		}
		
		//log.Println("Successful message")
		log.Println(responseRPC.GetMessage())
	}
}

//func accountexistence(user string, pass string)

//need to add authentication token
	


func main() {
	fmt.Println("Starting from project")
	go runServers()
	 
	
	//shttp.HandleFunc("/test", )
	http.HandleFunc("/register",register)
	http.HandleFunc("/",mainPage)
	http.HandleFunc("/login", login)
	err:= http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
	runServers()
	
}
