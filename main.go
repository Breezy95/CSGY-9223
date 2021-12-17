package main


import "fmt"
import "webAppProject/servers"

func main() {
	fmt.Println("Starting from project")
	//go servers.ClientRun()
	 servers.BackendRun()


}
