package main
import (
	"context" 
     "log"
     "net"

     pb "ch13/hellopb" 
     "google.golang.org/grpc" 
)
type server struct{} 
func (s *server) Say(ctx context.Context, \ 
 in *pb.HelloRequest) (*pb.HelloResponse, error) { 
     msg := "Hello " + in.Name + "!" 
     return &pb.HelloResponse{Message: msg}, nil 
} 
func main() {
     l, err := net.Listen("tcp", ":55555") 
     if err != nil { 
            ...  
     }  
     s := grpc.NewServer() 
     pb.RegisterHelloServer(s, &server{})
     s.Serve(l) 
} 
