//SERVER SOCKET: 

package main 

import ( 
	"fmt" 
	"net" 
	"bufio" 

	) 


// Program for sever socket
func main() { 
	fmt.Println("Server is Starting") 
	//Listens to port 8080
	ln, _ := net.Listen("tcp", ":8080") 
	//Accepts connection
	conn, _ := ln.Accept() 
	//Loop runs forever until 
	for {
		//Process data
		message, _ := bufio.NewReader(conn).ReadString('\n') 
		//output message
		fmt.Println("Message Received: ", string(message))
	}
} 

//cLIENT SOCKET

package main 

import ( 
	"fmt" 
	"net" 
	"bufio" 
	"os"
	) 

func main() { 
	//Connect server
	conn, _ := net.Dial("tcp", "127.0.0.1:8080") 
	for { 
		//Something to send
		reader := bufio.NewReader(os.Stdin) 
		text, _ := reader.ReadString('\n') 
		//send to sever
		fmt.Fprint(conn, text+"\n") 
		//Waite for the reply
		message, _ := bufio.NewReader(conn).ReadString('\n') 
		fmt.Println("Message from sever: "+ message)
	} 
}
